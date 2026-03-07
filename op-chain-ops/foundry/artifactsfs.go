package foundry

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

type StatDirFs interface {
	fs.StatFS
	fs.ReadDirFS
}

func OpenArtifactsDir(dirPath string) *ArtifactsFS {
	dir := os.DirFS(dirPath)
	if d, ok := dir.(StatDirFs); !ok {
		panic("Go DirFS guarantees changed")
	} else {
		return &ArtifactsFS{FS: d}
	}
}

type EmbedFS struct {
	FS      embed.FS
	RootDir string // Root directory within the embedded FS
}

// Open implements fs.FS
func (e *EmbedFS) Open(name string) (fs.File, error) {
	return e.FS.Open(path.Join(e.RootDir, name))
}

// Stat implements fs.StatFS
func (e *EmbedFS) Stat(name string) (fs.FileInfo, error) {
	file, err := e.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file.Stat()
}

// ReadDir implements fs.ReadDirFS
func (e *EmbedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return fs.ReadDir(e.FS, path.Join(e.RootDir, name))
}

// ArtifactsFS wraps a filesystem (read-only access) of a forge-artifacts bundle.
// The root contains directories for every artifact,
// each containing one or more entries (one per solidity compiler version) for a solidity contract.
// See OpenArtifactsDir for reading from a local directory.
// Alternative FS systems, like a tarball, may be used too.
type ArtifactsFS struct {
	FS StatDirFs
}

// ListArtifacts lists the artifacts. Each artifact matches a source-file name.
// This name includes the extension, e.g. ".sol"
// (no other artifact-types are supported at this time).
func (af *ArtifactsFS) ListArtifacts() ([]string, error) {
	entries, err := af.FS.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to list artifacts: %w", err)
	}
	out := make([]string, 0, len(entries))
	for _, d := range entries {
		// Some artifacts may be nested in directories not suffixed with ".sol"
		// Nested artifacts, and non-solidity artifacts, are not supported.
		if name := d.Name(); strings.HasSuffix(name, ".sol") {
			out = append(out, d.Name())
		}
	}
	return out, nil
}

// ListContracts lists the contracts of the named artifact, including the file extension.
// E.g. "Owned.sol" might list "Owned.0.8.15", "Owned.0.8.25", and "Owned".
func (af *ArtifactsFS) ListContracts(name string) ([]string, error) {
	f, err := af.FS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open artifact %q: %w", name, err)
	}
	defer f.Close()
	dirFile, ok := f.(fs.ReadDirFile)
	if !ok {
		return nil, fmt.Errorf("no dir for artifact %q, but got %T", name, f)
	}
	entries, err := dirFile.ReadDir(0)
	if err != nil {
		return nil, fmt.Errorf("failed to list artifact contents of %q: %w", name, err)
	}
	out := make([]string, 0, len(entries))
	for _, d := range entries {
		if name := d.Name(); strings.HasSuffix(name, ".json") {
			out = append(out, strings.TrimSuffix(name, ".json"))
		}
	}
	return out, nil
}

// ReadArtifact reads a specific JSON contract artifact from the FS.
// The contract name may be suffixed by a solidity compiler version, e.g. "Owned.0.8.25".
// The contract name does not include ".json", this is a detail internal to the artifacts.
// The name of the artifact is the source-file name, this must include the suffix such as ".sol".
// If name contains a path (e.g. "src/universal/Proxy.sol"), the full path is tried first;
// if that fails, a fallback to the base name (e.g. "Proxy.sol") is tried for flat artifact layouts.
// If that also fails, the FS is walked to find any path ending with name/contract.json (for nested layouts).
func (af *ArtifactsFS) ReadArtifact(name string, contract string) (*Artifact, error) {
	artifactPath := path.Join(name, contract+".json")
	f, err := af.FS.Open(artifactPath)
	if err != nil {
		// Fallback for flat artifact bundles that only have File.sol/Contract.json (no path prefix)
		if base := path.Base(name); base != name {
			artifactPath = path.Join(base, contract+".json")
			f, err = af.FS.Open(artifactPath)
		}
		if err != nil {
			// Fallback for nested layouts: find path ending with name/contract.json (e.g. Proxy.sol/Proxy.json)
			artifactPath, err = af.findArtifactPath(name, contract+".json")
			if err != nil {
				return nil, fmt.Errorf("failed to open artifact %s/%s: %w", name, contract, err)
			}
			f, err = af.FS.Open(artifactPath)
			if err != nil {
				return nil, fmt.Errorf("failed to open artifact %q: %w", artifactPath, err)
			}
		}
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var out Artifact
	if err := dec.Decode(&out); err != nil {
		return nil, fmt.Errorf("failed to decode artifact %q: %w", name, err)
	}
	return &out, nil
}

// findArtifactPath walks the FS to find a path ending with artifactName/contractFile (e.g. Proxy.sol/Proxy.json).
// Supports flat layout (File.sol/Contract.json at root) when the script requests a path like src/universal/Proxy.sol.
func (af *ArtifactsFS) findArtifactPath(artifactName, contractFile string) (string, error) {
	target := path.Join(artifactName, contractFile)
	var found string
	err := fs.WalkDir(af.FS, ".", func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		clean := path.Clean(p)
		if clean != p {
			p = clean
		}
		if p == target || strings.HasSuffix(p, "/"+target) {
			found = p
			return fs.SkipAll
		}
		if strings.HasSuffix(p, "/"+contractFile) && path.Base(path.Dir(p)) == artifactName {
			found = p
			return fs.SkipAll
		}
		return nil
	})
	if err != nil && err != fs.SkipAll {
		return "", err
	}
	if found == "" {
		return "", fmt.Errorf("no path ending with %s", target)
	}
	return found, nil
}
