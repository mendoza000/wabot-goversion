package utils

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func PrintFlag(name, desc string) {
	color.Yellow("  --%s\n", name)
	_, err := fmt.Fprintf(os.Stderr, "\t%s\n", desc)
	if err != nil {
		return
	}
}
