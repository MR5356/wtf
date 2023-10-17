package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"math/rand"
	"time"
)

func Commit(msg string) error {
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

	author := authors[rand.Intn(len(authors))]

	commit, err := w.Commit(msg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  author.Author.Name,
			Email: author.Author.Email,
			When:  time.Now(),
		},
	})

	fmt.Printf("commit with name %s and email %s\n", author.Author.Name, author.Author.Email)

	_, err = r.CommitObject(commit)
	return err
}
