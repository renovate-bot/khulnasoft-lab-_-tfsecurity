package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/khulnasoft-lab/tfsecurity/internal/app/tfsecurity/cmd"
)

const transitionMsg = `
======================================================
tfsecurity is joining the Tunnel family

tfsecurity will continue to remain available 
for the time being, although our engineering 
attention will be directed at Tunnel going forward.

You can read more here: 
https://github.com/khulnasoft-lab/tfsecurity/discussions/56
======================================================
`

func main() {
	fmt.Print(transitionMsg)
	if err := cmd.Root().Execute(); err != nil {
		if err.Error() != "" {
			fmt.Printf("Error: %s\n", err)
		}
		var exitErr *cmd.ExitCodeError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.Code())
		}
		os.Exit(1)
	}
}
