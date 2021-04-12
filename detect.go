package condaenvupdate

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

// Detect returns a packit.DetectFunc that will be invoked during the
// detect phase of the buildpack lifecycle.
//
// Detection always passes, and will contribute a Build Plan that provides conda.
func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		_, err := os.Stat(filepath.Join(context.WorkingDir, "environment.yml"))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return packit.DetectResult{}, packit.Fail
			}

			return packit.DetectResult{}, packit.Fail.WithMessage("failed trying to stat environment.yml: %w", err)
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "conda-environment"},
				},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "conda",
						Metadata: map[string]interface{}{
							"build": true,
						},
					},
				},
			},
		}, nil
	}
}
