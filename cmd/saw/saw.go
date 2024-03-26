package saw

import (
	"github.com/MR5356/wtf/pkg/saw"
	"github.com/spf13/cobra"
)

var (
	command = "sh"
	port    int
)

func NewSawCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "saw",
		Short: "Share the terminal with your newbie colleagues",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				command = args[0]
			}
			return saw.Terminal(command, port)
		},
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "port to run")

	return cmd
}
