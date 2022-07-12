#!/usr/bin/env bash

ego-go build -o app app.go
ego sign app
