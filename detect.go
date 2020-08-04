package gomodvendor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		_, err := os.Stat(filepath.Join(context.WorkingDir, "go.mod"))
		if err != nil {
			if os.IsNotExist(err) {
				return packit.DetectResult{}, packit.Fail
			}
			return packit.DetectResult{}, fmt.Errorf("failed to stat go.mod: %w", err)
		}

		_, err = os.Stat(filepath.Join(context.WorkingDir, "vendor"))
		if err == nil {
			return packit.DetectResult{}, packit.Fail
		} else {
			if !os.IsNotExist(err) {
				return packit.DetectResult{}, fmt.Errorf("failed to stat vendor directory: %w", err)
			}
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "go",
						Metadata: map[string]interface{}{
							"build": true,
						},
					},
				},
			},
		}, nil
	}
}