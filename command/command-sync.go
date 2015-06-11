package command

import (
	"encoding/json"
	"errors"
	"github.com/JakeClarke/supr/git"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
)

const reposDefs = ".supr"

func init() {
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "download project def and checkout latest on repos",
		Run:   sync,
	}
	rootCmd.AddCommand(syncCmd)
}

func sync(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		log.Print("Getting repo defs...")
		if err := downloadDef(args[0]); err != nil {
			log.Fatal(err)
		}
		log.Print("done!\n")
	}

	log.Println("Parsing repo defs...")
	repos, err := parseDef()
	if err != nil {
		log.Fatal(err)
	}

	for i := range repos {
		log.Printf("Cloning: %s from %s\n", repos[i].Name, repos[i].Uri)
		if err = repos[i].Clone(); err != nil {
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

func parseDef() ([]git.Repo, error) {
	r, err := os.Open(reposDefs)
	if err != nil {
		return nil, err
	}

	defer r.Close()
	decoder := json.NewDecoder(r)

	var repos []git.Repo
	err = decoder.Decode(&repos)
	return repos, err
}
