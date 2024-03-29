FROM node:16-alpine AS nodebuild

WORKDIR /app
ADD . .

RUN npm --prefix web-ui i && \
    npm --prefix web-ui run build


FROM golang:1.18.8-alpine3.17 AS gobuild

WORKDIR /app
ADD . .

ENV GO111MODULE on

RUN apk add -U --no-cache ca-certificates git make musl-dev gcc

COPY --from=nodebuild /app/web-ui/build ./web-ui/build

RUN make install && \
    make build

FROM alpine:3.17 as final

WORKDIR /home

COPY --from=gobuild /app/bin/brew-web .
COPY --from=gobuild /app/config.yml .

ENTRYPOINT [ "./brew-web" ]
CMD ["-type=pi"]
