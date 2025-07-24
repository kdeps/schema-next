// Code generated from Pkl module `org.kdeps.pkl.Session`. DO NOT EDIT.
package session

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/kdeps/schema/gen/utils"
)

type Session interface {
	utils.Utils
}

var _ Session = (*SessionImpl)(nil)

// Abstractions for Session records
type SessionImpl struct {
	*utils.UtilsImpl
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Session
func LoadFromPath(ctx context.Context, path string) (ret Session, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Session
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Session, error) {
	var ret SessionImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
