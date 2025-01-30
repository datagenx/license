# license-go

# To run locally

```
go run api/main.go
```

## To generate a license

```
curl --location 'localhost:9200/generate' \
--header 'Content-Type: application/json' \
--data '{
    "customer": "atul",
    "valid-from": "2021-01-01",
    "expiry-date": "2022-01-01",
    "hard-expiry-date": "2022-01-01",
    "seats": 1,
    "hard-seats": 1,
    "type": "test"
}'
```

## To generate a license

```
curl --location 'localhost:9200/validate' \
--header 'Content-Type: application/json' \
--data '{
    "customer":"atul",
    "valid-from":"2021-01-01",
    "expiry-date":"2022-01-01",
    "hard-expiry-date":"2022-01-01",
    "seats":1,
    "hard-seats":1,
    "type":"test",
    "signature":"IAZelv8qXR/xZcT1TIZozW6JryIMAyZOGEKpyiPKocXnJyvubde8m3UwbcmyzMfnGSdQP9HnlTNC8iMH3pjzDw=="
}'
```
