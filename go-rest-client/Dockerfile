FROM golang:1.20-bullseye as builder

WORKDIR /app
COPY *.go go.* GNUmakefile ./
COPY internal ./internal/
RUN make

FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=builder /app/go-rest-client ./

CMD ["/go-rest-client"]