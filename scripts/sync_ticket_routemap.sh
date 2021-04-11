#!/bin/sh

set -ex

tmpdir=`ttpage ./ticket/routemap.toml`

cp "$tmpdir/page/ticket/routemap.toml" ./ticket/routemap.toml
cp "$tmpdir/page/ticket/handlers.go" ./ticket/handlers.go
rsync -r "$tmpdir/page/ticket/templates/" ./ticket/templates/
go generate ./ticket