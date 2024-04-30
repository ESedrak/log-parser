# Log Parser

## Overview

### High Level Design

![docs/HLD.png](docs/HLD.png)

Log Parser counts the unique IP addresses, top most requested URL's, and the most active IP addresses for the following HTTP methods:

```text
GET, PUT, DELETE, POST, HEAD
```

## Dependencies

### Install Go version 1.21

<https://golang.org/doc/install>

#### On macs: go can be installed through homebrew

brew install go@1.21

### Makefile

Should be installed in most UNIX like systems.
