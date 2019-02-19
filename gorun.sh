export $(cat .env | xargs)
go run -tags=debug main.go
