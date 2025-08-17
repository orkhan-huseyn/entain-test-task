# Entain Test task

To run the application you just need to run `docker compose up -d`. This should run two containers,
one for postgres and one for the Go application itself.

The Go application has default port of `8080` and this port is exposed.

## Example requests

Since the application starts with three different users with ids `1`, `2` and `3`, you can check their balance by using get balance endpoint:

```bash
curl -s "http://localhost:8080/user/1/balance"
```

Alternatively, to create a transaction and update user's balance you can use create transaction endpoint:

```bash
curl -X POST http://localhost:8080/user/1/transaction \
  -H "Content-Type: application/json" \
  -H "Source-Type: game" \
  -d '{
    "state": "win",
    "amount": "5.75",
    "transactionId": "f3a2c7d1-8b4f-43b7-92b1-6d9a8e2c5e31"
  }'
```

You can generate UUID's using `uuidgen | tr '[:upper:]' '[:lower:]'` command on MacOS and Linux. Or you can run `post.sh` script inside `scripts` folder simulate some real traffic.

The `scripts` folder also contains `get.sh` where you can send 50 concurrent requests to get balance endpoint.
