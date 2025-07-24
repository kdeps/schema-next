// Code generated from Pkl module `org.kdeps.pkl.PklResource`. DO NOT EDIT.
package pklresource

import "github.com/apple/pkl-go/pkl"

// Relational Algebra Functions
type SelectionCondition struct {
	Field string `pkl:"field"`

	Operator string `pkl:"operator"`

	Value *pkl.Object `pkl:"value"`
}
