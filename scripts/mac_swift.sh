
#!/bin/bash

cd ..
DIR="./gen/swift/v1"

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

protoc --swift_out=$DIR --grpc-swift_out=$DIR -I ./protos/v1 idl.proto
echo "build : swift protobuf, [$DIR]"