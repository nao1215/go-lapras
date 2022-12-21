// Package main is go-lapras main pacakage
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/nao1215/go-lapras/go-lapras"
)

func main() {
	if len(os.Args) <= 1 {
		usage(os.Stdout)
		os.Exit(0)
	}

	lapras := lapras.NewLapras()
	person, err := lapras.GetPerson(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("Name             :%s\n", person.Name)
	fmt.Printf("Engineering Score:%v\n", person.EScore)
	fmt.Printf("Business Score   :%v\n", person.EScore)
	fmt.Printf("Score Influense  :%v\n", person.EScore)
	fmt.Println()

	for _, v := range person.GithubRepositories {
		if v.IsOwner && v.StargazersCount > 10 {
			fmt.Printf("Title       :%s\n", v.Title)
			fmt.Printf("Description :%s\n", v.Description)
			fmt.Printf("Stargazers  :%v\n", v.StargazersCount)
			fmt.Println()
		}
	}
}

// no help, no version :)
func usage(w io.Writer) {
	fmt.Fprintln(w, "lapras command is sample program for LAPRAS Inc. API (Version 0.0.1)")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "[Usage]")
	fmt.Fprintln(w, " lapras SHARE_ID")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "[License]")
	fmt.Fprintln(w, " MIT License - Copyright (c) 2022 CHIKAMATSU Naohiro")
}
