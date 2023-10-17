package commit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"math/rand"
)

func Test(file string) error {
	r, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	objects, err := r.CommitObjects()
	if err != nil {
		return err
	}

	authors := make([]*object.Commit, 0)

	err = objects.ForEach(func(c *object.Commit) error {
		authors = append(authors, c)
		return nil
	})
	if err != nil {
		return err
	}

	_, err = w.Add(file)
	if err != nil {
		return err
	}

	fmt.Println("git status --porcelain")
	status, err := w.Status()
	if err != nil {
		return err
	}

	fmt.Println(status)

	fmt.Println(authors[rand.Intn(len(authors))])
	return nil
}
