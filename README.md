# General

Simple AWS instances fetch script for creating, modifying and
deleting instances in Icinga 2 using the REST API (2.4+)

# Requirements

* Ruby, Gems (aws-sdk, rest-client, colorize)
* Icinga 2 2.4+

    sudo gem install aws-sdk
    sudo gem install rest-client

http://docs.aws.amazon.com/sdkforruby/api/index.html

## AWS Credentials

http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs
https://console.aws.amazon.com/iam/home?region=us-west-2#security_credential

    mkdir $HOME/.aws
    vim $HOME/.aws/credentials

    [default]
    aws_access_key_id = ACCESS_KEY
    aws_secret_access_key = SECRET_KEY
    aws_region = us-west-2

## Icinga 2

We'll be using the icinga2's client certificate and common name for authentication
instead of basic auth.

Therefore edit the ApiUser class and set the `client_cn` to NodeName.

    vim /usr/local/icinga2/etc/icinga2/conf.d/api-users.conf

    object ApiUser "root" {
      password = "icinga"
      client_cn = NodeName
    }

# Demo

* Start Icinga 2
* Start/Stop AWS EC2 instances

https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#Instances:sort=desc:statusChecks

* Run the script
