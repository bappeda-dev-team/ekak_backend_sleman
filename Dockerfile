ARG GO_VERSION=1.25.5

FROM registry.docker.com/library/golang:$GO_VERSION-alpine AS base

# app lives here
WORKDIR /app


# Throw-away build stage to reduce size of final image
FROM base as build

# Install packages needed to build
RUN apk update -qq && \
    apk add --no-cache git

COPY . .

# dummy .env
RUN touch .env

RUN go build -o ekak_backend_sleman main.go wire_gen.go

ENTRYPOINT ["/app/ekak_backend_sleman"]

CMD ["app/ekak_backend_sleman"]
