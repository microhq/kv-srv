package handler

import (
	"time"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-platform/kv"
	proto "github.com/micro/kv-srv/proto/store"

	"golang.org/x/net/context"
)

type Store struct{}

func (s *Store) Get(ctx context.Context, req *proto.GetRequest, rsp *proto.GetResponse) error {
	kk, ok := kv.FromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.srv.kv.get", "kv not initialised")
	}

	item, err := kk.Get(req.Key)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.kv.get", err.Error())
	}

	rsp.Item = &proto.Item{
		Key:        item.Key,
		Value:      item.Value,
		Expiration: int64(item.Expiration.Seconds()),
	}

	return nil
}

func (s *Store) Put(ctx context.Context, req *proto.PutRequest, rsp *proto.PutResponse) error {
	kk, ok := kv.FromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.srv.kv.get", "kv not initialised")
	}

	if err := kk.Put(&kv.Item{
		Key:        req.Item.Key,
		Value:      req.Item.Value,
		Expiration: time.Duration(req.Item.Expiration) * time.Second,
	}); err != nil {
		return errors.InternalServerError("go.micro.srv.kv.get", err.Error())
	}

	return nil
}

func (s *Store) Del(ctx context.Context, req *proto.DelRequest, rsp *proto.DelResponse) error {
	kk, ok := kv.FromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.srv.kv.get", "kv not initialised")
	}

	if err := kk.Del(req.Key); err != nil {
		return errors.InternalServerError("go.micro.srv.kv.get", err.Error())
	}

	return nil
}
