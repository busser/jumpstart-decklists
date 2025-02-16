package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/busser/jumpstart-decklists/pkg/config"
	"github.com/busser/jumpstart-decklists/pkg/export"
	"github.com/busser/jumpstart-decklists/pkg/render"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: jumpstart-decklists [export|serve]\n")
}

func run() error {
	if len(os.Args) != 2 {
		printUsage()
		return errors.New("invalid number of arguments")
	}

	switch os.Args[1] {
	case "export":
		return runExport()
	case "serve":
		return runServer()
	default:
		printUsage()
		return errors.New("unknown command")
	}
}

func runExport() error {
	decklistFile, err := os.Open("decks.yml")
	if err != nil {
		return err
	}
	defer decklistFile.Close()

	decklists, err := config.ReadDecklists(decklistFile)
	if err != nil {
		return err
	}

	if err := export.DecklistsAsPNG(context.Background(), decklists, "out"); err != nil {
		return err
	}

	return nil
}

func runServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decklistFile, err := os.Open("decks.yml")
		if err != nil {
			log.Printf("error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer decklistFile.Close()

		decklists, err := config.ReadDecklists(decklistFile)
		if err != nil {
			log.Printf("error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		html, err := render.DecklistsAsWebPage(decklists)
		if err != nil {
			log.Printf("error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(html)
	})

	log.Println("serving on http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}
