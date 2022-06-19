package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var server *config

func init() {
	server = new(config)
}

type config struct {
	ServerPort    uint32 `toml:"server_port"`
	Debug         bool
	ConsoleMode   bool `toml:"console_mode"`
	Cors          bool
	ShowSQL       bool             `toml:"showSQL"`
	Pprof         string           `toml:"pprof"`
	ServiceConfig string           `toml:"service_config"`
	Nodes         map[string]*Node `toml:"nodes"`
	Database      *Database
}

type Database struct {
	Host   string
	Port   uint32
	User   string
	Pwd    string
	DBName string `toml:"db_name"`
	Param  string
}

type RedisConfig struct {
	Port uint32
	Pwd  string
}

type Node struct {
	Rpc string `toml:"rpc"`
	//Scan bool   `toml:"scan"`
}

func Parse(path string) {
	_, err := toml.DecodeFile(path, server)
	if err != nil {
		panic(err)
	}
}

func IsConsoleMode() bool {
	return server.ConsoleMode
}

func DatabaseStr() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", server.Database.User, server.Database.Pwd,
		server.Database.Host, server.Database.Port, server.Database.DBName, server.Database.Param)
}

func Port() uint32 {
	return server.ServerPort
}

func IsDebug() bool {
	return server.Debug
}

func IsUseCors() bool {
	return server.Cors
}

func IsShowSQL() bool {
	return server.ShowSQL
}

func ServiceConfig() string {
	return server.ServiceConfig
}

func AllChainNode() map[string]*Node {
	return server.Nodes
}

func Pprof() string {
	return server.Pprof
}
