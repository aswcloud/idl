cd ..

if not exist .\gen\go\v1\ ( 
    mkdir .\gen\go\v1
)

protoc --go_out=.\gen\go\v1 --go_opt=paths=source_relative --go-grpc_out=.\gen\go\v1 --go-grpc_opt=paths=source_relative -I .\protos\v1 idl.proto

