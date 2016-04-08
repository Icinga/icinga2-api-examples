#!/usr/bin/env python

import requests
import json

def handle_response_execution_time(response):
    execution_times = {}

    results = response["results"]
    for r in results:
        obj_type = r["type"]
	obj_name = r["name"]
	obj_attrs = r["attrs"]

	for attr, value in obj_attrs.iteritems():
	    if attr in "last_check_result" and value is not None:
	        execution_time = float(value["execution_end"]) - float(value["execution_start"])
                execution_times[obj_name] = execution_time

	# prefer joined hosts, and override them each instead of firing two queries
	obj_joins = r["joins"]

	for join, value in obj_joins.iteritems():
	    if join in "host":
	        cr = value["last_check_result"]
		host_name = value["name"]

                if cr is not None:
                    execution_times[host_name] = float(cr["execution_end"]) - float(cr["execution_start"])


    for execution_time, obj in sorted([(value,key) for (key,value) in execution_times.items()], reverse=False):
        print "Execution time: %2f seconds, Object name '%s'" % (float(execution_time), obj)

def main():
    # Replace 'localhost' with your FQDN and certificate CN
    # for SSL verification
    #fqdn = "localhost"
    fqdn = "mbmif.int.netways.de"
    port = 5665
    user = "root"
    password = "icinga"
    request_url = "https://%s:%s/v1/objects/services" % (fqdn, port)
    headers = {
            'Accept': 'application/json',
            'X-HTTP-Method-Override': 'GET'
            }
    data = {
            "attrs": [ "name", "state", "last_check_result" ],
            "joins": [ "host.name", "host.state", "host.last_check_result" ],
    }
    
    r = requests.post(request_url,
            headers=headers,
            auth=(user, password), # modify auth
            data=json.dumps(data),
            verify="/usr/local/icinga2/etc/icinga2/pki/ca.crt") # modify trusted ca path
    
    print "Request URL: " + str(r.url)
    print "Status code: " + str(r.status_code)
    
    if (r.status_code == 200):
    	handle_response_execution_time(r.json())
    else:
            print r.text
            r.raise_for_status()


if __name__ == "__main__":
    main()
