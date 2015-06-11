package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const gitExec = "git"

type Repo struct {
	Uri           string
	Name          string
	DefaultBranch string
}

func (r *Repo) Clone() error {
	params := []string{"clone"}
	if len(r.DefaultBranch) > 0 {
		params = append(params, "-b", r.DefaultBranch, r.Uri, r.Name)
	} else {
		params = append(params, r.Uri, r.Name)
	}

	cmd := exec.Command(gitExec, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to run git: %v", err)
	}

	return nil
}

func (r *Repo) Branch(branch string) error {
	cmd := exec.Command(gitExec, "checkout", "-b", branch)
	cmd.Dir = r.Name

	return cmd.Run()
}

func (r *Repo) Tag(tag, message string) error {
	param := []string{"tag", "-a"}
	if len(message) > 0 {
		param = append(param, "-m", message)
	} else {
		param = append(param, "-m", tag)
	}

	if len(tag) == 0 {
		return errors.New("Need a tag.")
	}
	param = append(param, tag)

	cmd := exec.Command(gitExec, param...)
	cmd.Dir = r.Name
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("Failed to tag %s with %s, reason: %s(%v)", r.Name, tag, out, err)
	}
	return nil
}

func (r *Repo) Checkout(branch string) error {
	cmd := exec.Command(gitExec, "checkout", branch)
	cmd.Dir = r.Name

	return cmd.Run()
}
