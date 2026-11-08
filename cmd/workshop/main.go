// Command workshop is OpenSSL Workshop. Subcommands arrive with the phases in docs/PHASES.md.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "workshop: scaffold; nothing runs yet. See docs/PHASES.md.")
	if len(os.Args) > 1 {
		os.Exit(2)
	}
}
