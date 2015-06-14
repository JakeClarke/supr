package command

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	branchgCmd = &cobra.Command{
		Use:   "branch [name] [repos]",
		Short: "Creates a branch on the managed repos and checks it out.",
		Run:   branch,
	}
)

func init() {
	rootCmd.AddCommand(branchgCmd)
}

func branch(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("Supply a branch name")
	}

	log.Println("Parsing repo defs...")
	repos, err := parseDef()
	if err != nil {
		log.Fatal(err)
	}

	if len(args) > 1 {
		repos = filterDef(repos, args[1:])
	}

	if len(repos) == 0 {
		log.Fatal("Filted all repos. Nothing to do.")
	}

	for i := range repos {
		log.Printf("Branching: %s\n", repos[i].Name)
		if err = repos[i].Branch(args[0]); err != nil {
			log.Fatal(err)
		}
	}
}
