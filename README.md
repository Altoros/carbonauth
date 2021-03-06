# carbonauth

Authorization microservice for the carbon stack with forward proxy / load balancer functionality.

Using admin credentials you can manage users and their privileges for accessing underlying carbonapi servers via http basic auth.

Developed and tested as an additional layer between [grafana](https://github.com/grafana/grafana) and [carbonapi](https://github.com/go-graphite/carbonapi).

## Compatibility

* `/metrics/find`
* `/render` only with `format=json`

## Configuration

Copy `config.yml.example` to `config.yml` and adjust configuration values.

## Usage

Create or update user:
```
$ curl admin:secret@localhost:8082/users -d '{
  "username": "test",
  "password": "secret",
  "globs":    ["foo.*", "bar.baz"]
}'
```

List users:
```
$ curl admin:secret@localhost:8082/users
```

Find user:
```
$ curl admin:secret@localhost:8082/users/test
```

Delete user:
```
$ curl -X DELETE admin:secret@localhost:8082/users/test
```

Querying carbon:
```
$ curl test:secret@localhost:8082/metrics/find?query=foo.*
$ curl test:secret@localhost:8082/render -d "target=foo.*&format=json"
```
