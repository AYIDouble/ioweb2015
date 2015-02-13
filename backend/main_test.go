package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.Dir = "app"
	config.Env = "dev"
	config.Prefix = "/myprefix"
	cache = newMemoryCache()
	os.Exit(m.Run())
}

func overrideEnv(env string) func() {
	orig := config.Env
	config.Env = env
	return func() { config.Env = orig }
}
