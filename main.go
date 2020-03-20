package main

import (
	"fmt"

	i2c "github.com/d2r2/go-i2c"
	logger "github.com/d2r2/go-logger"
	sht3x "github.com/d2r2/go-sht3x"
)

var lg = logger.NewPackageLogger("main",
	logger.InfoLevel,
)

func main() {
	defer logger.FinalizeLogger()
	i2c, err := i2c.NewI2C(0x45, 1)
	if err != nil {
		lg.Fatal(err)
	}
	defer i2c.Close()

	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	logger.ChangePackageLogLevel("sht3x", logger.InfoLevel)

	sensor := sht3x.NewSHT3X()

	temp, rh, err := sensor.ReadTemperatureAndRelativeHumidity(i2c, sht3x.RepeatabilityLow)
	if err != nil {
		lg.Fatal(err)
	}

	fmt.Printf("weather,location=rpi,type=sht31,sid=45 temperature=%v\n", temp)
	fmt.Printf("weather,location=rpi,type=sht31,sid=45 humidity=%v\n", rh)
}
