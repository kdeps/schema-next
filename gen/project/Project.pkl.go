// Code generated from Pkl module `org.kdeps.pkl.Project`. DO NOT EDIT.
package project

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Project interface {
}

var _ Project = (*ProjectImpl)(nil)

// Abstractions for Kdeps Project Settings
//
// This module defines the structure for project-specific settings in the Kdeps system. It includes
// configurations related to the API server, Docker agent settings, and security settings. These
// settings allow customization of the project's environment, such as enabling API server mode or
// configuring Docker and security parameters.
type ProjectImpl struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Project
func LoadFromPath(ctx context.Context, path string) (ret Project, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Project
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Project, error) {
	var ret ProjectImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
