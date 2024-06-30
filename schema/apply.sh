#!/bin/sh
atlas schema fmt --config file://schema.hcl && atlas schema apply \
  --url "postgresql://postgres:postgres@localhost:5432/webhook-multiplexer?sslmode=disable" \
  --to "file://schema.hcl"

#atlas schema fmt --config file://schema.hcl && atlas schema apply \
#  --url "postgresql://{username}:{password}@{host}:{port}/webhook-multiplexer?sslmode=disable" \
#  --to "file://schema.hcl"