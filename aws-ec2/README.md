# General

Simple AWS EC2 instances script for creating, modifying and
deleting instances as host objects in Icinga 2 using the REST API (2.4+)

**This is intended for demo purposes and not for production usage.**
You may use the provided examples in your own implementation.

# Requirements

* Ruby, Gems (aws-sdk, rest-client, colorize)
* Icinga 2 2.4+

Example:

    sudo gem install aws-sdk
    sudo gem install rest-client

More on the AWS Ruby SDK can be found [here](http://docs.aws.amazon.com/sdkforruby/api/index.html).

## AWS Credentials

The AWS rest client requires the credentials being stored
in your user's home directory in `.aws/credentials`.

More information can be found [here](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs).
Fetch the access and secret key from your [AWS settings](https://console.aws.amazon.com/iam/home?region=us-west-2#security_credential)
and also specify the AWS region.

Example:

    mkdir $HOME/.aws
    vim $HOME/.aws/credentials

    [default]
    aws_access_key_id = ACCESS_KEY
    aws_secret_access_key = SECRET_KEY
    aws_region = us-west-2

## Icinga 2

Icinga 2 provides either basic auth or client certificates for authentication.

Therefore add a new ApiUser object to your Icinga 2 configuration:

    vim /etc/icinga2/conf.d/api-users.conf

    object ApiUser "aws" {
      password = "icinga"
      client_cn = "icinga2a"
      permissions = [ "*" ]
    }

In case you want to use client certificates, set the `client_cn` from your connecting
host and put the client certificate files (private and public key, ca.crt) in the `pki`
directory.

> **Note**
>
> The script will attempt to use client certificates once found in the `pki/` directory
> instead of basic auth.

## Script

Adapt the credentials in the main method at the bottom.

    #i2.set_node_name("icinga2a")
    i2.set_api_username("aws")
    i2.set_api_password("icinga")


# Demo

* Start Icinga 2
* Start/Stop AWS EC2 instances

[AWS EC2 Console](https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#Instances:sort=desc:statusChecks)

* Run the script
