package shared

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml"
)

// Config represents the application configuration
type Config struct {
	Appearance struct {
		SelectedTheme string `toml:"selected_theme"`
		FontSize      int    `toml:"font_size"`
		Theme         struct {
			Light ThemeColors `toml:"light"`
			Dark  ThemeColors `toml:"dark"`
		} `toml:"theme"`
	} `toml:"appearance"`
}

// ThemeColors represents the colors for a theme
type ThemeColors struct {
	Background string `toml:"background"`
	Foreground string `toml:"foreground"`
	Primary    string `toml:"primary"`
	Secondary  string `toml:"secondary"`
	Tertiary   string `toml:"tertiary"`
}

// LoadConfig loads the configuration from the config file
func LoadConfig() (*Config, error) {
	logger := GetLogger()
	configDir := filepath.Join(xdg.ConfigHome, "MultiCrypt")
	configFile := filepath.Join(configDir, "config.toml")

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		logger.Info("Config file not found, creating default config")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, err
		}
		if err := createDefaultConfig(configFile); err != nil {
			return nil, err
		}
	}

	config := &Config{}
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

// SaveConfig saves the configuration to the config file
func SaveConfig(config *Config) error {
	configDir := filepath.Join(xdg.ConfigHome, "MultiCrypt")
	configFile := filepath.Join(configDir, "config.toml")

	data, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, data, 0644)
}

// createDefaultConfig creates a default configuration file
func createDefaultConfig(configFile string) error {
	defaultConfig := &Config{}
	defaultConfig.Appearance.SelectedTheme = "dark"
	defaultConfig.Appearance.FontSize = 14
	defaultConfig.Appearance.Theme.Light = ThemeColors{
		Background: "F0F4F8",
		Foreground: "1A202C",
		Primary:    "3182CE",
		Secondary:  "38A169",
		Tertiary:   "E53E3E",
	}
	defaultConfig.Appearance.Theme.Dark = ThemeColors{
		Background: "1A202C",
		Foreground: "F0F4F8",
		Primary:    "4299E1",
		Secondary:  "48BB78",
		Tertiary:   "F56565",
	}

	data, err := toml.Marshal(defaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, data, 0644)
}
