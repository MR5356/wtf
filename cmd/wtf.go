package main

import (
	"github.com/MR5356/wtf/cmd/git"
	"github.com/MR5356/wtf/pkg/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "wtf",
	Version: version.Version,
}

func init() {
	rootCmd.AddCommand(git.NewGitCommand())
}
func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
