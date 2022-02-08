
#!/bin/bash

DIR="./gen/swift/v1/AswProtobuf/Sources/AswProtobuf"

if [ ! -d "$DIR" ]; then  
    mkdir -p ./gen/swift/v1
fi

if brew ls --versions protobuf > /dev/null; then
    echo "checked : protobuf installed"
else
    brew install protobuf
    echo "install : protobuf"
fi
if brew ls --versions grpc-swift > /dev/null; then
    echo "checked : grpc-swift installed"
else
    brew install grpc-swift
    echo "install : grpc-swift"
fi
if brew ls --versions swift-protobuf > /dev/null; then
    echo "checked : swift-protobuf installed"
else
    brew install swift-protobuf
    echo "install : swift-protobuf"
fi


for file in "./protos/v1"/*
do
	protoc --swift_out=$DIR --grpc-swift_out=$DIR -I ./protos/v1 "$file"
	echo "WORK : [$file]"
done

echo "build : swift protobuf, [$DIR]"

