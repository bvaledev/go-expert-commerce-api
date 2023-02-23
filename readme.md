# Commerce

## Install Swag

`go install github.com/swaggo/swag/cmd/swag@latest`

go env

and see where is installed your go
 
 ex /home/brendo/projects-go

 navigate to bin folder inside.. cd bin

 you must see the swag bin

 set your env with `export PATH=$PATH:$GOPATH/bin`
 

init swag

swag init -g ./cmd/server/main.go
go run ./cmd/server/main.go