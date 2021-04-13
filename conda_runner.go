package condaenvupdate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/pexec"
)

//go:generate faux --interface Executable --output fakes/executable.go
type Executable interface {
	Execute(pexec.Execution) error
}

type CondaRunner struct {
	executable Executable
}

func NewCondaRunner(executable Executable) CondaRunner {
	return CondaRunner{
		executable: executable,
	}
}

func (c CondaRunner) Execute(condaLayerPath string, condaCachePath string, workingDir string) error {
	err := c.executable.Execute(pexec.Execution{
		Args: []string{
			"env",
			"update",
			"--prefix", condaLayerPath,
			"--file", filepath.Join(workingDir, "environment.yml"),
		},
		Env: append(os.Environ(), fmt.Sprintf("CONDA_PKGS_DIRS=%s", condaCachePath)),
	})
	if err != nil {
		return err
	}

	return nil
}
