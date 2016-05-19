# KV Service

KV service is a microservice that serves as a key-value store using gossip. It leverages go-platform/kv 
and participates in a consistent hash ring by broadcasting its availability as a server. 

Why is this here? It's a method of offloading key-value storage to a backend service while keeping it 
in the realm of Micro services. It's built, deployed and run the same way as everything else.

Why is there nothing here? Most of the code exists in go-platform/kv.


## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -dev -advertise=127.0.0.1
	```

3. Download and start the service

	```shell
	go get github.com/micro/kv-srv
	kv-srv
	```

	OR as a docker container

	```shell
	docker run microhq/kv-srv --registry_address=YOUR_REGISTRY_ADDRESS
	```
## API

Store
- Get
- Put
- Del
