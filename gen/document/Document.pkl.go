// Code generated from Pkl module `org.kdeps.pkl.Document`. DO NOT EDIT.
package document

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/kdeps/schema/gen/utils"
)

type Document interface {
	utils.Utils
}

var _ Document = (*DocumentImpl)(nil)

// Common parser and document renderer functions used across all resources.
//
// Tools for Parsing and Generating JSON, YAML and XML documents
type DocumentImpl struct {
	*utils.UtilsImpl
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Document
func LoadFromPath(ctx context.Context, path string) (ret Document, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Document
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Document, error) {
	var ret DocumentImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
