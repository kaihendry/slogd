# slog duration

No need for heavyweight tracing tools like X-ray, Jaeger or Zipkin. "slogd"
measures the time for a function to return an error object.

Inspired by @[tj](https://github.com/tj)'s [Apex logs's Trace function](https://pkg.go.dev/github.com/apex/log#Trace)

- https://github.com/apex/log/blob/master/_examples/trace/trace.go
- https://github.com/apex/log/blob/8da83152b5d6177b4bfe3d12810a5afd25355170/entry.go#L142-L158

## [slogd example](https://github.com/kaihendry/slogfest)

Sample code:

    slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout)))
    url := "https://httpbin.org/delay/2"
    var err error
    // Caveat: Only one defer slogd per function!
    defer slogd.New("fetching", "url", url).Stop(&err)
    _, err = http.Get(url)

Sample output:

    time=2023-02-10T20:21:46.564+08:00 level=INFO msg=fetching url=https://httpbin.org/delay/2 duration=3.716787544s
