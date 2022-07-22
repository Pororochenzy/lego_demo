SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
cd ../
go build -o ./bin/console.exe ./services/console/main.go
go build -o ./bin/gate.exe ./services/gate/main.go
go build -o ./bin/live.exe ./services/live/main.go
