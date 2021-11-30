CGO_ENABLED=1 
GO111MODULE=on 
GOOS=linux 
GOARCH=amd64 
go build -o ./bin/console ./services/console/main.go
go build -o ./bin/gate ./services/gate/main.go
go build -o ./bin/live ./services/live/main.go