package main

import (
	"fmt"
	"log"

	"github.com/MarianGheorghiu/blog_aggregator/internal/config"
)

func main() {
	// Citește fișierul de config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	// Setează user-ul
	err = cfg.SetUser("marian")
	if err != nil {
		log.Fatalf("Error writing: %v", err)
	}

	// Citește din nou config-ul pentru a verifica modificarea
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading: %v", err)
	}

	// Afișează structura Config în terminal
	fmt.Printf("Final Config: %+v\n", cfg)
}
