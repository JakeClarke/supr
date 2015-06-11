package command

import (
	"encoding/json"
	"os"

	"github.com/JakeClarke/supr/git"
)

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
