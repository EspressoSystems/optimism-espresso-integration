package enclave_tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

const (
	// ArgDeliveryPort is the vsock port for batcher arg delivery. Must match NC_PORT in enclave-entrypoint.bash.
	ArgDeliveryPort uint16 = 8337
	// ReadinessPort is the vsock port for the readiness handshake. Must match READY_PORT in enclave-entrypoint.bash.
	ReadinessPort uint16 = 8338
	// ArgDeliveryHostPort is the host-side TCP port docker --publish maps to ArgDeliveryPort.
	ArgDeliveryHostPort = 9000
)

type EnclaverManifestSources struct {
	App string `yaml:"app"`
}

type EnclaverManifestDefaults struct {
	CpuCount uint `yaml:"cpu_count"`
	MemoryMb uint `yaml:"memory_mb"`
}

type EnclaverManifestKmsProxy struct {
	ListenPort uint16 `yaml:"listen_port,omitempty"`
}

type EnclaverManifestEgress struct {
	Allow     []string `yaml:"allow"`
	Deny      []string `yaml:"deny"`
	ProxyPort uint16   `yaml:"proxy_port,omitempty"`
}

type EnclaverManifestIngress struct {
	ListenPort uint16 `yaml:"listen_port"`
}

type EnclaverManifest struct {
	Version  string                    `yaml:"version"`
	Name     string                    `yaml:"name"`
	Target   string                    `yaml:"target"`
	Sources  *EnclaverManifestSources  `yaml:"sources,omitempty"`
	Defaults *EnclaverManifestDefaults `yaml:"defaults,omitempty"`
	KmsProxy *EnclaverManifestKmsProxy `yaml:"kms_proxy,omitempty"`
	Egress   *EnclaverManifestEgress   `yaml:"egress,omitempty"`
	Ingress  []EnclaverManifestIngress `yaml:"ingress"`
}

func DefaultManifest(name string, target string, source string, cpuCount uint, memoryMb uint) EnclaverManifest {
	return EnclaverManifest{
		Version: "v1",
		Name:    name,
		Target:  target,
		Sources: &EnclaverManifestSources{
			App: source,
		},
		Defaults: &EnclaverManifestDefaults{
			CpuCount: cpuCount,
			MemoryMb: memoryMb,
		},
		Egress: &EnclaverManifestEgress{
			ProxyPort: 10000,
			Allow:     []string{"0.0.0.0/0", "**", "::/0"},
		},
		Ingress: []EnclaverManifestIngress{
			{ListenPort: ArgDeliveryPort}, // batcher arg delivery
			{ListenPort: ReadinessPort},   // readiness handshake
		},
	}
}

type EnclaverBuildOutput struct {
	Measurements EnclaveMeasurements `json:"Measurements"`
}

type EnclaverCli struct{}

// BuildEnclave builds an enclaver EIF image using the provided manifest. If build is successful,
// it returns the image's Measurements.
func (*EnclaverCli) BuildEnclave(ctx context.Context, manifest EnclaverManifest) (EnclaveMeasurements, error) {
	tempfile, err := os.CreateTemp("", "enclaver-manifest")
	if err != nil {
		return EnclaveMeasurements{}, err
	}
	defer os.Remove(tempfile.Name())

	if err := yaml.NewEncoder(tempfile).Encode(manifest); err != nil {
		return EnclaveMeasurements{}, err
	}

	var stdout bytes.Buffer
	cmd := exec.CommandContext(
		ctx,
		"enclaver",
		"build",
		"--file",
		tempfile.Name(),
	)
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return EnclaveMeasurements{}, err
	}

	// Find measurements in the output
	re := regexp.MustCompile(`\{[\s\S]*"Measurements"[\s\S]*\}`)
	jsonMatch := re.Find(stdout.Bytes())
	if jsonMatch == nil {
		return EnclaveMeasurements{}, fmt.Errorf("could not find measurements JSON in output")
	}

	var output EnclaverBuildOutput
	if err := json.Unmarshal(jsonMatch, &output); err != nil {
		return EnclaveMeasurements{}, fmt.Errorf("failed to parse measurements JSON: %w", err)
	}

	return output.Measurements, nil
}

// RunEnclave runs an enclaver EIF image `name` with the provided arguments.
// Uses 'docker run' directly (not 'enclaver run') to support --publish.
// --publish=127.0.0.1:ArgDeliveryHostPort:ArgDeliveryPort instead of --net=host keeps
// the container off the host network stack, blocking EC2 metadata-service access
// (requires IMDSv2 with HttpPutResponseHopLimit=1 on the instance).
func (*EnclaverCli) RunEnclave(ctx context.Context, name string, args []string) error {
	// We'll append this to container name to avoid conflicts
	nameSuffix := uuid.New().String()[:8]

	cmd := exec.CommandContext(
		ctx,
		"docker",
		"run",
		"--rm",
		"-d",
		"--privileged",
		fmt.Sprintf("--publish=127.0.0.1:%d:%d", ArgDeliveryHostPort, ArgDeliveryPort),
		"--name=batcher-enclaver-"+nameSuffix,
		"--device=/dev/nitro_enclaves",
		name,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Info("Starting enclave container", "command", cmd.Args)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start enclave container: %w", err)
	}

	// Wait for container to start
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("enclave exited with an error: %w", err)
	}

	// Send arguments to enclave via nc listener
	if err := sendArgsToEnclave(ctx, args); err != nil {
		return fmt.Errorf("failed to send arguments to enclave: %w", err)
	}

	return nil
}

// sendArgsToEnclave sends arguments to the enclave's nc listener as null-separated values
func sendArgsToEnclave(ctx context.Context, args []string) error {
	// Prepare arguments as null-separated bytes
	var buf bytes.Buffer
	for _, arg := range args {
		buf.WriteString(arg)
		buf.WriteByte(0) // null separator
	}
	buf.WriteByte(0) // double null to signal end

	// Create a dialer with short timeout for individual attempts
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	// Retry connecting for up to 1 minute
	retryDuration := 60 * time.Second
	retryInterval := 2 * time.Second
	deadline := time.Now().Add(retryDuration)

	for time.Now().Before(deadline) {
		// Connect to the enclave's listener
		conn, err := dialer.DialContext(ctx, "tcp", fmt.Sprintf("127.0.0.1:%d", ArgDeliveryHostPort))
		if err != nil {
			// If we still have time, wait and retry
			if time.Now().Add(retryInterval).Before(deadline) {
				time.Sleep(retryInterval)
				continue
			}
			return fmt.Errorf("failed to connect to enclave listener after %v: %w", retryDuration, err)
		}
		defer conn.Close()

		// Send the arguments
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			conn.Close()
			return fmt.Errorf("failed to send arguments to enclave: %w", err)
		}

		return nil
	}

	return fmt.Errorf("timeout connecting to enclave listener after %v", retryDuration)
}
