package pkg

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(environment string) error {

	var err error

	if environment == "production" {
		Log, err = zap.NewProduction()
	} else {
		Log, err = zap.NewDevelopment()
	}

	if err != nil {
		return err
	}

	return nil
}
