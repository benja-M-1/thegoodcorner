# Fizzbuzz [![Build Status](https://travis-ci.com/benja-M-1/thegoodcorner.svg?branch=master)](https://travis-ci.com/benja-M-1/thegoodcorner) [![Go Report Card](https://goreportcard.com/badge/github.com/benja-M-1/thegoodcorner)](https://goreportcard.com/report/github.com/benja-M-1/thegoodcorner)

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by “fizz”,
all multiples of 5 by “buzz”, and all multiples of 15 by “fizzbuzz”. The output would look like this: 
“1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...”.

Your goal is to implement a web server that will expose a REST API endpoint that: 
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all 
multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

## Bonus question :

Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request

---

## Run the API

First install the project running `make install`.

Then create the database replacing ${DB_USERNAME}, ${DB_HOSTNAME} and ${DB_DATABASE} with the values contained in `.env` file:
```bash
docker-compose exec db psql --username=${DB_USENAME} --host=${DB_HOSTNAME} --dbname=${DB_DATABASE} -a --file=init_db.sql
```

Finally run `make serve` to start the API.

Make a request on `/fizzbuzz`:

```bash
curl -X GET "http://localhost:80/fizzbuzz?int1=2&int2=4&str1=fizz&str2=buzz&limit=50"
```

Fetch the statistics on `/statistics`:

```bash
curl -X GET http://localhost:80/statistics
```

## Run the tests

Simply run `make test`