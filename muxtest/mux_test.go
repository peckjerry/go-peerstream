package muxtest

import (
	"testing"

	multiplex "github.com/jbenet/go-peerstream/Godeps/_workspace/src/github.com/jbenet/go-stream-muxer/multiplex"
	multistream "github.com/jbenet/go-peerstream/Godeps/_workspace/src/github.com/jbenet/go-stream-muxer/multistream"
	muxado "github.com/jbenet/go-peerstream/Godeps/_workspace/src/github.com/jbenet/go-stream-muxer/muxado"
	spdy "github.com/jbenet/go-peerstream/Godeps/_workspace/src/github.com/jbenet/go-stream-muxer/spdystream"
	yamux "github.com/jbenet/go-peerstream/Godeps/_workspace/src/github.com/jbenet/go-stream-muxer/yamux"
)

func TestYamuxTransport(t *testing.T) {
	SubtestAll(t, yamux.DefaultTransport)
}

func TestSpdyStreamTransport(t *testing.T) {
	SubtestAll(t, spdy.Transport)
}

func TestMultiplexTransport(t *testing.T) {
	SubtestAll(t, multiplex.DefaultTransport)
}

func TestMuxadoTransport(t *testing.T) {
	SubtestAll(t, muxado.Transport)
}

func TestMultistreamTransport(t *testing.T) {
	SubtestAll(t, multistream.NewTransport())
}
