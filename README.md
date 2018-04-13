# golang_webapidemo

Web API demo with MySQL support

## Starting the sample application

### Clone this repository

```bash
git clone git@github.com:norberteder/golang_webapidemo.git
cd golang_webapidemo
```

### Start the sample application using Docker Compose
```bash
docker-compose up -d
```

## Endpoints

Create a new contact

```bash
http post http://localhost:8081/api/1/contacts firstName=Norbert lastName=Eder email=norberteder@fakemail.com
```

Get all contacts

```bash
http get http://localhost:8081/api/1/contacts
```

Expected result:

```json
[
    {
        "email": "norberteder@fakemail.com",
        "firstName": "Norbert",
        "id": 1,
        "lastName": "Eder"
    }
]
```

(Sample API calls use  [HTTPie](https://httpie.org/))

## Troubleshooting

If the http calls don't work it might caused by the crashed app container.
This can happen if the database is starting to slow and is not available when the app container is trying to access it on startup (the app should use some retry for database connections).

To validate this, just run `docker logs golangwebapidemo_app_1`.

The result might be:
```bash
2018/04/13 20:15:28 Initializing database ...
2018/04/13 20:15:28 dial tcp 172.31.0.2:3306: connect: connection refused
panic: dial tcp 172.31.0.2:3306: connect: connection refused
```

If this is the result, just wait a few seconds and re-run `docker-compose up -d`