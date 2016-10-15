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

# Configuration

**TODO**: Hardcoded inside the script.

* API URL, Username, Password
* SSL Verification

# Run

Build and run the binary. There are currently no external libraries required.

    export GOPATH=`pwd`
    go build
    ./go-icinga2-events

# TODO

* Config file for API credentials
* Support for multiple types (currently CheckResult hardcoded)
 * This requires mapping the static structs to the JSON messages
* Support for definable hooks (e.g. "OnStateChange", "OnFlappingDetected", etc.) for easier demos
* Support for triggering API actions on specific events (reschedule a check, auto-acknowledge problems based on a specific custom attributes, etc.)
* Support for querying objects on demand (the event message payload doesn't provide them)
