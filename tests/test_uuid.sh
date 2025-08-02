curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/create" -H "accept: application/json" -H "content-type: application/json" -d '{ "ip": "string", "mask": 0, "default_gateways": [ "10.105.2.1" ], "tags": [ { "key": "namespace", "value": "infiniblock.io" } ], "status": 1}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/pool/create" -H "accept: application/json" -H "content-type: application/json" -d '{ "ip": "10.105.2.0", "mask": 24, "pool": { "start": "10.105.2.100", "end": "10.105.2.200", "status": 1, "tags": [ { "key": "namespace", "value": "infiniblock.io" } ] }}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356121212"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh2343561212123"}'
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.101/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.1010/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.100/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.1002/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.102/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh2343561212123"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356121212"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh23435667"}'
curl -X  GET "http://127.0.0.1:25577/api/v0/ip/list" -H "accept: application/json" -H "content-type: application/json" 
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356678"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356679"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356678"}'
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.102/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.1025/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.105/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356679"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh2343566710"}'
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh2343566711"}'
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.105/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.102/deactivate" -H "accept: application/json" -H "content-type: application/json"
curl -X POST "http://127.0.0.1:25577/api/v0/network/10.105.2.0/24/drawip" -H "accept: application/json" -H "content-type: application/json" -d '{ "uuid": "ab.cd.ef.asdh234356679"}'
