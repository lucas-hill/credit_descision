#!/bin/bash
set -e

ENDPOINT="http://localhost:8000"

echo "Waiting for DynamoDB Local..."
until aws dynamodb list-tables --endpoint-url "$ENDPOINT" --region us-east-1 > /dev/null 2>&1; do
  sleep 1
done

echo "Creating CreditApplications table..."
aws dynamodb create-table \
  --endpoint-url "$ENDPOINT" \
  --region us-east-1 \
  --table-name CreditApplications \
  --attribute-definitions \
    AttributeName=PK,AttributeType=S \
    AttributeName=SK,AttributeType=S \
  --key-schema \
    AttributeName=PK,KeyType=HASH \
    AttributeName=SK,KeyType=RANGE \
  --billing-mode PAY_PER_REQUEST \
  2>/dev/null && echo "Table created." || echo "Table already exists."

echo "DynamoDB Local ready."
