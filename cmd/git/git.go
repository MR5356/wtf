package git

import (
	"github.com/MR5356/wtf/cmd/git/commit"
	"github.com/spf13/cobra"
)

func NewGitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git",
		Short: "Very cool alternative to git",
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmd.AddCommand(commit.NewGitCommitCommand())
	return cmd
}
