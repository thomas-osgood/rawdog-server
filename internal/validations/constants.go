package validations

// slice containing unix socket types.
var UNIXSOCKET_CONTYPES []string = []string{"unix", "unixpacket"}

// slice containing a list of valid connection
// types the user can specify when creating a
// new teamserver.
var VALID_CONNTYPES []string = []string{"tcp", "tcp4", "tcp6", "unix", "unixpacket"}
