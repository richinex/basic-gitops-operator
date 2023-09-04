package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
)

func main() {
	// Define the sync interval duration.
	timerInterval := 5 * time.Second

	// URL to the GitOps repository.
	gitopsRepo := "https://github.com/richinex/basic-gitops-operator.git"

	// Local directory where the repo will be cloned.
	localPath := "tmp/"

	// Path within the cloned repo containing manifests to apply.
	pathToApply := "gitops-operator-config"

	// Start an infinite loop for syncing the repo and applying manifests.
	for {
		// Start synchronizing the repository.
		fmt.Println("start repo sync")
		err := syncRepo(gitopsRepo, localPath)
		if err != nil {
			// If there's an error in syncing, report it and stop the loop.
			fmt.Printf("repo sync error: %s", err)
			return
		}

		// Start applying the manifests from the repository.
		fmt.Println("start manifests apply")
		err = applyManifestsClient(path.Join(localPath, pathToApply))
		if err != nil {
			// If there's an error in applying manifests, report it.
			fmt.Printf("manifests apply error: %s", err)
		}

		// Set a timer to wait before the next sync.
		syncTimer := time.NewTimer(timerInterval)
		fmt.Printf("\n next sync in %s \n", timerInterval)

		// Block and wait for the timer to expire.
		<-syncTimer.C
	}
}

// syncRepo tries to clone a repository from the given URL into the local path.
// If the repository already exists, it pulls the latest changes.
func syncRepo(repoUrl, localPath string) error {
	// Attempt to clone the repository.
	_, err := git.PlainClone(localPath, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout, // Display progress to standard output.
	})

	// If the repository already exists locally, try to pull the latest changes.
	if err == git.ErrRepositoryAlreadyExists {
		// Open the existing repository.
		repo, err := git.PlainOpen(localPath)
		if err != nil {
			return err
		}

		// Get the working tree of the repo.
		w, err := repo.Worktree()
		if err != nil {
			return err
		}

		// Pull the latest changes.
		err = w.Pull(&git.PullOptions{
			RemoteName: "origin",
			Progress:   os.Stdout, // Display progress to standard output.
		})

		// The library returns an "Already up to date" error if there's nothing to pull.
		// In our case, we don't consider it an error.
		if err == git.NoErrAlreadyUpToDate {
			return nil
		}
		return err
	}
	return err
}

// applyManifestsClient applies Kubernetes manifests from the given local path using kubectl.
func applyManifestsClient(localPath string) error {
	// Get the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Construct the kubectl command to apply the manifests.
	cmd := exec.Command("kubectl", "apply", "-f", path.Join(dir, localPath))
	cmd.Stdout = os.Stdout // Forward stdout of the command.
	cmd.Stderr = os.Stderr // Forward stderr of the command.

	// Run the command.
	err = cmd.Run()
	return err
}
