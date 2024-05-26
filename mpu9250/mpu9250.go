package mpu9250

import (
	"context"

	"github.com/golang/geo/r3"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

var (
	mpuFamily = resource.NewModelFamily("rand", "mpu9250")
	Model     = mpuFamily.WithModel("mpu9250")
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

func (m *mpu9250) Properties(ctx context.Context, extra map[string]interface{}) (
	*movementsensor.Properties, error,
) {
	return &movementsensor.Properties{
		LinearAccelerationSupported: true,
		AngularVelocitySupported:    true,
		OrientationSupported:        true,
	}, nil
}

func (m *mpu9250) AngularVelocity(ctx context.Context, extra map[string]interface{}) (spatialmath.AngularVelocity, error) {

	return spatialmath.AngularVelocity{}, nil
}

func (m *mpu9250) LinearAcceleration(ctx context.Context, extra map[string]interface{}) (r3.Vector, error) {
	return r3.Vector{}, nil
}

func (m *mpu9250) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}
