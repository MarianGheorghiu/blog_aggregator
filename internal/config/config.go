package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Structura de bază pentru configurarea aplicației.
// ATENȚIE: momentan config-ul este simplu, dar în aplicații mai mari recomandăm împărțirea pe module.
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	// Setăm userul curent și scriem în fișier.
	// Aici am putea adăuga validare (ex: să nu fie string gol).
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	// Obținem path-ul complet la fișierul de config
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Deschidem fișierul
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Parsăm JSON-ul în structura noastră Config
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	// Obținem directorul home al userului și construim path-ul către fișierul de config.
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Suprascriem complet fișierul.
	// Ar fi de luat în calcul și o scriere atomică (ex: scriem într-un fișier temporar, apoi rename).
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
