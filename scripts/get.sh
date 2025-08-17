#!/usr/bin/env bash

USERS=(1 2 3)

for i in $(seq 1 50); do
  USER_ID=${USERS[$RANDOM % ${#USERS[@]}]}
  echo "Dispatching GET request $i -> user $USER_ID"

  (
    RESPONSE=$(curl -s "http://localhost:8080/user/$USER_ID/balance")
    echo "Response $i: $RESPONSE"
  ) &
done

wait

echo "✅ Finished sending 50 GET requests."
