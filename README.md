# General

These examples for the [Icinga 2 API](http://docs.icinga.com/icinga2/latest/doc/module/icinga2/chapter/icinga2-api#icinga2-api)
should help you get started with your own projects.

Please read the API documentation thoroughly before looking
into the scripting details.

# Support

These examples remain generally unsupported, you should not put them in
production without your own review and knowledge.

Discuss your questions on the [community channels](https://www.icinga.com/community/get-involved/).

# Integrations

If any tool or script is missing, please send a patch/PR :)

## Libraries

Name												| Language	| Description
------------------------------------------------------------------------------------------------|---------------|--------------------------------------------------------
[python-icinga2-api](https://pypi.python.org/pypi/python-icinga2api)				| Python	| Python bindings for Icinga 2 interaction
[go-icinga2](https://github.com/xert/go-icinga2)						| Golang	| Golang functions and type definitions
[go-icinga2-api](https://github.com/lrsmith/go-icinga2-api/)					| Golang	| Golang implementation used inside the Terraform provider
[go-icinga2-client](https://github.com/Nexinto/go-icinga2-client)     | Golang  | Golang implementation for the Rancher integration.
[Monitoring::Icinga2::Client::REST](https://metacpan.org/release/THESEAL/Monitoring-Icinga2-Client-REST-2.0.0) | Perl | Perl bindings.

## Status

Name												| Language	| Description
------------------------------------------------------------------------------------------------|---------------|--------------------------------------------------------
[Dashing](https://github.com/icinga/dashing-icinga2)						| Ruby, HTML	| Dashboard for Dashing querying the REST API for current host/service/global status
[InfluxDB Telegraf Input](https://github.com/influxdata/telegraf/pull/2668)			| Golang	| [Telegraf](https://github.com/influxdata/telegraf) is an agent written in Go for collecting, processing, aggregating, and writing metrics.
[icinga2bot](https://github.com/reikoNeko/icinga2bot)						| Python	| [Errbot](http://errbot.io/en/latest/user_guide/setup.html) plugin to fetch status and event stream information and forward to XMPP, IRC, etc.
[IcingaBusyLightAgent](https://github.com/stdevel/IcingaBusylightAgent) 			| C#		| Notification Agent in Systray
[BitBar for OSX](https://getbitbar.com/plugins/Dev/Icinga2/icinga2.24m.py)			| Python	| macOS tray app for highlighting the host/service status
[Icinga 2 Multistatus](https://chrome.google.com/webstore/detail/icinga-multi-status/khabbhcojgkibdeipanmiphceeoiijal/related)	| - 	| Chrome Extension
[Clippy.js](clippy.js/)										| PHP, JS	| Funny demo for presenting alerts in your browser

## Manage Objects

Name												| Language	| Description
------------------------------------------------------------------------------------------------|---------------|--------------------------------------------------------
[Icinga Director](https://www.icinga.org/products/icinga-web-2-modules/) 			| PHP, JS	| Icinga 2 configuration interface integrated into Icinga Web 2
[AWS/EC2](aws-ec2/)										| Ruby		| Example script for creating and deleting AWS instances in Icinga 2
[Foreman Smart Proxy Monitoring](https://github.com/theforeman/smart_proxy_monitoring)		| Ruby		| Smart Proxy extension for Foreman creating and deleting hosts and services in Icinga 2
[Terraform Provider](https://github.com/hashicorp/terraform/pull/8306)				| Golang	| Register hosts from Terraform in Icinga 2. [Official docs](https://www.terraform.io/docs/providers/icinga2/index.html).
[Rancher integration](https://github.com/Nexinto/rancher-icinga)              | Golang  | Registers [Rancher](http://rancher.com/rancher/) resources in Icinga 2 for monitoring.

## Event Streams

Name												| Language	| Description
------------------------------------------------------------------------------------------------|---------------|--------------------------------------------------------
[Request Tracker ticket integration](https://github.com/bytemine/icinga2rt)			| Golang	| Create and update RT tickets
[Elastic icingabeat](https://github.com/icinga/icingabeat)					| Golang	| Process events and send to Elasticsearch/Logstash outputs
[Logstash input event stream](https://github.com/bobapple/logstash-input-icinga_eventstream)	| Ruby		| Forward events as Logstash input
[Flapjack events](https://github.com/sol1/flapjack-icinga2)					| Golang	| Dumping events into Redis for Flapjack processing
[Stackstorm integration](https://github.com/StackStorm-Exchange/stackstorm-icinga2)		| Python	| Processing events and fetching status information

## Actions

Name												| Language	| Description
------------------------------------------------------------------------------------------------|---------------|--------------------------------------------------------
[mqttwarn](https://github.com/jpmens/mqttwarn#icinga2)						| Python	| Forward check results from mqttwarn to Icinga 2
[Lita handler](https://github.com/tuxmea/lita-icinga2)						| Ruby		| List, recheck and acknowledge through a #chatops bot called [Lita](https://github.com/litaio/lita)
[Sakuli forwarder](http://sakuli.readthedocs.io/en/dev/forwarder-icinga2api/)			| Java		| Forward check results from tests from [Sakuli](https://github.com/ConSol/sakuli) to Icinga 2
[OpsGenie actions](https://www.opsgenie.com/docs/integrations/icinga2-integration)		| Golang, Java	| Integrate Icinga 2 into OpsGenie


## REST API Clients

Name												| Language	| Description
------------------------------------------------------------------------------------------------|---------------|--------------------------------------------------------
Browser plugins											| -		| [Postman for Chrome](https://www.getpostman.com), [REST Easy for Firefox](https://addons.mozilla.org/en-US/firefox/addon/rest-easy/?src=userprofile)
[Paw for MacOS](https://paw.cloud)								| (exported)	| Examples [here](paw/)
[Icinga Studio](http://docs.icinga.org/icinga2/latest/doc/module/icinga2/toc#!/icinga2/latest/doc/module/icinga2/chapter/icinga2-api#icinga2-api-clients-icinga-studio)	| C++	| Application for visualizing the status information
[icinga2 console](http://docs.icinga.org/icinga2/latest/doc/module/icinga2/toc#!/icinga2/latest/doc/module/icinga2/chapter/icinga2-api#icinga2-api-clients-cli-console)	| C++	| CLI tool for running config expressions against the API

## Misc

Several [Scripts](scripts/).

Name												| Language	| Description
------------------------------------------------------------------------------------------------|---------------|--------------------------------------------------------
[go-icinga2-events](go-icinga2-events/)								| Golang	| Connect to the event stream and output state changes
[console](scripts/console/)									| -		| Examples for using the icinga2 console CLI command
[events](scripts/events/)									| -		| Examples for event streams
[objects](scripts/objects/)									| PHP, Python	| Examples for fetching status and managing objects
