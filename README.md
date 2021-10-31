# ldt
Load Testing Utility using Vegeta.

A very basic utility to test Get And Post APIs utilizing the powerful Vegeta library
https://github.com/tsenart/vegeta. 

### PreRequisite
Golang 1.16

### Usage
`make build`

`./out/ldt ldt -m GET -u http://localhost:8040/load-test -r 100 -d 60`

The above command will test the `load-test` GET API for 60 seconds at 100 requests per second.
