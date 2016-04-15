#!/bin/bash

curl -k -s -u root:icinga 'https://localhost:5665/v1/objects/services?attrs=__name&attrs=state&filter=service.state==ServiceCritical' | grep -c __name
