# golang-api-example
## Getting Started
1. Change the env file to fit your environmental variables
2. Run `docker-compose up --build`
3. ...
4. Profit!

## Docs
### GET /api/v1/users

Returns list of users

```
curl -X GET http://localhost:8080/api/v1/users
```

Output:

```
[
  {
    "ID": 1,
    "CreatedAt": "2017-05-23T05:37:13Z",
    "UpdatedAt": "2017-05-23T05:37:13Z",
    "DeletedAt": null,
    "Age": 21,
    "Name": "NewFoo",
    "Email": "newfoobar"
  },
  {
    "ID": 2,
    "CreatedAt": "2017-05-23T06:00:14Z",
    "UpdatedAt": "2017-05-23T06:00:14Z",
    "DeletedAt": null,
    "Age": 21,
    "Name": "JP",
    "Email": "jp@jpvillavicencio.com"
  }
]
```

### GET /api/v1/users/{userId}

Returns list of users

```
curl -X GET http://localhost:8080/api/v1/users/1
```

Output:

```
[
  {
    "ID": 1,
    "CreatedAt": "2017-05-23T05:37:13Z",
    "UpdatedAt": "2017-05-23T05:37:13Z",
    "DeletedAt": null,
    "Age": 21,
    "Name": "NewFoo",
    "Email": "newfoobar"
  },
]
```

### POST /api/v1/users/add

Creates user and adds it to DB.

```
curl -X POST \
	http://localhost:8080/api/v1/users/add \
	-H 'content-type: application/json' \
	-d '{
		"Age": 21,
		"Name": "JP",
		"Email": "jp@jpvillavicencio.com"
	}'
```

Output:

```
{
	"ID":4,
	"CreatedAt":"2017-05-23T06:00:14.030481565Z",
	"UpdatedAt":"2017-05-23T06:00:14.030481565Z",
	"DeletedAt":null,
	"Age":21,
	"Name":"JP",
	"Email":"jp@jpvillavicencio.com"
}
```
