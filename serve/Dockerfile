FROM alpine:latest

WORKDIR /app/server
COPY ./srv /app/server/
COPY ./config.dev.yaml /app/server/

EXPOSE 8888
ENTRYPOINT ./srv -c config.dev.yaml
