package main

import (
	"fmt"
	"os"

	"github.com/cloudlift/lift/pkg/commands"
)

func main() {
	rootCmd := commands.RootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
