package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const gitExec = "git"

type Repo struct {
	Uri  string
	Name string
}

func (r *Repo) Clone() error {
	cmd := exec.Command(gitExec, "clone", r.Uri, r.Name)
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
	}

	if len(tag) == 0 {
		return errors.New("Need a tag.")
	}
	param = append(param, tag)

	cmd := exec.Command(gitExec, param...)
	cmd.Dir = r.Name
	return cmd.Run()
}

func (r *Repo) Checkout(branch string) error {
	cmd := exec.Command(gitExec, "checkout", branch)
	cmd.Dir = r.Name

	return cmd.Run()
}
