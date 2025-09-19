package main

import (
	"log"
	"os"

	"github.com/MarianGheorghiu/blog_aggregator/internal/config"
)

type state struct {
	// Aplicăm un wrapper pentru Config pentru a păstra starea programului.
	// Ar fi bine să includem aici și alți clienți (ex: DB, logger custom, servicii externe)
	cfg *config.Config
}

func main() {
	// Citim configurația aplicației din fișierul JSON
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Instanțiem starea programului. 
	// ATENȚIE: `cfg` este returnat ca valoare, iar aici facem un pointer spre valoare → nu persistă modificările între reîncărcări
	programState := &state{
		cfg: &cfg,
	}

	// Inițializăm mapa de comenzi
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	// Înregistrăm prima comandă disponibilă ("login").
	cmds.register("login", handlerLogin)

	// Validăm că userul a introdus cel puțin o comandă.
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Rulăm comanda
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
