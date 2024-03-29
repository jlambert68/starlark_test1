package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"log"
)

// Define a function that can be called from Starlark
func helloWorld(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, _ []starlark.Tuple) (starlark.Value, error) {
	if args.Len() != 1 {
		return nil, fmt.Errorf("expected exactly one argument")
	}

	name, ok := starlark.AsString(args.Index(0))
	if !ok {
		return nil, fmt.Errorf("argument is not a string")
	}

	return starlark.String("Hello, " + name), nil
}

func main() {
	// Starlark script
	script := `print(hello_world("Starlark"))`

	// Create a Starlark thread
	thread := &starlark.Thread{Name: "main"}

	// Expose the Go function to Starlark
	builtins := starlark.StringDict{
		"hello_world": starlark.NewBuiltin("hello_world", helloWorld),
	}

	// Execute the Starlark script
	_, err := starlark.ExecFile(thread, "hello.star", script, builtins)
	if err != nil {
		log.Fatal(err)
	}
}
