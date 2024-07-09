//go:build !linux

package mpu9250

import (
	"context"

	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)


func init() {
	resource.RegisterComponent(movementsensor.API, Model, resource.Registration[movementsensor.MovementSensor, *resource.NoNativeConfig]{
		Constructor: newMpu9250,
	})
}

func newMpu9250(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	movementsensor.MovementSensor, error) {
	m := &mpu9250{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	return m, nil
}

type mpu9250 struct {
	resource.Named
	resource.AlwaysRebuild
	resource.TriviallyCloseable
	logger logging.Logger
}
