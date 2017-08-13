regenerate:
	go install github.com/vaniila/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/vaniila/protobuf/protobuf/ \
	--gogo_out=. events.proto 
