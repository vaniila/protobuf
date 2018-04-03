package vproto

import (
	"errors"
	"testing"

	proto "github.com/vaniila/protobuf/proto"
)

func TestNullErrorExtendedResponse(t *testing.T) {
	o := &ExtendedResponse{}
	b, err := proto.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("[response] before -> ", o)
	t.Log("[error]    before -> ", o.Error)
	var r = new(ExtendedResponse)
	err = proto.Unmarshal(b, r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("[response] after -> ", r)
	t.Log("[error]    after -> ", r.Error)
}

func TestNonNullErrorExtendedResponse(t *testing.T) {
	o := &ExtendedResponse{
		Error: errors.New("random error message"),
	}
	b, err := proto.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("[response] before -> ", o)
	t.Log("[error]    before -> ", o.Error)
	var r = new(ExtendedResponse)
	err = proto.Unmarshal(b, r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("[response] after -> ", r)
	t.Log("[error]    after -> ", r.Error)
}
