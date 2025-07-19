#!/bin/bash

migrate -path scripts/postgresql/migrations -database postgres://postgres:postgres@localhost:5432/gajiandb?sslmode=disable up