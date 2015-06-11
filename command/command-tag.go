package command

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	tagCmd = &cobra.Command{
		Use:   "tag [name] [repos]",
		Short: "tag the currently checked out branch on the child repos",
		Run:   tag,
	}

	message *string
)

func init() {
	message = tagCmd.Flags().String("msg", "", "Message that you want to be attached to the tag. If not set, tag name is used")
	rootCmd.AddCommand(tagCmd)
}

func tag(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("Supply a tag name")
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
		log.Printf("Tagging: %s\n", repos[i].Name)
		if err = repos[i].Tag(args[0], *message); err != nil {
			log.Fatal(err)
		}
	}
}
