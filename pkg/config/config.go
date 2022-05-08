package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Emitter Emitter `json:"emitter"`
	Server  Server  `json:"server"`
}

type Emitter struct {
	Connection Connection `json:"connection"`
	Protocol   string     `json:"protocol"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Connection struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func CheckEmitter(c *Config) error {
	if c.Emitter.Connection.Host == "" || c.Emitter.Connection.Port == "" {
		log.Fatalln("ERROR - EMITTER: host and port are required")
	}
	if c.Emitter.Protocol == "" {
		fmt.Println("INFO - EMITTER", "No protocol was provided in config, checking cli flags.")
	}
	return nil

}

func CheckServer(c *Config) error {
	if c.Server.Port == "" || c.Server.Host == "" {
		log.Fatalln("ERROR - SERVER: server requires port to listen on")
	}
	return nil
}

func FindConfigFile(isTest bool) (f string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("ERROR: unable to determine pwd")
		return ""
	}
	root, _, _ := strings.Cut(cwd, "Med")
	file := "med.config.json"
	if isTest {
		file = "config_test.json"
	}
	f = filepath.Join(root, "Med", file)
	return f
}

func LoadConfig(c *Config, trg string, isTest bool) error {
	p := FindConfigFile(isTest)
	file, err := os.Open(p)
	if err != nil {
		return err
	}
	fmt.Println("STATUS: successfully opened config", p)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println("STATUS: successfully unmarshalled", p)
	json.Unmarshal([]byte(bytes), c)
	if trg == "emitter" {
		CheckEmitter(c)
	}
	if trg == "server" {
		CheckServer(c)
	}
	return nil
}
