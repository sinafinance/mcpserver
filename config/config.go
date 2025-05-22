package config

import (
	"encoding/json"
	"os"
)

// Config 表示服务器配置
type Config struct {
	Port     string `json:"port"`
	Password string `json:"password"`
	MaxConns int    `json:"max_connections"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Port:     "25575",
		Password: "",
		MaxConns: 1,
	}
}

// LoadConfig 从文件加载配置
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return DefaultConfig(), err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return DefaultConfig(), err
	}

	return &config, nil
}

// SaveConfig 保存配置到文件
func SaveConfig(filename string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
