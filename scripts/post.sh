#!/usr/bin/env bash

generate_uuid() {
  uuidgen | tr '[:upper:]' '[:lower:]'
}

USERS=(1 2 3)
SOURCES=("game" "payment" "server")
STATES=("win" "lose")

for i in $(seq 1 30); do
  USER_ID=${USERS[$RANDOM % ${#USERS[@]}]}
  SOURCE=${SOURCES[$RANDOM % ${#SOURCES[@]}]}
  STATE=${STATES[$RANDOM % ${#STATES[@]}]}
  AMOUNT=$(printf "%.2f" "$(awk -v min=0.1 -v max=10 'BEGIN{srand(); print min+rand()*(max-min)}')")
  TX_ID=$(generate_uuid)

  echo "Dispatching request $i -> user $USER_ID, source=$SOURCE, state=$STATE, amount=$AMOUNT, txId=$TX_ID"

  (
    RESPONSE=$(curl -s -X POST "http://localhost:8080/user/$USER_ID/transaction" \
      -H "Content-Type: application/json" \
      -H "Source-Type: $SOURCE" \
      -d "{
        \"state\": \"$STATE\",
        \"amount\": \"$AMOUNT\",
        \"transactionId\": \"$TX_ID\"
      }")
    echo "Response $i: $RESPONSE"
  ) &
done

wait

echo "✅ Finished sending 30 parallel requests."
