// Code generated from Pkl module `org.kdeps.pkl.Kdeps`. DO NOT EDIT.
package kdeps

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/kdeps/schema/gen/kdeps/gpu"
	"github.com/kdeps/schema/gen/kdeps/path"
	"github.com/kdeps/schema/gen/kdeps/runmode"
)

// Abstractions for Kdeps Configuration
type Kdeps struct {
	// The mode of execution for Kdeps, defaulting to "docker".
	Mode *runmode.RunMode `pkl:"Mode"`

	// The GPU type to use for Kdeps, defaulting to "cpu".
	DockerGPU *gpu.GPU `pkl:"DockerGPU"`

	// The directory where Kdeps files are stored, defaulting to ".kdeps".
	KdepsDir *string `pkl:"KdepsDir"`

	// The path where Kdeps configurations are stored, defaulting to "user".
	KdepsPath *path.Path `pkl:"KdepsPath"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Kdeps
func LoadFromPath(ctx context.Context, path string) (ret *Kdeps, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Kdeps
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Kdeps, error) {
	var ret Kdeps
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
