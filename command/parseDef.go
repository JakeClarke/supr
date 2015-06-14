package command

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JakeClarke/supr/git"
	"github.com/spf13/cobra"
)

const reposDefs = ".supr"

// Filters repos by names provided.
// If 0 names are provided then repos is returned.
func filterDef(repos []*git.Repo, names []string) (res []*git.Repo) {
	if len(names) == 0 {
		res = repos
		return
	}

	for n := range names {
		for r := range repos {
			if repos[r].Name == names[n] {
				res = append(res, repos[r])
			}
		}
	}

	return
}

// Parses the repos def file.
func parseDef() ([]*git.Repo, error) {
	r, err := os.Open(reposDefs)
	if err != nil {
		return nil, err
	}

	defer r.Close()
	decoder := json.NewDecoder(r)

	var repos []*git.Repo
	err = decoder.Decode(&repos)
	return repos, err
}

// Generates a cmd handler function with repos loaded and prefiltered. With n Arguments
func genCmdHandlerFnN(fn func(*cobra.Command, []string, []*git.Repo), n int) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		log.Println("Parsing repo defs...")
		repos, err := parseDef()
		if err != nil {
			log.Fatal(err)
		}

		if len(args) > n {
			// filter and consume command line args
			repos = filterDef(repos, args[n:])
			args = args[:n]

			if len(repos) == 0 {
				log.Fatal("Filted all repos. Nothing to do.")
			}
		}

		fn(cmd, args, repos)
	}
}

// Generates a cmd handler function with repos loaded and prefiltered. With 1 argument.
func genCmdHandlerFn(fn func(*cobra.Command, []string, []*git.Repo)) func(*cobra.Command, []string) {
	return genCmdHandlerFnN(fn, 1)
}
