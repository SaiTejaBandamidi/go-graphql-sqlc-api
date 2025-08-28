package config

import "testing"

func TestLoadEnv_NoPanic(t *testing.T) {
    // Just ensure LoadEnv can be called without panicking when .env is missing.
    LoadEnv()
}
