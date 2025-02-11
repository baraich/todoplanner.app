package resource

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

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

func init() {
	val := reflect.ValueOf(&Resource).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)
		envVarName := fmt.Sprintf("SST_RESOURCE_%s", typeField.Name)
		envValue, exists := os.LookupEnv(envVarName)
		if !exists {
			panic(fmt.Sprintf("Environment variable %s is required!", envVarName))
		}
		if err := json.Unmarshal([]byte(envValue), field.Addr().Interface()); err != nil {
			panic(err)
		}
	}
}
