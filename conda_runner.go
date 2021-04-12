package condaenvupdate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/pexec"
)

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

func (c CondaRunner) Execute(condaEnvPath string, condaCachePath string, workingDir string) error {
	return c.executable.Execute(pexec.Execution{
		Args: []string{"env", "update", "--prefix", condaEnvPath, "--file", filepath.Join(workingDir, "environment.yml")},
		Env:  append(os.Environ(), fmt.Sprintf("CONDA_PKGS_DIRS=%s", condaCachePath)),
	})
}
