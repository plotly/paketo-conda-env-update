package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/pexec"
	"github.com/paketo-buildpacks/packit/scribe"
	condaenvupdate "github.com/paketo-community/conda-env-update"
)

func main() {
	logger := scribe.NewLogger(os.Stdout)
	condaRunner := condaenvupdate.NewCondaRunner(pexec.NewExecutable("conda"))

	packit.Run(condaenvupdate.Detect(), condaenvupdate.Build(condaRunner, logger))
}
