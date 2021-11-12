MAIN_FILE="$1"
GOOS="$2"
GOARCH="$3"

echo "running tests"
go test

echo "building app"
go build .
