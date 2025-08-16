# Entain Test task

```bash
$ migrate -path=./migrations -database=$ENTAIN_DB_DSN up
```

```bash
curl -X POST http://localhost:8080/user/2/transaction \
  -H "Content-Type: application/json" \
  -H "Source-Type: game" \
  -d '{
    "state": "lose",
    "amount": "0.15",
    "transactionId": "a47f9b2d-1c3e-4f6a-8d2b-5e9c7a1b3f4d"
  }'
```
