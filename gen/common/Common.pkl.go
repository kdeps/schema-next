// Code generated from Pkl module `org.kdeps.pkl.Common`. DO NOT EDIT.
package common

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Common interface {
}

var _ Common = (*CommonImpl)(nil)

// Common utility functions used across all PKL modules
// This module provides standardized implementations of frequently used patterns
// to ensure consistency and reduce code duplication across resource modules.
//
// **MEMORY-ONLY PROCESSING POLICY:**
// All functions in this module support the kdeps memory-first approach:
// - No temporary file creation during processing
// - All data processing happens in-memory for optimal performance
// - Functions prioritize memory-efficient operations over file I/O
// - Caching and memoization used extensively to avoid redundant processing
type CommonImpl struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Common
func LoadFromPath(ctx context.Context, path string) (ret Common, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Common
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Common, error) {
	var ret CommonImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
