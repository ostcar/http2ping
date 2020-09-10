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

Per default, it sends a http2 ping every 5 seconds. The wait time can be set in
seconds with the `-wait` flag. The example sends a ping once per minute.

```
http2ping -w 60 https://localhost:9000/system/autoupdate
```