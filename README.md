# brew-web

`mDNS` drivers need to be installed on system running go server code. They are used to discover esp

Connects to esp running in network(can be anything else with http endpoints).

Temperature is taken every 5 seconds.

Go code uses go modules so go version of 1.11 or higher is required.

directory ./espKegFirmware contains arduino project source code for esp.

## Go Server for brewing can be run in two modes

- prod mode, connects to esp or other device discovered by msdn

```bash
go run main.go -type=esp
```

- mock mode, mocks connections to esp

```bash
go run main.go -type=mock
```

## React UI

UI is written in react and can be built with following command:

```bash
webpack --progress
```

## TODOS

- Nice ui
- error handling
- finish recipes

## Static assets

Binary contains static assets loaded through `https://github.com/rakyll/statik`

install

```
go get github.com/rakyll/statik
npm i --prefix web-ui
npm run build --prefix web-ui
statik -f -src=web-ui/build
```