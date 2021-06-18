FROM alpine:latest

WORKDIR /root

RUN apk --no-cache add ca-certificates

COPY cmdb ./cmdb

ENTRYPOINT ["/root/cmdb"]