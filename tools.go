package dingo

import (
	"math/rand/v2"

	"github.com/spf13/viper"
)

func RandRange(min, max int) int {
	return rand.IntN(max-min) + min
 }

func InitConfig() error {
	viper.AddConfigPath("/home/r_rmarsu/3d_in_go")
	viper.SetConfigFile("config.yaml")
     return viper.ReadInConfig()
}

func Initialize() (*Config, error) {
	err := InitConfig()
	if err!= nil {
          return nil , err
     }

	width  := viper.GetInt("width")
	height := viper.GetInt("height")

	cellsperrun      := viper.GetInt("cellsPerRun")
	cellslifetime     := viper.GetInt("cellsLifetime")
	cellsdistance     := viper.GetInt("cellsDistance")
	timebeforehungry := viper.GetInt("cellsTimeHungry")

     config := &Config{
		Width:  width,
          Height: height,
		CellsPerRun:      cellsperrun,
          CellsLifetime:     cellslifetime,
          CellsDistance:     cellsdistance,
          CellsTimeHungry: timebeforehungry,
     }
	return config, nil
}
