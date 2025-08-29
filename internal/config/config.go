package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Database Database `json:"database"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}

// Init parses config file named "config.json" from same directory as executable binary
// and return Config struct instance.
func Init() (Config, error) {
	exec, err := os.Executable()
	if err != nil {
		return Config{}, fmt.Errorf("os.Executable: %w", err)
	}

	dir, _ := filepath.Split(exec)
	configPath := filepath.Join(dir, "config.json")
	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("os.Open: %w", err)

	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("decoder.Decode: %w", err)

	}

	return cfg, err
}
