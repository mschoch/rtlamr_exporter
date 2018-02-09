# rtlamr_exporter

A [Prometheus](https://prometheus.io/) exporter for data collected by [rtlamr](https://github.com/bemasher/rtlamr).

## Usage

Pipe the output of rtlamr that you want to expose:

```
rtlamr -filterid=<yourdevice> -msgtype=scm -format=json -centerfreq=915000000 | rtlamr_exporter
```

By default this presents the data for prometheus on port 9415, you can override this with the `-addr` flag.
