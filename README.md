# carbonauth

```
curl -u admin:admin -X POST localhost:8080/users -d '{
  "username": "user",
  "password": "secret",
  "globs":    ["foo.*", "bar.baz"]
}'
```
