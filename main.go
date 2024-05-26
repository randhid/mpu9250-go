package main

import (
	"context"
	"mpu9250/mpu9250"

	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("Universal Robotics Arm Go Module"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
	// instantiates the module itself
	mpu9250module, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	// Models and APIs add helpers to the registry during their init().
	// They can then be added to the module here.s
	if err = mpu9250module.AddModelFromRegistry(ctx, movementsensor.API, mpu9250.Model); err != nil {
		return err
	}
	// Each module runs as its own process
	err = mpu9250module.Start(ctx)
	logger.Warn("starting module")
	defer mpu9250module.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
