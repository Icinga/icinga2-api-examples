## List all hosts
curl -X "GET" "https://192.168.33.5:5665/v1/objects/hosts?attrs=__name&attrs=address6" \
     -u root:icinga

## List localhost
curl -X "GET" "https://192.168.33.5:5665/v1/objects/hosts?host=icinga2" \
     -H "Content-Type: application/octet-stream" \
     -u root:icinga

## List localhost service
curl -X "GET" "https://192.168.33.5:5665/v1/objects/services?service=icinga2!ping6" \
     -u root:icinga

## Process checkresult
curl -X "POST" "https://192.168.33.5:5665/v1/actions/process-check-result?service=icinga2!ping6&type=Service" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:icinga \
     -d "{\"exit_status\":2,\"plugin_output\":\"PING CRITICAL - Packet loss = 100%\",\"performance_data\":[\"rta=5000.000000ms;3000.000000;5000.000000;0.000000\",\"pl=100%;80;100;0\"],\"check_source\":\"example.localdomain\"}"

## Reschedule check
curl -X "POST" "https://192.168.33.5:5665/v1/actions/reschedule-check?filter=true&type=Service" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:icinga \
     -d "{}"

## Add comment
curl -X "POST" "https://192.168.33.5:5665/v1/actions/add-comment" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:icinga \
     -d "{\"type\":\"Service\",\"filter\":\"service.name==\\\"ping6\\\"\",\"author\":\"michi\",\"comment\":\"I don't care. Gin 4 the win\"}"

## Create host icinga.com
curl -X "PUT" "https://192.168.33.5:5665/v1/objects/hosts/icinga.com" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:icinga \
     -d "{\"templates\":[\"generic-host\"],\"attrs\":{\"address\":\"192.168.1.1\",\"check_interval\":30,\"check_command\":\"hostalive\",\"vars.os\":\"Linux\"}}"

## Modify host icinga.com
curl -X "POST" "https://192.168.33.5:5665/v1/objects/hosts/icinga.com" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:icinga \
     -d "{\"attrs\":{\"address\":\"192.168.33.5\"}}"

## Create host github.com
curl -X "PUT" "https://192.168.33.5:5665/v1/objects/hosts/github.com" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:icinga \
     -d "{\"templates\":[\"generic-host\"],\"attrs\":{\"address\":\"192.168.1.1\",\"check_command\":\"hostalive\",\"vars.os\":\"Linux\"}}"

## Delete multiple hosts
curl -X "DELETE" "https://192.168.33.5:5665/v1/objects/hosts" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:***** Hidden credentials ***** \
     -d "{\"filter\":\"match(\\\"*com*\\\", host.name)\",\"cascade\":true}"

## Acknowledge all problems
curl -X "POST" "https://192.168.33.5:5665/v1/actions/acknowledge-problem" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:icinga \
     -d "{\"type\":\"Service\",\"filter\":\"service.state!=ServiceOK\",\"author\":\"michi\",\"comment\":\"Global outage. Working on it.\",\"notify\":true}"

## Remove all acknowledgements
curl -X "POST" "https://192.168.33.5:5665/v1/actions/remove-acknowledgement" \
     -H "Accept: application/json" \
     -H "Content-Type: text/plain; charset=utf-8" \
     -u root:icinga \
     -d $'{
  "type": "Service",
  "filter": "service.state!=ServiceOK"
}'

## Create realtime load check
curl -X "PUT" "https://192.168.33.5:5665/v1/objects/services/icinga2!load-realtime" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:***** Hidden credentials ***** \
     -d "{\"templates\":[\"generic-service\"],\"attrs\":{\"check_command\":\"load\",\"check_interval\":1,\"retry_interval\":1}}"

## Delete realtime load check
curl -X "DELETE" "https://192.168.33.5:5665/v1/objects/services/icinga2!load-realtime?cascade=1" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json; charset=utf-8" \
     -u root:***** Hidden credentials ***** \
     -d "{}"

## Delete all service comments
curl -X "POST" "https://192.168.33.5:5665/v1/actions/remove-comment?type=Service&filter=true" \
     -H "Accept: application/json" \
     -u root:icinga

## Delete all host comments
curl -X "POST" "https://192.168.33.5:5665/v1/actions/remove-comment?type=Host&filter=true" \
     -H "Accept: application/json" \
     -u root:icinga
