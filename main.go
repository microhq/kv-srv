package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-platform/kv"
	"github.com/micro/kv-srv/handler"
	proto "github.com/micro/kv-srv/proto/store"

	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.kv"),
		micro.Version("latest"),

		micro.Flags(
			cli.StringFlag{
				Name:   "namespace",
				EnvVar: "NAMESPACE",
				Usage:  "Namespace for the key-value store",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			opts := []kv.Option{
				kv.Client(service.Client()),
				kv.Server(service.Server()),
			}

			if len(c.String("namespace")) > 0 {
				opts = append(opts, kv.Namespace(c.String("namespace")))
			}

			kk := kv.NewKV(opts...)

			service.Init(micro.WrapHandler(func(fn server.HandlerFunc) server.HandlerFunc {
				return func(ctx context.Context, req server.Request, rsp interface{}) error {
					ctx = kv.NewContext(ctx, kk)
					return fn(ctx, req, rsp)
				}
			}))
		}),
	)

	proto.RegisterStoreHandler(service.Server(), new(handler.Store))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
