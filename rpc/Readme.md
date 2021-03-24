# RPC

this example is to demonstrate how RPC may work.

it can work over several transports:
- tcp
- http

not yet implemented:
- udp
- websocket
- (tls,mtls,https)?


Example Usage:
---
- `$ ./rpc <mode> -<transport> :<port>`
- `$ ./rpc server -tcp :1234`
- `$ ./rpc client -tcp :1234 <<< "ping"`
- `$ ./rpc client -tcp :1234 <<< "put <key> <value>"`
- `$ ./rpc client -tcp :1234 <<< "get <key>"`
- `$ ./rpc client -tcp :1234 <<< "del <key>"`
- `$ ./rpc client -tcp :1234 <<< "ls"`
- `$ ./rpc server -http :1234`
- `$ ./rpc client -http :1234 <<< "ping"`



