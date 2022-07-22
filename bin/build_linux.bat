set GOOS=linux
set CGO_ENABLED=0
cd ../
del bin/console,bin/gate,bin/live
go build -o ./bin/console ./services/console/main.go
go build -o ./bin/gate ./services/gate/main.go
go build -o ./bin/live ./services/live/main.go
REM pause