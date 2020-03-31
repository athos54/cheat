package installer

import (
	"fmt"
	"os"
	"os/exec"
)

const cloneURL = "https://github.com/cheat/cheatsheets.git"

// Clone clones the community cheatsheets
func Clone(path string) error {

	// perform the clone in a shell
	cmd := exec.Command("git", "clone", cloneURL, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clone cheatsheets: %v", err)
	}

	return nil
}

const clonePOURL = "https://gitlab.com/athos.oc/po-cheatsheets"

// Clone clones the community cheatsheets
func ClonePO(path string) error {

	// perform the clone in a shell
	cmd := exec.Command("git", "clone", clonePOURL, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clone cheatsheets: %v", err)
	}

	return nil
}

func PullPO(path string) error {
	fmt.Println("path", path)
	// perform the clone in a shell
	a := "--git-dir=/home/athos/.config/cheat/cheatsheets/po/.git"
	b := "--work-tree=/home/athos/.config/cheat/cheatsheets/po"
	cmd := exec.Command("git", a, b, "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to pull cheatsheets: %v", err)
	}

	return nil
}

func PushPO(path string) error {
	fmt.Println("PushPO", path)
	// perform the clone in a shell
	a := "--git-dir=/home/athos/.config/cheat/cheatsheets/po/.git"
	b := "--work-tree=/home/athos/.config/cheat/cheatsheets/po"
	cmd := exec.Command("git", a, b, "add", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to add git: %v", err)
	}

	cmd = exec.Command("git", a, b, "commit", "-m", "\"cheat-auto-commit\"")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to commit git: %v", err)
	}

	cmd = exec.Command("git", a, b, "push", "origin", "master")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to push cheatsheets: %v", err)
	}

	return nil
}
