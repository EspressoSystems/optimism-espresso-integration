package environment

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os/exec"
	"strings"
)

type DockerContainerInfo struct {
	// The container ID of the Docker container that is running the
	// Espresso Dev Node.
	// This is useful for further interaction with docker concerning
	// the specific Dev Node
	ContainerID string

	// The Port Map of the Resulting Docker Container
	PortMap map[string][]string
}

// DockerContainerConfig is a configuration struct that is used to configure
// the launching of a Docker Container
type DockerContainerConfig struct {
	Image string

	Environment map[string]string

	Ports []string

	AutoRM bool
}

// DockerCli is a simple implementation of a Docker Client that is used to
// launch Docker Containers
type DockerCli struct{}

// LaunchContainer launches a Docker Container with the given configuration
// and returns the resulting Docker Container Info
//
// The Container will automatically be stopped when the given context is
// completed. This is done by spawning a goroutine that is blocked by the
// context that is passed in's Done channel.
func (d *DockerCli) LaunchContainer(ctx context.Context, config DockerContainerConfig) (DockerContainerInfo, error) {
	originalContext := ctx

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	outputBuffer := new(bytes.Buffer)
	var args []string
	// Let's build the arguments for the docker launch command
	{

		args = append(args, "run", "-d")

		if config.AutoRM {
			args = append(args, "--rm")
		}

		for _, port := range config.Ports {
			args = append(args, "-p", port)
		}

		for key, value := range config.Environment {
			args = append(args, "-e", key+"="+value)
		}

		args = append(args, config.Image)
	}

	var containerID string
	{
		launchContainerCmd := exec.CommandContext(
			ctx,
			"docker",
			args...,
		)

		// A buffer to collect the output of the command, so we can retrieve the
		// Container ID.
		launchContainerCmd.Stdout = outputBuffer

		if err := launchContainerCmd.Run(); err != nil {
			return DockerContainerInfo{}, err
		}

		containerID = strings.TrimSpace(outputBuffer.String())
	}

	// Let's setup a cleanup function to stop the container, should we
	// need to.

	stopContainer := func() error {
		return d.StopContainer(context.Background(), containerID)
	}

	// We spin up a goroutine that will clean us up when the original context
	// dies
	go (func(ctx context.Context) {
		// Wait for the context that governs us to tell us to die
		<-ctx.Done()

		stopContainer()
	})(originalContext)

	// We have the container ID.  Let's get our Ports

	portMap := map[string][]string{}
	containerInfo := DockerContainerInfo{ContainerID: containerID, PortMap: portMap}
	{
		for _, portToFind := range config.Ports {
			outputBuffer.Reset()
			// Let's find out what our assigned ports ended up being
			determinePortCmd := exec.CommandContext(
				ctx,
				"docker",
				"port",
				containerID,
				portToFind,
			)
			determinePortCmd.Stdout = outputBuffer

			if err := determinePortCmd.Run(); err != nil {
				return containerInfo, err
			}

			lineReader := bufio.NewReader(outputBuffer)

			for {
				line, _, err := lineReader.ReadLine()
				if err == io.EOF {
					// we consumed all of it
					break
				}

				if err != nil {
					return DockerContainerInfo{ContainerID: containerID}, err
				}

				if len(line) == 0 {
					// empty line, ignore
					continue
				}

				portMap[portToFind] = append(portMap[portToFind], string(line))
			}
		}
	}

	return containerInfo, nil
}

// StopContainer stops a Docker Container with the given container ID
func (d *DockerCli) StopContainer(ctx context.Context, containerID string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stopCmd := exec.CommandContext(
		ctx,
		"docker",
		"stop",
		containerID,
	)

	return stopCmd.Run()
}
