package main

import (
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use: "japlay",
}

func main() {
	cmdRoot.Execute()
}
