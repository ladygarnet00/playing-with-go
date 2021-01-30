# playing-with-go
Repository to test Golang library

```bash
go build .
```


```bash
go run .
```

# Test the application
To test the web server, using postman do the following http requests:

## Create new user
```json
curl --location --request POST 'localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "FirstName": "Carolina",
    "LastName": "Martinez"
}'
```

## Obtain all users
```json
curl --location --request GET 'localhost:3000/users' \
--header 'Content-Type: application/json' \
```

## Obtain user by id
```json
curl --location --request GET 'localhost:3000/users/1' \
--header 'Content-Type: application/json' \
```

