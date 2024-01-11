package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	projectName = "go-spotify-cli"
	ServerUrl   = "http://localhost:4949"
)

type Config struct {
	ClientId        string
	ClientSecret    string
	RequestedScopes string
}

type EnvVarConfig struct {
	ClientId     string `yaml:"ClientId"`
	ClientSecret string `yaml:"ClientSecret"`
}

type ConfigService struct {
	config *Config
}

func NewConfigService() *ConfigService {
	cfg, _ := LoadConfiguration() // handle error
	return &ConfigService{config: cfg}
}

func (c *ConfigService) GetConfig() *Config {
	return c.config
}

func LoadConfiguration() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return nil, err
	}

	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+".yaml")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config EnvVarConfig
	// Unmarshal the YAML data into a Configuration struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logrus.WithError(err).Error("Error unmarshalling YAML data")
		return nil, err
	}

	return &Config{
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
	}, nil
}

func VerifyConfigExists(cfg *Config) bool {
	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+".yaml")

	// Check if the folder exists
	folderExists, err := os.Stat(folderPath)
	if err != nil || os.IsNotExist(err) || !folderExists.IsDir() {
		return false
	}

	// Check if the file exists
	_, err = os.Stat(filePath)

	if err != nil {
		return false
	}

	if len(cfg.ClientSecret) == 0 || len(cfg.ClientId) == 0 {
		return false
	}

	return true
}
