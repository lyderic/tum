echo "linux amd64"
GOOS=linux GOARCH=amd64 go install

echo "darwin amd64"
GOOS=darwin GOARCH=amd64 go install

echo "linux arm"
GOOS=linux GOARCH=arm go install

echo "windows amd64"
GOOS=windows GOARCH=amd64 go install

echo "freebsd amd64"
GOOS=freebsd GOARCH=amd64 go install
