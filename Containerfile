FROM alpine


COPY _output/app /app

ENTRYPOINT ["/app"]