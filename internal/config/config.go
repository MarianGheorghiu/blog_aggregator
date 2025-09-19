package config

import (
	"encoding/json" // Pentru citirea și scrierea JSON
	"os"            // Pentru manipularea fișierelor
	"path/filepath" // Pentru construirea portabilă a căilor
)

// Numele fișierului de configurare
const configFileName = ".gatorconfig.json"

// Structura care reprezintă config-ul
type Config struct {
	DBURL           string `json:"db_url"`            // Cheia JSON "db_url" → DBURL
	CurrentUserName string `json:"current_user_name"` // Cheia JSON "current_user_name" → CurrentUserName
}

// SetUser modifică CurrentUserName și scrie automat config-ul pe disc
func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName // Modifică structura în memorie
	return write(*cfg)             // Scrie imediat modificarea pe disc
}

// Read citește config-ul din fișierul JSON și îl decodează într-o structură Config
func Read() (Config, error) {
	fullPath, err := getConfigFilePath() // Obține calea completă a fișierului de config
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath) // Deschide fișierul pentru citire
	if err != nil {
		return Config{}, err
	}
	defer file.Close() // Se asigură că fișierul se închide la final

	decoder := json.NewDecoder(file) // Creează un decoder JSON pentru fișier
	cfg := Config{}                  // Inițializează structura Config
	err = decoder.Decode(&cfg)       // Decodează JSON-ul în structura Config
	if err != nil {
		return Config{}, err
	}

	return cfg, nil // Returnează config-ul citit
}

// getConfigFilePath construiește calea completă a fișierului de config
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir() // Obține directorul home al utilizatorului
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName) // Concatenează directorul home + numele fișierului
	return fullPath, nil
}

// write scrie structura Config în fișierul JSON
func write(cfg Config) error {
	fullPath, err := getConfigFilePath() // Obține calea completă a fișierului
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath) // Creează fișierul sau îl suprascrie dacă există
	if err != nil {
		return err
	}
	defer file.Close() // Închide fișierul la final

	encoder := json.NewEncoder(file) // Creează un encoder JSON pentru fișier
	err = encoder.Encode(cfg)        // Encodează structura Config direct în fișier
	if err != nil {
		return err
	}

	return nil // Returnează nil dacă scrierea a fost cu succes
}
