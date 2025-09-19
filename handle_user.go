package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// Validăm input-ul. Aici cerem fix 1 argument: numele userului.
	// Dacă vrem să extindem CLI-ul, putem implementa și `flags` sau parsing mai robust (ex: cobra/urfave/cli).
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	// Actualizăm userul curent în config și scriem în fișierul JSON.
	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	// Feedback pentru utilizator.
	fmt.Println("User switched successfully!")
	return nil
}
