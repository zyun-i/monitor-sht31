package main

import (
	"flag"
	"fmt"
	"os"

	i2c "github.com/d2r2/go-i2c"
	logger "github.com/d2r2/go-logger"
	sht3x "github.com/d2r2/go-sht3x"
)

var lg = logger.NewPackageLogger("main",
	logger.InfoLevel,
)

func usageExit() {
	fmt.Println("hello")
	os.Exit(0)
}

func main() {
	flag.Usage = func() { usageExit() }
	flag.Parse()
	args := flag.Args()

	var address uint8 = 0x44
	var device int = 1

	if len(args) > 0 {
		address = 0x45
	}

	defer logger.FinalizeLogger()
	i2c, err := i2c.NewI2C(address, device)
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

	fmt.Printf("weather,location=rpi,type=sht31,sid=%x temperature=%v\n", address, temp)
	fmt.Printf("weather,location=rpi,type=sht31,sid=%x humidity=%v\n", address, rh)
}
