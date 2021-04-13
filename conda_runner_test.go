package condaenvupdate_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	condaenvupdate "github.com/paketo-community/conda-env-update"
	"github.com/paketo-community/conda-env-update/fakes"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testCondaRunner(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir     string
		condaLayerPath string
		condaCachePath string

		executable *fakes.Executable

		runner condaenvupdate.CondaRunner
	)

	it.Before(func() {
		var err error

		workingDir, err = os.MkdirTemp("", "working-dir")
		Expect(err).NotTo(HaveOccurred())
		condaLayerPath = "a-conda-layer"
		condaCachePath = "a-conda-cache-path"

		executable = &fakes.Executable{}
		runner = condaenvupdate.NewCondaRunner(executable)
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("Execute", func() {
		it("runs conda env update", func() {
			err := runner.Execute(condaLayerPath, condaCachePath, workingDir)
			Expect(err).NotTo(HaveOccurred())

			Expect(executable.ExecuteCall.Receives.Execution.Args).To(Equal([]string{
				"env",
				"update",
				"--prefix", "a-conda-layer",
				"--file", filepath.Join(workingDir, "environment.yml"),
			}))
			Expect(executable.ExecuteCall.Receives.Execution.Env).To(Equal(append(os.Environ(), fmt.Sprintf("CONDA_PKGS_DIRS=%s", "a-conda-cache-path"))))
		})

		context("failures cases", func() {
			context("when the conda command fails to run", func() {
				it.Before(func() {
					executable.ExecuteCall.Returns.Error = errors.New("failed to run conda command")
				})

				it("returns an error", func() {
					err := runner.Execute(condaLayerPath, condaCachePath, workingDir)
					Expect(err).To(MatchError("failed to run conda command"))
				})
			})
		})
	})
}
