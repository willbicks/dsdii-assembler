package main

import (
	"fmt"
	"runtime/debug"
)

// These variables are set by the linker at build time.
var (
	version string // program version string
	vcsHash string // git commit hash
)

// parseVersionInfo overwrites the package variables with embedded build info, if variables
// are not already set by linker flags.
func parseVersionInfo() {
	dbg, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	if version == "" {
		version = dbg.Main.Version
		for _, s := range dbg.Settings {
			if s.Key == "vcs.revision" {
				vcsHash = s.Value
			}
		}
	}
}

// printVersion prints versioning information from the package variables to the terminal.
func printVersion() {
	fmt.Printf("dsdii-assembler %v \n", version)
	if vcsHash != "" {
		fmt.Println(vcsHash)
	}
}
