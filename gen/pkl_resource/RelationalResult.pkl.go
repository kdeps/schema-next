// Code generated from Pkl module `org.kdeps.pkl.PklResource`. DO NOT EDIT.
package pklresource

import "github.com/apple/pkl-go/pkl"

type RelationalResult struct {
	Rows []*pkl.Object `pkl:"rows"`

	Columns []string `pkl:"columns"`

	Query string `pkl:"query"`

	Ttl string `pkl:"ttl"`
}
