#！/usr/bin/env bash
#

curl --location --request POST 'http://127.0.0.1:8001/user/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName": "user1001",
    "pwd": "123456"
}'
