package main

import (
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use: "ja_player",
}

func main() {
	cmdRoot.Execute()
}
