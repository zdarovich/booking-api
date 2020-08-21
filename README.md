# Run
```shell script
git clone https://github.com/zdarovich/booking-api
cd booking-api
go run main.go
```

# Create class
```shell script
curl --location --request POST 'http://127.0.0.1:8081/classes' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--data-raw '
            {
               "start_date": "2021-08-21",
               "end_date": "2021-08-27",
               "capacity": 10,
               "name": "Rob Pike"
            }
'
```

# Create booking
```shell script
curl --location --request POST 'http://127.0.0.1:8081/bookings' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--data-raw '
            {
               "date": "2021-08-22",
               "name": "Rob Pike"
            }
'
```

# Get classes
```shell script
curl --location --request GET 'http://127.0.0.1:8081/classes' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' 
```