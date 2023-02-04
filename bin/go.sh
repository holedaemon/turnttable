#!/bin/bash

PROJECT_BASE=/home/max/git/holedaemon/turnttable

qtc -dir $PROJECT_BASE/internal/web/templates

sass --style compressed --no-source-map $PROJECT_BASE/internal/web/style/turnttable.scss:$PROJECT_BASE/internal/web/static/turnttable.css

go run github.com/holedaemon/turnttable