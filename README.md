# carbonauth

Authorization microservice for the carbon stack with forward proxy / load balancer feature.

### Usage

Create user:

```
curl -X POST admin:admin@localhost:8082/users -d '{
  "username": "test",
  "password": "secret",
  "globs":    ["foo.*", "bar.baz"]
}'
```

Delete user:
```
curl -X DELETE admin:admin@localhost:8082/users?username=test
```
