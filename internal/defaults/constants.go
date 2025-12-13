package defaults

import "time"

// default address and port the TeamServer will
// listen for incoming connections on.
const DEFAULT_ADDRESS string = "0.0.0.0:8080"

// default connection type the teamserver will listen
// for if none is specified by the user.
const DEFAULT_CONNTYPE string = "tcp"

// default read timeout for server.
const DEFAULT_READ_TIMEOUT time.Duration = 15 * time.Second

// default send timeout for server.
const DEFAULT_SEND_TIMEOUT time.Duration = 5 * time.Second

var VALID_CONNTYPES []string = []string{"tcp", "tcp4", "tcp6", "unix", "unixpacket"}
