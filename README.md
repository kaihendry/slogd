# Go trace

No need for heavyweight tracing tools like X-ray, Jaeger or Zipkin. Just use the Go trace package, a simple function to log a function that returns an error object and measure the duration.

Inspired by @[tj](https://github.com/tj)'s [Apex logs's Trace function](https://pkg.go.dev/github.com/apex/log#Trace)

- https://github.com/apex/log/blob/master/_examples/trace/trace.go
- https://github.com/apex/log/blob/8da83152b5d6177b4bfe3d12810a5afd25355170/entry.go#L142-L158

      {"fields":{"app":"myapp","env":"prod","path":"Readme.md"},"level":"info","timestamp":"2023-02-02T15:36:57.222212167+08:00","message":"opening"}
      {"fields":{"app":"myapp","duration":0,"env":"prod","path":"Readme.md"},"level":"info","timestamp":"2023-02-02T15:36:57.22255411+08:00","message":"opening"}

## [gotrace example](https://github.com/kaihendry/slogfest)

Sample code:

    slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))
    var err error
    defer gotrace.Trace("fetching").Stop(&err)
    _, err = http.Get("https://httpbin.org/delay/2")

Sample output:

    {"time":"2023-02-02T15:58:42.77263618+08:00","level":"INFO","msg":"fetching"}
    {"time":"2023-02-02T15:58:45.686097016+08:00","level":"INFO","msg":"fetching","duration":2913}
