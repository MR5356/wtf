package git

import (
	"github.com/MR5356/wtf/pkg/git"
	"github.com/spf13/cobra"
)

func NewGitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "git",
		RunE: func(cmd *cobra.Command, args []string) error {
			return git.Commit(args[0])
		},
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	return cmd
}
