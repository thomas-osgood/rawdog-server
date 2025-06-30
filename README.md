# Rawdog: Server

## Overview

This module defines a generic server object that implements the Rawdog TCP Communication protocol.

```golang

import (
    rdserver "github.com/thomas-osgood/rawdog-server"
)

/*
    function designed to echo back the data that was
    transmitted by the client.
*/
func TestEndpoint(c net.Conn, md comms.TcpHeader, data *bytes.Buffer) (result string, err error) {
        return data.String(), nil
}

/*
    main function that creates a Rawdog TeamServer,
    assigns a new endpoint/func, and starts the server
    listening on port 8080 (default).
*/
func main() {
        srv, e := rdserver.NewTeamServer()
        if e != nil {
                log.Fatalf(e.Error())
        }

        srv.AddEndpoint(1, TestEndpoint)

        log.Fatal(srv.Start())
}
```
