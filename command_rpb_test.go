package riak

import (
	rpbRiak "github.com/basho-labs/riak-go-client/rpb/riak"
	rpbRiakKV "github.com/basho-labs/riak-go-client/rpb/riak_kv"
	proto "github.com/golang/protobuf/proto"
	"reflect"
	"testing"
)

func TestEnsureCorrectRequestAndResponseCodes(t *testing.T) {
	var cmd Command
	var msg proto.Message
	// Misc commands
	// Ping
	cmd = &PingCommand{}
	if expected, actual := rpbCode_RpbPingReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbPingResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if msg != nil {
		t.Error("expected nil response protobuf message")
	}
	// FetchBucketProps
	cmd = &FetchBucketPropsCommand{}
	if expected, actual := rpbCode_RpbGetBucketReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbGetBucketResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiak.RpbGetBucketResp); !ok {
		t.Errorf("error casting %v to RpbGetBucketResp", reflect.TypeOf(msg))
	}
	// StoreBucketProps
	cmd = &StoreBucketPropsCommand{}
	if expected, actual := rpbCode_RpbSetBucketReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbSetBucketResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if msg != nil {
		t.Error("expected nil response protobuf message")
	}

	// KV commands
	// FetchValue
	cmd = &FetchValueCommand{}
	if expected, actual := rpbCode_RpbGetReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbGetResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbGetResp); !ok {
		t.Errorf("error casting %v to RpbGetResp", reflect.TypeOf(msg))
	}
	// StoreValue
	cmd = &StoreValueCommand{}
	if expected, actual := rpbCode_RpbPutReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbPutResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbPutResp); !ok {
		t.Errorf("error casting %v to RpbPutResp", reflect.TypeOf(msg))
	}
	// DeleteValue
	cmd = &DeleteValueCommand{}
	if expected, actual := rpbCode_RpbDelReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbDelResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if msg != nil {
		t.Error("expected nil response protobuf message")
	}
	// ListBuckets
	cmd = &ListBucketsCommand{}
	if expected, actual := rpbCode_RpbListBucketsReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbListBucketsResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbListBucketsResp); !ok {
		t.Errorf("error casting %v to RpbListBucketsResp", reflect.TypeOf(msg))
	}
}
