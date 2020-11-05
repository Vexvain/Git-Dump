package cmd

import (
	"fmt"
	"github.com/Vexvain/Git-Dump/pkg/Git-Dump"
	"github.com/spf13/cobra"
	"os"
)

var force bool
var list bool
var rootCmd = &cobra.Command{
	Use:   "git-dump",
	Short: "An extremely fast tool that grabs sources from exposed .git folders",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dir string
		if len(args) >= 2 {
			dir = args[1]
		}
		if list {
			if err := git-dump.CloneList(args[0], dir, force); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			if err := git-dump.Clone(args[0], dir, force); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "overrides DIR if it already exists")
	rootCmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "allows you to supply the name of a file containing a list of domain names instead of just one domain")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
