// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package synclet

import (
	"context"

	"github.com/windmilleng/tilt/internal/container"
	"github.com/windmilleng/tilt/internal/docker"
)

// Injectors from wire.go:

func WireSynclet(ctx context.Context, runtime container.Runtime) (*Synclet, error) {
	clusterEnv := docker.ProvideEmptyClusterEnv()
	localEnv, err := docker.ProvideLocalEnv(ctx, clusterEnv)
	if err != nil {
		return nil, err
	}
	localClient, err := docker.ProvideLocalCli(ctx, localEnv)
	if err != nil {
		return nil, err
	}
	client := docker.ProvideLocalAsDefault(localClient)
	synclet := NewSynclet(client)
	return synclet, nil
}
