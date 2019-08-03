FROM golang:alpine AS build
ADD . /src

RUN apk add -U --no-cache ca-certificates git

RUN cd /src && \
    go mod vendor && \
    go build -o brew-web

# final stage
FROM alpine as bare
WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/ecs-go /app/

RUN apk add jq
ENV PATH="/app:${PATH}"
ENTRYPOINT brew-web