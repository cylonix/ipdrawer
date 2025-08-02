#! /bin/bash

curl -X  GET "http://127.0.0.1:25577/api/v0/ip/list" -H \
    "accept: application/json" -H "content-type: application/json"
for i in {0..200}
do
    echo "Release ip 10.105.2.$i"
    curl -X POST "http://127.0.0.1:25577/api/v0/ip/10.105.2.$i/deactivate" \
        -H "accept: application/json" -H "content-type: application/json"
    echo "Release ip 100.2.2.$i"
    curl -X POST "http://127.0.0.1:25577/api/v0/ip/100.2.2.$i/deactivate" \
        -H "accept: application/json" -H "content-type: application/json"
done
curl -X  GET "http://127.0.0.1:25577/api/v0/ip/list" -H \
    "accept: application/json" -H "content-type: application/json"