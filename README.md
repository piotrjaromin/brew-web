# brew-web

mdsn drivers needs to be installed on host system

Connectes to esp runing in network(can be anything else with http ednpoints)

Temperature is taken every 5 seconds

directory ./espKegFirmware contains arduino project source code for esp

1. prod mode
```bash
go run main.go -type=esp
```

2. mock mode
```bash
go run main.go -type=mock
```

3. to build ui
```bash
webpack --progress
```

TODO Nice ui
TODO error handling
TODO finish recepies