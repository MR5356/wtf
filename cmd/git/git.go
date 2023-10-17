package git

import (
	"github.com/MR5356/wtf/pkg/git"
	"github.com/spf13/cobra"
)

var exampleForPushCmd = `
  wtf gc "new commit"
`

func NewGitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "git",
		Example: exampleForPushCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return cmd.Help()
			}
			return git.Commit(args[0])
		},
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	return cmd
}
