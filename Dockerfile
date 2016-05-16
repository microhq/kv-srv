FROM alpine:3.2
ADD kv-srv /kv-srv
ENTRYPOINT [ "/kv-srv" ]
