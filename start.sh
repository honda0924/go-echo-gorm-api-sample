#!/bin/bash -eu
go mod tidy
air
go run main.go router.go