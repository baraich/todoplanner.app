package resource

// TodoPlannerKye structs holds the public and private key information
// related to the TodoPlanner application. These keys are used for
// authentication purposed and were set using the (node)sst.Linkable class.
type TodoPlannerKey struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}

// _Resource acts as a container for various resources used by
// the TodoPlanner application.
type _Resource struct {
	TodoPlannerKey TodoPlannerKey
}

// Resource is a global variable of type _Resource that stores the
// application's key data. It can be accessed throughout the program
// for tasks that require access to infrastructure.
var Resource _Resource

// TODO: Implementation of reflection technique to populate resources.
