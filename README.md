# monitor-sht31

## telegraf

```conf
[[inputs.exec]]
  commands = [
    "sudo /usr/local/bin/monitor-sht31"
  ]
  timeout = "5s"
  data_format = "influx"
```
