package command

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const reposDefs = ".supr"

var (
	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "download project def and checkout latest on repos",
		Run:   sync,
	}

	uri *string
)

func init() {
	uri = syncCmd.Flags().StringP("uri", "u", "", "URI to copy from. Uses local copy if not provided.")
	rootCmd.AddCommand(syncCmd)
}

func sync(cmd *cobra.Command, args []string) {
	if len(*uri) > 0 {
		log.Print("Getting repo defs...")
		if err := downloadDef(*uri); err != nil {
			log.Fatal(err)
		}
		log.Print("done!\n")
	}

	log.Println("Parsing repo defs...")
	repos, err := parseDef()
	if err != nil {
		log.Fatal(err)
	}

	if len(args) > 0 {
		repos = filterDef(repos, args)
	}

	if len(repos) == 0 {
		log.Fatal("Filted all repos. Nothing to do.")
	}

	for i := range repos {
		log.Printf("Syncing: %s from %s. Default branch: %s\n", repos[i].Name, repos[i].Uri, repos[i].DefaultBranch)
		if err = repos[i].Sync(); err != nil {
			log.Fatal(err)
		}
	}
}

func downloadDef(uri string) error {
	if len(uri) == 0 {
		return errors.New("No uri provided.")
	}

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	r, err := os.Create(reposDefs)
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = io.Copy(r, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
