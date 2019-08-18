# build golang
FROM golang:alpine AS build
ADD . /src

RUN apk add -U --no-cache ca-certificates git

RUN cd /src && \
    go mod vendor && \
    go build -o brew-web

# build frontend
FROM node:alpine AS frontend

WORKDIR /web-ui

ADD web-ui .

RUN REACT_APP_BACKEND_URL=http://localhost:3001 npm run build

# final stage
FROM alpine as bare
WORKDIR /app

# backend
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/brew-web /app/

# frontend
COPY --from=frontend /web-ui/build  /app/web-ui/build

RUN apk add jq
ENV PATH="/app:${PATH}"
ENTRYPOINT brew-web