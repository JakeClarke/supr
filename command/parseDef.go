package command

import (
	"encoding/json"
	"os"

	"github.com/JakeClarke/supr/git"
)

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
