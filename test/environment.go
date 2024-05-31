package test

import (
	"context"
	"gin-rush-template/config"
	"gin-rush-template/internal/global/database"
	"gin-rush-template/tools"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)
import (
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

const (
	EnvFileName   = "docker-compose.env.yml"
	ConfigFilName = "config.yaml"
)

func SetupEnvironment(t *testing.T) {
	compose, err := tc.NewDockerCompose(tools.SearchFile(EnvFileName))
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal))
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	require.NoError(t,
		compose.WaitForService("mysql", wait.ForLog("port: 3306  MySQL Community Server")).Up(ctx, tc.Wait(true)),
	)

	config.Read(tools.SearchFile(ConfigFilName))
	database.Init()
}
