package tls_example

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/joomcode/errorx"
)

// dialHandler specifies the dial function for creating unencrypted TCP connections.
type dialHandler func(ctx context.Context, network, addr string) (net.Conn, error)

func createTlsConnection() (*tls.Conn, error) {
	d := createDialContext("1.1.1.1:853", 5*time.Second)
	t := tls.Config{ServerName: "1.1.1.1"}
	rawConn, err := d(context.TODO(), "tcp", "")
	if err != nil {
		return nil, err
	}
	conn := tls.Client(rawConn, &t)
	err = conn.Handshake()
	if err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}

func createDialContext(address string, timeout time.Duration) (dialContext dialHandler) {
	dialer := &net.Dialer{
		Timeout:   timeout,
		DualStack: true,
	}

	dialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		con, err := dialer.DialContext(ctx, network, address)
		if err == nil {
			return con, err
		}

		return nil, errorx.Decorate(err, "dialer failed to initialize connection")
	}
	return
}
