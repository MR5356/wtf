package commit

import (
	"github.com/MR5356/wtf/pkg/git"
	"github.com/spf13/cobra"
)

var (
	message string
)

func NewGitCommitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "commit",
		ValidArgs: []string{"message"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return git.Commit(message)
		},
	}

	cmd.PersistentFlags().StringVarP(&message, "message", "m", "", "commit message")
	return cmd
}
