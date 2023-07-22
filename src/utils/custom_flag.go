package utils

import "flag"

// BoolFlag create flag for short and long argument for bool with false as the defautl arg
func CustomBoolFlag(arg *bool, name string, description string) {
	// Long flag
	flag.BoolVar(arg, name, false, "")

	// Short flag
	flag.BoolVar(arg, string(name[0]), false, description)
}

// StringFlag create flag for short and long argument for bool with "" as the default argument
func CustomStringFlag(arg *string, name string, description string) {
	// Long flag
	flag.StringVar(arg, name, "", "")

	// Short flag
	flag.StringVar(arg, string(name[0]), "", description)
}
