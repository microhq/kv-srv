package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-platform/kv"
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
				kv.Internal(false),
			}

			if len(c.String("namespace")) > 0 {
				opts = append(opts, kv.Namespace(c.String("namespace")))
			}

			// we don't need to store it
			// just initialise
			kv.NewKV(opts...)
                }),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
