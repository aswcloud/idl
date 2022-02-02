cd ..

if not exist .\gen\go\v1\ ( 
    mkdir .\gen\go\v1
)

for %%c in (.\protos\v1\*.proto) do protoc --go_out=.\gen\go\v1 --go_opt=paths=source_relative --go-grpc_out=.\gen\go\v1 --go-grpc_opt=paths=source_relative -I .\protos\v1 %%c

cd scripts