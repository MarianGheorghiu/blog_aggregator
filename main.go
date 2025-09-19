package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/MarianGheorghiu/blog_aggregator/internal/config"
	"github.com/MarianGheorghiu/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	// Aplicăm un wrapper pentru Config pentru a păstra starea programului.
	// Ar fi bine să includem aici și alți clienți (ex: DB, logger custom, servicii externe)
	cfg *config.Config
	db  *database.Queries
}

func main() {
	// Citim configurația aplicației din fișierul JSON
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// 1. deschidem conexiunea la baza de date
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()
	// 2. creăm instanța Queries din pachetul generat de sqlc
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handleReset)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
