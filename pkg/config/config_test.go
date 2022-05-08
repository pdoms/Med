package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindConfigFile(t *testing.T) {
	cwd, _ := os.Getwd()
	root, _, _ := strings.Cut(cwd, "Med")
	want := filepath.Join(root, "Med", "config_test.json")
	got := FindConfigFile(true)
	if want != got {
		t.Errorf("got %q want %q for host", got, want)

	}
}

func TestLoadConfig(t *testing.T) {
	var conf Config
	LoadConfig(&conf, "emitter", true)
	wantHost := "127.0.0.1"
	wantPort := "8080"
	if wantHost != conf.Emitter.Connection.Host {
		t.Errorf("got %q want %q for host", conf.Emitter.Connection.Host, wantHost)
	}
	if wantPort != conf.Emitter.Connection.Port {
		t.Errorf("got %q want %q for host", conf.Emitter.Connection.Port, wantHost)
	}
}
