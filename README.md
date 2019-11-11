# GoTee

Duplicate TCP stream. Listen for TCP connection, and forward stream to output port 1. When output port 2 is also given,
dupliacte the stream to output port 2.


## Usage

```sh
usage: gotee [<flags>]

Flags:
      --help                  Show context-sensitive help (also try --help-long and --help-man).
  -l, --listen=":8000"        Listen port.
  -1, --output port1=":8001"  Output port 1.
  -2, --output port2=":8002"  Output port 2.
  -d, --debug                 Debug.
```