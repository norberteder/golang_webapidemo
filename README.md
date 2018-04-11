# golang_webapidemo

Web API demo with MySQL support

## Installation

Execute `contact.sql`. This creates a new table `contact`. A schema named `contacts` must exist.

## Endpoints

Get all contacts

```
GET http://localhost:8081/api/1/contacts
```

Create a new contact

```
POST http://localhost:8081/api/1/contacts
```
