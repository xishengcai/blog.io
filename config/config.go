package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"sync"
)

// blog server config
// use: config.Config() return *BlogConfig single instance

var (
	cfg   *BlogConfig
	sOnce sync.Once
	lock  = &sync.RWMutex{}
)

// 应用服务器配置
type BlogConfig struct {
	ServerName        string      `toml:"server_name"`
	LogEnable         bool        `toml:"log_enable"`
	LogPath           string      `toml:"log_path"`
	ImaPath           string      `toml:"img_path"`
	ServerPort        int         `toml:"server_port"`
	EnvProd           bool        `toml:"enviroment_prod"`
	SecretKey         string      `toml:"secret_key"`
	DefaultClientUser string      `toml:"default_client_user"`
	DBCfg             DBConfig    `toml:"database"`
	RedisCfg          RedisConfig `toml:"redis"`
}

// 数据库配置
type DBConfig struct {
	Dtype    string `toml:"db_type"`
	Name     string `toml:"db_name"`
	Addr     string `toml:"db_addr"`
	Port     string `toml:"db_port"`
	User     string `toml:"db_user"`
	Password string `toml:"db_password"`
	Debug    bool   `toml:"db_debug"`
}

// Redis 配置
type RedisConfig struct {
	Host string `toml:"r_host"`
	Port string `toml:"r_port"`
}

// return single config instance
func Config() *BlogConfig{
	sOnce.Do(decodeConfig)
	lock.RLocker()
	defer lock.RUnlock()
	return cfg
}

// decode config with toml file
func decodeConfig() {
	path := os.Getenv("APP_CONFIG_PATH")
	if len(path) <= 0 {
		path = "./config.toml"
	}

	fp, err := filepath.Abs(path)
	if err != nil {
		panic(fmt.Errorf("Read Config Path Err: %s", err.Error()))
	}

	config := new(BlogConfig)
	if _, err := toml.DecodeFile(fp, config); err != nil {
		panic(fmt.Errorf("DecodeFile Config Error: %s", err.Error()))
	}
	lock.Lock()
	cfg = config
	lock.Unlock()
}
