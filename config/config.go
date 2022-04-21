package config

import (
	"Kasir_Test/manager"
	"Kasir_Test/util"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type ApiConfig struct {
	Url string
}

type Manager struct {
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

type DbConfig struct {
	Host     string
	User     string
	Port     string
	Password string
	Name     string
}

type Config struct {
	Manager
	DbConfig
	ApiConfig
	LogLevel string
}

func (c Config) Configuration(path, fileName string) Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Name:     os.Getenv("MYSQL_DBNAME"),
	}

	c.ApiConfig = ApiConfig{Url: v.GetString("api.url")}

	c.LogLevel = v.GetString("api.log_level")

	return c
}

func NewConfig(path, name string) Config {
	cfg := Config{}
	cfg = cfg.Configuration(path, name)
	util.NewLog(cfg.LogLevel)

	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	cfg.InfraManager = manager.NewInfraManager(dataSourceName)
	cfg.RepoManager = manager.NewRepoManager(cfg.InfraManager)
	cfg.UseCaseManager = manager.NewUseCaseManager(cfg.RepoManager)

	return cfg
}
