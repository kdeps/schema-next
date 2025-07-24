// Code generated from Pkl module `org.kdeps.pkl.Resource`. DO NOT EDIT.
package resource

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Resource interface {
	GetActionID() string

	GetName() *string

	GetDescription() *string

	GetCategory() *string

	GetRequires() *[]string

	GetItems() *[]string

	GetRun() *ResourceAction
}

var _ Resource = (*ResourceImpl)(nil)

// Abstractions for Kdeps Resources
//
// This module defines the structure for resources used within the Kdeps framework,
// including actions that can be performed on these resources, validation checks,
// and error handling mechanisms. Each resource can define its actionID, name, description,
// category, dependencies, and how it runs.
//
// **MEMORY-ONLY PROCESSING POLICY:**
// - All resource processing is done in-memory to maximize performance
// - No temporary files are created during resource execution
// - APIResponse blocks are processed directly in memory and stored for later use
// - Only the final target action response is persisted to disk
// - Intermediate resource responses remain in memory-only storage
//
// **EXECUTION FLOW:**
// - Resources execute in dependency order until the target action is reached
// - Each resource with APIResponse stores its response in memory
// - Processing continues beyond intermediate response resources
// - Only when TargetActionID is reached does the workflow terminate
// - This allows for complex multi-step workflows with intermediate API responses
type ResourceImpl struct {
	// The unique identifier for the resource, validated against [isValidActionID].
	ActionID string `pkl:"ActionID"`

	// The name of the resource.
	Name *string `pkl:"Name"`

	// A description of the resource, providing additional context.
	Description *string `pkl:"Description"`

	// The category to which the resource belongs.
	Category *string `pkl:"Category"`

	// A listing of dependencies required by the resource, validated against [isValidDependency].
	Requires *[]string `pkl:"Requires"`

	// Defines the action items to be processed individually in a loop.
	Items *[]string `pkl:"Items"`

	// Defines the action to be taken for the resource.
	Run *ResourceAction `pkl:"Run"`
}

// The unique identifier for the resource, validated against [isValidActionID].
func (rcv *ResourceImpl) GetActionID() string {
	return rcv.ActionID
}

// The name of the resource.
func (rcv *ResourceImpl) GetName() *string {
	return rcv.Name
}

// A description of the resource, providing additional context.
func (rcv *ResourceImpl) GetDescription() *string {
	return rcv.Description
}

// The category to which the resource belongs.
func (rcv *ResourceImpl) GetCategory() *string {
	return rcv.Category
}

// A listing of dependencies required by the resource, validated against [isValidDependency].
func (rcv *ResourceImpl) GetRequires() *[]string {
	return rcv.Requires
}

// Defines the action items to be processed individually in a loop.
func (rcv *ResourceImpl) GetItems() *[]string {
	return rcv.Items
}

// Defines the action to be taken for the resource.
func (rcv *ResourceImpl) GetRun() *ResourceAction {
	return rcv.Run
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Resource
func LoadFromPath(ctx context.Context, path string) (ret Resource, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Resource
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Resource, error) {
	var ret ResourceImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
