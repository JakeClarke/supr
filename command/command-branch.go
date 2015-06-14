package command

import (
	"log"

	"github.com/JakeClarke/supr/git"
	"github.com/spf13/cobra"
)

var (
	branchgCmd = &cobra.Command{
		Use:   "branch [name] [repos]",
		Short: "Creates a branch on the managed repos and checks it out.",
		Run:   genCmdHandlerFn(branch),
	}
)

func init() {
	rootCmd.AddCommand(branchgCmd)
}

func branch(cmd *cobra.Command, args []string, repos []*git.Repo) {
	if len(args) == 0 {
		log.Fatal("Supply a branch name")
	}

	for i := range repos {
		log.Printf("Branching: %s\n", repos[i].Name)
		if err := repos[i].Branch(args[0]); err != nil {
			log.Fatal(err)
		}
	}
}
