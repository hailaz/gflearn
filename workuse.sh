for file in $(find . -name go.mod); do
    go work use $(dirname $file)
done