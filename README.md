# Wait4Port
wait4port is a simple commandline tool to wait for a tcp port to become open.
It is written in go.
## License
This tool is licensed under the MIT License

## Building
```
go get -v github.com/joernott/wait4port
```

## Running
wait4port accepts the following commandline parameters:

|Short | Long            | Purpose |                         | Default     |
|------|-----------------|-----------------------------------|-------------|
|  -h  | --help          | help for wait4port                |             |
|  -s, | --server string | Server name                       | "localhost" |
|  -p, | --port int      | Network port                      | 22          |
|  -t, | --timeout int   | Timeout for connection in seconds | 30          |
|  -r  | --retry int     | Number of retries                 | 10          |
|  -w, | --wait int      | Seconds to wait between tries     | 30          |
|  -v, | --verbose       | Verbose output                    | false       |
