# monitor-sht31

Monitoring SHT31 i2c sensor.

## Exec

```sh
$ sudo /usr/local/sbin/monitor-sht31 -a 0x45 -d 1 -l example
weather,location=example,type=sht31,sid=45 temperature=27.97
weather,location=example,type=sht31,sid=45 humidity=66.4
```

## telegraf

```conf
[[inputs.exec]]
  commands = [
    "sudo /usr/local/bin/monitor-sht31",
    "sudo /usr/local/bin/monitor-sht31 -a 0x45"
  ]
  timeout = "5s"
  data_format = "influx"
```

## License

[BSD-3-Clause](https://github.com/zyun-i/monitor-sht31/blob/master/LICENSE)

## Author

Isida Zyun'iti
