package main

import (
	_ "embed"
	"fmt"
	"go.starlark.net/starlark"
	"log"
)

//go:embed starlark_scripts/firstFunction.star
var firstFunction []byte

// A Go function to be called from Starlark
func myGoFunction(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	// Example: Expecting one string argument
	if len(args) != 1 {
		return nil, fmt.Errorf("expected exactly one argument")
	}

	str, ok := starlark.AsString(args[0])
	if !ok {
		return nil, fmt.Errorf("argument is not a string")
	}

	// Process the argument and return a result
	return starlark.String("Hello, Jonas " + str), nil
}

func main() {
	// Starlark script
	script := string(firstFunction) //`print(hello_world("Starlark"))`

	// Create a Starlark thread
	thread := &starlark.Thread{Name: "main"}

	// Expose the Go function to Starlark
	builtins := starlark.StringDict{
		"my_go_function": starlark.NewBuiltin("my_go_function", myGoFunction),
		"parameter":      starlark.String("Denna kommer från Go"),
	}

	// ****************************************
	// Insert 'parameter' and Execute the full script and extract response variable

	// Execute the Starlark script
	globals, err := starlark.ExecFile(thread, "", script, builtins)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("************************Insert 'parameter' and Execute the full script and extract response variable***********************************")
	fmt.Println("From Go", globals["response"])
	fmt.Println("***********************************************************")
	fmt.Println("")

	// ****************************************
	// Execute a specific starlark function

	// Retrieve the function
	processFn, ok := globals["printHello"]
	if !ok {
		log.Fatalf("Function 'printHello' not found")
	}

	// Prepare the arguments for the function
	args := starlark.Tuple{starlark.String("Hello, Starlark from Go!")}

	// Call the Starlark function
	result, err := starlark.Call(thread, processFn, args, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Use the result
	fmt.Println("************************ Execute a specific starlark function***********************************")
	fmt.Printf("Result from Starlark function: %v\n", result)
	fmt.Println("***********************************************************")
	fmt.Println("")

	// ****************************************
	// Call a Starlark function that then calls a Go function

	// Retrieve the function
	processFn2, ok := globals["starlark_def_function"]
	if !ok {
		log.Fatalf("Function 'printHello' not found")
	}

	// Call the Starlark function
	result2, err := starlark.Call(thread, processFn2, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Use the result
	fmt.Println("************************ Execute a specific starlark function***********************************")
	fmt.Printf("Result from Starlark function: %v\n", result2)
	fmt.Println("***********************************************************")
	fmt.Println("")

}
