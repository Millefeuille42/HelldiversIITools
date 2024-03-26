FROM golang:1.20-alpine3.18 as builder

RUN apk add gcc g++

ADD go.mod /Helldivers2Tools/go.mod
ADD go.sum /Helldivers2Tools/go.sum

WORKDIR /Helldivers2Tools
RUN go mod download

ADD pkg /Helldivers2Tools/pkg
ADD cmd /Helldivers2Tools/cmd

RUN CGO_ENABLED=1 GOOS=linux go build -v -o /bin/ ./...

FROM alpine as api
COPY --from=builder /bin/api /
ADD data /Helldivers2Tools/data
ENTRYPOINT ["/api"]

FROM alpine as bot
COPY --from=builder /bin/bot /
ENTRYPOINT ["/bot"]

FROM alpine as updater
COPY --from=builder /bin/updater /
ENTRYPOINT ["/updater"]
