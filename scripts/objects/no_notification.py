#!/usr/bin/env python

# pip install icinga2api
# https://github.com/tobiasvdk/icinga2api
from icinga2api.client import Client

# use a helper to fetch our cut down object names
def getObjects(client, type_name):
    return client.objects.list(type_name, attrs=['__name'])

# use a helper to convert the full blown object dictionary into a list of __name elements
def getNameList(objects):
    return map(lambda x : x['attrs']['__name'], objects)

def diffList(l1,l2):
    l2 = set(l2)
    return [x for x in l1 if x not in l2]

# add connection details
client = Client('https://localhost:5665', 'root', 'icinga')

hosts = getObjects(client, 'Host')
services = getObjects(client, 'Service')
notifications = getObjects(client, 'Notification')

# debug
#print hosts
#print services
#print notifications

h_names = getNameList(hosts)
s_names = getNameList(services)
n_names = getNameList(notifications)

found_h_names = []
found_s_names = []

for n in n_names:
    split_arr = n.split("!")

    if len(split_arr) == 2:
        for h in h_names:
            # debug
            #print n + " " + h
            if h in split_arr[0]:
                found_h_names.append(h)

    if len(split_arr) == 3:
        for s in s_names:
            # debug
            #print n + " " + s

            # rebuild the matched full service name with host!service
            split_s_full_name = split_arr[0] + "!" + split_arr[1]

            if s in split_s_full_name:
                found_s_names.append(s)

print "Hosts without notification"
print ", ".join(diffList(h_names, found_h_names))
print ""
print "Services without notification"
print ", ".join(diffList(s_names, found_s_names))



