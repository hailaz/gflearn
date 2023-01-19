#!/bin/bash
echo "go 1.18"> go.work
for file in $(find . -name go.mod); do
    go work use $(dirname $file)
done