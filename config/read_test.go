package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadFromEnv(t *testing.T) {
	os.Setenv("HOST", "3333")
	os.Setenv("PORT", "env-user")
	os.Setenv("MODE", "env-pass")
	defer os.Clearenv()

	Init("../config.example.yaml")

	cfg := Get()

	assert.Equal(t, "3333", cfg.Host)
	assert.Equal(t, "env-user", cfg.Port)
	assert.Equal(t, Mode("env-pass"), cfg.Mode)

	t.Log(cfg)
}

func TestLoadFromFiles(t *testing.T) {
	Init("../config.example.yaml")

	cfg := Get()

	assert.Equal(t, "0.0.0.0", cfg.Host)
	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, ModeDebug, cfg.Mode)

	t.Log(cfg.OTel)
	t.Log(cfg.Mysql)
}
