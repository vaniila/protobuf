regenerate:
	go install github.com/vaniila/protobuf/protoc-min-version
	go install github.com/vaniila/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/vaniila/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/vaniila/protobuf/types:. timefail.proto 
