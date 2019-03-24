SRC="$HOME/go/src/github.com/tozastation/go-grpc-streaming-example/idl/output"
# Generate User
protoc -I idl/ --plugin=$HOME/go/bin/protoc-gen-go --go_out=plugins=grpc:$SRC idl/music.proto