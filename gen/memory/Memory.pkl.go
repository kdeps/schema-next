// Code generated from Pkl module `org.kdeps.pkl.Memory`. DO NOT EDIT.
package memory

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/kdeps/schema/gen/utils"
)

type Memory interface {
	utils.Utils
}

var _ Memory = (*MemoryImpl)(nil)

// Abstractions for Memory records
type MemoryImpl struct {
	*utils.UtilsImpl
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Memory
func LoadFromPath(ctx context.Context, path string) (ret Memory, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Memory
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Memory, error) {
	var ret MemoryImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
