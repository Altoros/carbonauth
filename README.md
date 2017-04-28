# carbonauth

Authorization microservice for the carbon stack with forward proxy / load balancer functionality.

Using admin credentials you can manage users and their privileges for accessing underlying carbonapi servers via http basic auth.

## Compatibility

* `/metrics/find`
* `/render` only with `format=json`

## Configuration

Copy `config.yml.example` to `config.yml` and adjust configuration values.

## Usage

Create or update user:
```
curl -X POST admin:admin@localhost:8082/users -d '{
  "username": "test",
  "password": "secret",
  "globs":    ["foo.*", "bar.baz"]
}'
```

Delete user:
```
curl -X DELETE admin:admin@localhost:8082/users/test
```

Querying carbon:
```
curl test:secret@localhost:8082/metrics/find?query=foo.*
```
