# General

Simple script fetching event streams from the Icinga 2 API written
in Golang.

**This is intended for demo purposes and not for production usage.**
You may use the provided examples in your own implementation.

# Requirements

* Golang 1.4+
* Icinga 2 2.5+

## Icinga 2

Icinga 2 provides either basic auth for authentication.

Therefore add a new ApiUser object to your Icinga 2 configuration:

    vim /etc/icinga2/conf.d/api-users.conf

    object ApiUser "streams" {
      password = "icinga"
      permissions = [ "events" ]
    }

## Script

Build and run the binary.

    export GOPATH=`pwd`
    go build
    ./go-icinga2-events

# Configuration

**TODO**: Hardcoded inside the script.

* API URL, Username, Password
* SSL Verification
