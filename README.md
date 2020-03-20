# monitor-sht31

## Installation

```sh
go get github.com/zyun-i/monitor-sht31
```

## telegraf

```conf
[[inputs.exec]]
  commands = [
    "sudo /usr/local/bin/monitor-sht31"
  ]
  timeout = "5s"
  data_format = "influx"
```

## License

[BSD-3-Clause](https://github.com/zyun-i/monitor-sht31/blob/master/LICENSE)

## Author

Isida Zyun'iti
