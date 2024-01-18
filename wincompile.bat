go env -w GOARCH="386" CGO_ENABLED=1
go build -o dist/game-deal.exe
go env -w GOARCH="amd64" CGO_ENABLED=0
echo "done"