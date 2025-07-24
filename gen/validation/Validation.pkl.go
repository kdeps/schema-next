// Code generated from Pkl module `org.kdeps.pkl.Validation`. DO NOT EDIT.
package validation

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Validation interface {
}

var _ Validation = (*ValidationImpl)(nil)

// Common validation patterns and regex definitions
// This module provides standardized validation functions and regular expressions
// to ensure consistent input validation across all resource modules.
type ValidationImpl struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Validation
func LoadFromPath(ctx context.Context, path string) (ret Validation, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Validation
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Validation, error) {
	var ret ValidationImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
