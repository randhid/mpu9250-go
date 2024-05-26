package mpu9250

import (
	"context"
	"math"

	"github.com/golang/geo/r3"
	geo "github.com/kellydunn/golang-geo"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/grpc"
	"go.viam.com/rdk/spatialmath"
)

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

func (m mpu9250) Orientation(ctx context.Context, extra map[string]interface{}) (spatialmath.Orientation, error) {
	return nil, grpc.UnimplementedError
}

func (m *mpu9250) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

// Unimplemented methods because of hardware limitations
func (m *mpu9250) Position(ctx context.Context, extra map[string]interface{}) (*geo.Point, float64, error) {
	return nil, math.NaN(), grpc.UnimplementedError
}

func (m *mpu9250) CompassHeading(ctx context.Context, extra map[string]interface{}) (float64, error) {
	return math.NaN(), grpc.UnimplementedError
}

func (m *mpu9250) LinearVelocity(ctx context.Context, extra map[string]interface{}) (r3.Vector, error) {
	return r3.Vector{}, grpc.UnimplementedError
}

func (m *mpu9250) Accuracy(ctx context.Context, extra map[string]interface{}) (*movementsensor.Accuracy, error) {
	return nil, grpc.UnimplementedError
}
