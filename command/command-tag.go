package command

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	tagCmd = &cobra.Command{
		Use:   "tag",
		Short: "tag the currently checked out branch on the child repos",
		Run:   tag,
	}
)

func init() {
	rootCmd.AddCommand(tagCmd)
}

func tag(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("Supply a tag name")
	}
	var message string
	if len(args) > 1 {
		message = args[1]
	}

	log.Println("Parsing repo defs...")
	repos, err := parseDef()
	if err != nil {
		log.Fatal(err)
	}

	for i := range repos {
		log.Printf("Tagging: %s\n", repos[i].Name)
		if err = repos[i].Tag(args[0], message); err != nil {
			log.Fatal(err)
		}
	}
}