package condaenvupdate

import (
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

//go:generate faux --interface Runner --output fakes/runner.go
type Runner interface {
	Execute(condaEnvPath string, condaCachePath string, workingDir string) error
}

func Build(runner Runner, logger scribe.Logger) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		// Get conda-env-layer
		condaLayer, err := context.Layers.Get("conda-env")
		if err != nil {
			return packit.BuildResult{}, err
		}

		condaLayer, err = condaLayer.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		condaLayer.Launch = true

		condaCacheLayer, err := context.Layers.Get("conda-env-cache")
		if err != nil {
			return packit.BuildResult{}, err
		}

		condaCacheLayer.Cache = true

		err = runner.Execute(condaLayer.Path, condaCacheLayer.Path, context.WorkingDir)
		if err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Layers: []packit.Layer{
				condaLayer,
				condaCacheLayer,
			},
		}, nil
	}
}
