package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	git2go "gopkg.in/libgit2/git2go.v22"
)

func main() {

	var r *git2go.Repository
	var err error

	repo_log := "./test_repo"

	r, err = git2go.OpenRepository(repo_log)
	if err != nil {

		r, err = git2go.InitRepository(repo_log, false)
		if err != nil {
			log.Fatal(err)
		}
	}

	idx, err := r.Index()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(repo_log+"/happy.txt", []byte(time.Now().String()+"\n"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = idx.AddByPath("happy.txt")
	if err != nil {
		log.Fatal(err)
	}

	idx_id, err := idx.WriteTree()
	if err != nil {
		log.Fatal(err)
	}

	err = idx.Write()
	if err != nil {
		log.Fatal(err)
	}

	tree, err := r.LookupTree(idx_id)
	if err != nil {
		log.Fatal(err)
	}

	head, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}

	commit, err := r.LookupCommit(head.Target())
	if err != nil {
		log.Println("lookup commit")
		log.Fatal(err)
	}

	sig := &git2go.Signature{Name: "Rob Cameron", Email: "rwcameron@gmail.com", When: time.Now()}

	id, err := r.CreateCommit("HEAD", sig, sig, "Updating happiness", tree, commit)
	if err != nil {
		log.Println("commit")
		log.Fatal(err)
	}

	log.Println(id)
}
