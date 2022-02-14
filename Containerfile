FROM alpine

ENV SERVER_PORT=80

COPY _output/app /app

ENTRYPOINT ["/app"]