package snet

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/quic-go/quic-go"
)

func process() {
	ctx, cancelF := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancelF()

	conf := &quic.Config{}
	tlsConf := &tls.Config{}
	quic.DialAddr(ctx, "", tlsConf, conf)
}
