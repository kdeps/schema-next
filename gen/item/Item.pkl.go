// Code generated from Pkl module `org.kdeps.pkl.Item`. DO NOT EDIT.
package item

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/kdeps/schema/gen/utils"
)

type Item interface {
	utils.Utils
}

var _ Item = (*ItemImpl)(nil)

// Abstractions for Item iteration records
//
// This module provides functions to interact with records representing iterations or elements in a for loop.
// The module supports retrieving, navigating, and listing records without requiring a specific identifier.
type ItemImpl struct {
	*utils.UtilsImpl
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Item
func LoadFromPath(ctx context.Context, path string) (ret Item, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Item
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Item, error) {
	var ret ItemImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
