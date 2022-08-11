# joshpay
joshpay is server with multiple go modules using go workspaces


test using base64
curl --location --request POST 'http://localhost:8080/pay' \
--header 'Authorization: BASIC Sm9obklzVmFsaWRVc2Vy' \
--header 'Content-Type: application/json' \
--data-raw '{
"credit_card_token":"xxxxx",
"amount": 1002
}’

test using jwt
curl --location --request POST 'http://localhost:8080/pay' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c \
--header 'Content-Type: application/json' \
--data-raw '{
"credit_card_token":"xxxxx",
"amount": 1002
}'

