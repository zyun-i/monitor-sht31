package main

import (
	"flag"
	"fmt"

	i2c "github.com/d2r2/go-i2c"
	logger "github.com/d2r2/go-logger"
	sht3x "github.com/d2r2/go-sht3x"
)

var lg = logger.NewPackageLogger("main",
	logger.InfoLevel,
)

func main() {
	var device = flag.Int("d", 1, "I2C device")
	var address = flag.Int("a", 0x44, "I2C address")
	flag.Parse()

	defer logger.FinalizeLogger()
	i2c, err := i2c.NewI2C(uint8(*address), *device)
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

	fmt.Printf("sensor_sht31_temperature{address=\"%x\"} %v\n", *address, temp)
	fmt.Printf("sensor_sht31_humidity{address=\"%x\"} %v\n", *address, rh)
}
