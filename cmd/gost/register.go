package main

import (
	// Register connectors
	_ "github.com/liukeqqs/x/connector/direct"
	_ "github.com/liukeqqs/x/connector/forward"
	_ "github.com/liukeqqs/x/connector/http"
	_ "github.com/liukeqqs/x/connector/http2"
	_ "github.com/liukeqqs/x/connector/relay"
	_ "github.com/liukeqqs/x/connector/serial"
	_ "github.com/liukeqqs/x/connector/sni"
	_ "github.com/liukeqqs/x/connector/socks/v4"
	_ "github.com/liukeqqs/x/connector/socks/v5"
	_ "github.com/liukeqqs/x/connector/ss"
	_ "github.com/liukeqqs/x/connector/ss/udp"
	_ "github.com/liukeqqs/x/connector/sshd"
	_ "github.com/liukeqqs/x/connector/tcp"
	_ "github.com/liukeqqs/x/connector/tunnel"
	_ "github.com/liukeqqs/x/connector/unix"

	// Register dialers
	_ "github.com/liukeqqs/x/dialer/direct"
	_ "github.com/liukeqqs/x/dialer/dtls"
	_ "github.com/liukeqqs/x/dialer/ftcp"
	_ "github.com/liukeqqs/x/dialer/grpc"
	_ "github.com/liukeqqs/x/dialer/http2"
	_ "github.com/liukeqqs/x/dialer/http2/h2"
	_ "github.com/liukeqqs/x/dialer/http3"
	_ "github.com/liukeqqs/x/dialer/http3/wt"
	_ "github.com/liukeqqs/x/dialer/icmp"
	_ "github.com/liukeqqs/x/dialer/kcp"
	_ "github.com/liukeqqs/x/dialer/mtcp"
	_ "github.com/liukeqqs/x/dialer/mtls"
	_ "github.com/liukeqqs/x/dialer/mws"
	_ "github.com/liukeqqs/x/dialer/obfs/http"
	_ "github.com/liukeqqs/x/dialer/obfs/tls"
	_ "github.com/liukeqqs/x/dialer/pht"
	_ "github.com/liukeqqs/x/dialer/quic"
	_ "github.com/liukeqqs/x/dialer/serial"
	_ "github.com/liukeqqs/x/dialer/ssh"
	_ "github.com/liukeqqs/x/dialer/sshd"
	_ "github.com/liukeqqs/x/dialer/tcp"
	_ "github.com/liukeqqs/x/dialer/tls"
	_ "github.com/liukeqqs/x/dialer/udp"
	_ "github.com/liukeqqs/x/dialer/unix"
	_ "github.com/liukeqqs/x/dialer/ws"

	// Register handlers
	_ "github.com/liukeqqs/x/handler/auto"
	_ "github.com/liukeqqs/x/handler/dns"
	_ "github.com/liukeqqs/x/handler/file"
	_ "github.com/liukeqqs/x/handler/forward/local"
	_ "github.com/liukeqqs/x/handler/forward/remote"
	_ "github.com/liukeqqs/x/handler/http"
	_ "github.com/liukeqqs/x/handler/http2"
	_ "github.com/liukeqqs/x/handler/http3"
	_ "github.com/liukeqqs/x/handler/metrics"
	_ "github.com/liukeqqs/x/handler/redirect/tcp"
	_ "github.com/liukeqqs/x/handler/redirect/udp"
	_ "github.com/liukeqqs/x/handler/relay"
	_ "github.com/liukeqqs/x/handler/serial"
	_ "github.com/liukeqqs/x/handler/sni"
	_ "github.com/liukeqqs/x/handler/socks/v4"
	_ "github.com/liukeqqs/x/handler/socks/v5"
	_ "github.com/liukeqqs/x/handler/ss"
	_ "github.com/liukeqqs/x/handler/ss/udp"
	_ "github.com/liukeqqs/x/handler/sshd"
	_ "github.com/liukeqqs/x/handler/tap"
	_ "github.com/liukeqqs/x/handler/tun"
	_ "github.com/liukeqqs/x/handler/tunnel"
	_ "github.com/liukeqqs/x/handler/unix"

	// Register listeners
	_ "github.com/liukeqqs/x/listener/dns"
	_ "github.com/liukeqqs/x/listener/dtls"
	_ "github.com/liukeqqs/x/listener/ftcp"
	_ "github.com/liukeqqs/x/listener/grpc"
	_ "github.com/liukeqqs/x/listener/http2"
	_ "github.com/liukeqqs/x/listener/http2/h2"
	_ "github.com/liukeqqs/x/listener/http3"
	_ "github.com/liukeqqs/x/listener/http3/h3"
	_ "github.com/liukeqqs/x/listener/http3/wt"
	_ "github.com/liukeqqs/x/listener/icmp"
	_ "github.com/liukeqqs/x/listener/kcp"
	_ "github.com/liukeqqs/x/listener/mtcp"
	_ "github.com/liukeqqs/x/listener/mtls"
	_ "github.com/liukeqqs/x/listener/mws"
	_ "github.com/liukeqqs/x/listener/obfs/http"
	_ "github.com/liukeqqs/x/listener/obfs/tls"
	_ "github.com/liukeqqs/x/listener/pht"
	_ "github.com/liukeqqs/x/listener/quic"
	_ "github.com/liukeqqs/x/listener/redirect/tcp"
	_ "github.com/liukeqqs/x/listener/redirect/udp"
	_ "github.com/liukeqqs/x/listener/rtcp"
	_ "github.com/liukeqqs/x/listener/rudp"
	_ "github.com/liukeqqs/x/listener/serial"
	_ "github.com/liukeqqs/x/listener/ssh"
	_ "github.com/liukeqqs/x/listener/sshd"
	_ "github.com/liukeqqs/x/listener/tap"
	_ "github.com/liukeqqs/x/listener/tcp"
	_ "github.com/liukeqqs/x/listener/tls"
	_ "github.com/liukeqqs/x/listener/tun"
	_ "github.com/liukeqqs/x/listener/udp"
	_ "github.com/liukeqqs/x/listener/unix"
	_ "github.com/liukeqqs/x/listener/ws"
)
