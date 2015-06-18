# slackcmd

A simple server to receive and handle Slack Commands written in GO (known as Golang)

------

[![Build Status](https://travis-ci.org/ngdinhtoan/slackcmd.svg)](https://travis-ci.org/ngdinhtoan/slackcmd)
[![GoDoc](https://godoc.org/github.com/ngdinhtoan/slackcmd?status.svg)](https://godoc.org/github.com/ngdinhtoan/slackcmd)
[![Join the chat at https://gitter.im/ngdinhtoan/slackcmd](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ngdinhtoan/slackcmd?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Install and Import SlackCmd

Of course, you have to [install GO](https://golang.org/doc/install) first if you do not have GO on your system.

Get SlackCmd by running command

    go get github.com/ngdinhtoan/slackcmd

and import slackcmd package into your project

```go
import "github.com/ngdinhtoan/slackcmd"
```

If you want to use stable version,
don't want the changes in master branch affect to your project,
use SlackCmd `v1` by

```go
import (
    slackcmd "gopkg.in/ngdinhtoan/slackcmd.v1"
)
```

## How to use SlackCmd

Checkout package `github.com/ngdinhtoan/slackcmd/example` to see how to use it

### How to write new commander

TBD: `Hello World` commander!!!

You can check package `github.com/ngdinhtoan/slackcmd/jira` as an example

### Use package `webhook` to send rich message to Slack

TBD

## Contribution

If you have a contribution, new commander or any idea to share, feel free to create a pull request or open a ticket,
or join to chat at [![https://gitter.im/ngdinhtoan/slackcmd](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ngdinhtoan/slackcmd?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## License

SlackCmd is licensed under the [MIT License](https://github.com/ngdinhtoan/slackcmd/blob/master/LICENSE)
