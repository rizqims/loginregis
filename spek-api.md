### Register

#### POST api/v1/users/register

request
```json
{
    "name":"string",
    "address":"string",
    "age":int,
    "username":"string",
    "password":"string"
}
```

response
```json
{
    "status":int,
    "message":"string",
    "data":{
        "name":"string",
        "address":string",
        "age":int,
        "username":"string",
        "password":"string
    }
}
```

#### POST api/v1/users/login

request
```json
{
    "username":"string",
    "password":"string"
}
```

response
```json
{
    "status":int,
    "message":"string",
    "data": 0
}