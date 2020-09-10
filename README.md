# HTTP2 Ping

http2ping is a helper tool to test http2 pings from a client to the server.

It only supports connections with https, but does not verify the certificate. So
it can be used with selfsigned certificates.

## Install

The tool can be installed with go:

```
go get github.com/ostcar/http2ping
```

## Usage

The tools requries an https addres. For example:

```
http2ping https://localhost:9000/system/autoupdate
```