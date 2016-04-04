package main

import (
	"log"

	git2go "github.com/libgit2/git2go"
)

func main() {

	r, err := git2go.InitRepository("./test_repo", false)
	if err != nil {
		log.Fatal(err)
	}

	c, err := r.Config()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(c)
}
