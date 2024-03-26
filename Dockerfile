FROM golang:1.20-alpine3.18 as base

RUN apk add gcc g++

ADD go.mod /Helldivers2Tools/go.mod
ADD go.sum /Helldivers2Tools/go.sum

WORKDIR /Helldivers2Tools
RUN go mod download

ADD pkg/shared /Helldivers2Tools/pkg/shared
ADD cmd/cli /Helldivers2Tools/cmd/cli
# Building cli here since it is shared
#  This also caches shared dependencies
RUN CGO_ENABLED=1 GOOS=linux go build -v -o /cli ./cmd/cli

# --- BUILD STAGE ---

FROM base as api_builder
ADD pkg/api /Helldivers2Tools/pkg/api
ADD cmd/api /Helldivers2Tools/cmd/api
RUN CGO_ENABLED=1 GOOS=linux go build -v -o /api ./cmd/api

FROM base as bot_builder
ADD pkg/bot /Helldivers2Tools/pkg/bot
ADD cmd/bot /Helldivers2Tools/cmd/bot
RUN CGO_ENABLED=1 GOOS=linux go build -v -o /bot ./cmd/bot

FROM base as updater_builder
ADD cmd/updater /Helldivers2Tools/cmd/updater
RUN CGO_ENABLED=1 GOOS=linux go build -v -o /updater ./cmd/updater

# --- FINAL STAGE ---

FROM alpine as api
COPY --from=api_builder /cli /cli
COPY --from=api_builder /api /api
ADD data /data
ENTRYPOINT ["/api"]

FROM alpine as bot
COPY --from=bot_builder /cli /cli
COPY --from=bot_builder /bot /bot
ENTRYPOINT ["/bot"]

FROM alpine as updater
COPY --from=updater_builder /cli /cli
COPY --from=updater_builder /updater /updater
ENTRYPOINT ["/updater"]
