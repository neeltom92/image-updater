package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/Masterminds/semver"
)

func main() {
	// Clone the repository
	repo, err := git.PlainClone("neeltom92/image-updater", false, &git.CloneOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get the latest tag
	tags, err := repo.Tags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tags.Close()

	var latestTag *object.Tag
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsTag() {
			tagObj, err := repo.TagObject(ref.Hash())
			if err != nil {
				return err
			}
			if latestTag == nil || tagObj.Tagger.When.After(latestTag.Tagger.When) {
				latestTag = tagObj
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Parse the tag to a semantic version
	version, err := semver.NewVersion(latestTag.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the version to the console
	fmt.Println(version.String())
}
