# Log Parser

## Overview

### High Level Design

![docs/HLD.png](docs/HLD.png)

Log Parser counts the unique IP addresses, top most requested URL's, and the most active IP addresses for the following HTTP methods:

```text
GET, PUT, DELETE, POST, HEAD
```

The number of requested URL's and most active IP addresses can be set by the requestedCountURL and requestCountIP. Currently - both are set to 3.

The format of the logs will need to be in:

```text
$client_ip - - \[$time_stamp\] "$http_method $request_path $_" $response_code $_ "-" "$user_agent"
```

E.g

```text
177.71.128.21 - - [10/Jul/2018:22:22:08 +0200] "GET /blog/2018/08/survey-your-opinion-matters/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"
```

These logs are currently being stored and read from logs/log_file.log

## Dependencies

### Install Go version 1.22

<https://golang.org/doc/install>

#### On macs: go can be installed through homebrew

brew install go@1.22

### Makefile

Should be installed in most UNIX like systems.

## Makefile steps

To build and test the application:

```sh
make all
```

To run the application:

```sh
make run
```

Run all tests:

```sh
make test
```

Run benchmark tests:

```sh
make bench-all
```

Remove all compiled executables instance

```sh
make clean
```

> [!TIP]
Please see the Makefile for all commands.
