package command

import (
	"log"

	"github.com/JakeClarke/supr/git"
	"github.com/spf13/cobra"
)

var (
	tagCmd = &cobra.Command{
		Use:   "tag [name] [repos]",
		Short: "tag the currently checked out branch on the child repos",
		Run:   genCmdHandlerFn(tag),
	}

	message *string
)

func init() {
	message = tagCmd.Flags().StringP("msg", "m", "", "Message that you want to be attached to the tag. If not set, tag name is used")
	rootCmd.AddCommand(tagCmd)
}

func tag(cmd *cobra.Command, args []string, repos []*git.Repo) {
	if len(args) == 0 {
		log.Fatal("Supply a tag name")
	}

	for i := range repos {
		log.Printf("Tagging: %s\n", repos[i].Name)
		if err := repos[i].Tag(args[0], *message); err != nil {
			log.Fatal(err)
		}
	}
}
