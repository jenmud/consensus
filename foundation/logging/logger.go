package logging

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type Options struct {
	// Level is the minimum level that will be logged.
	Level slog.Leveler

	// Subject is the NATS subject/topic to publish messages to.
	Subject string

	// NATSAddr is the address of the NATS server.
	NATSAddr string
}

// NATSHandler is a slog.Handler that publishes messages to NATS.
type NATSHandler struct {
	lock sync.RWMutex
	opts Options
	nc   *nats.Conn
}

// NewNATSHandler returns a new NATSHandler.
func NewNATSHandler(ctx context.Context, opt Options) slog.Handler {
	if opt.NATSAddr == "" {
		opt.NATSAddr = nats.DefaultURL
	}

	if opt.Subject == "" {
		opt.Subject = "consensus.logs"
	}

	nc, _ := nats.Connect(opt.NATSAddr, nats.Name("consensus"), nats.Timeout(time.Second))

	handler := &NATSHandler{nc: nc, opts: opt}

	if handler.opts.Level == nil {
		handler.opts.Level = slog.LevelDebug
	}

	return handler
}

// WithNATS returns a new NATSHandler with the given NATS connection.
func (h *NATSHandler) WithNC(nc *nats.Conn) *NATSHandler {
	h.nc = nc
	return h
}

// Enabled returns true if the handler is enabled for the given level.
func (h *NATSHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= h.opts.Level.Level()
}

// WithGroup returns a new NATSHandler with the given group name.
func (h *NATSHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return h
}

// WithAttrs returns a new NATSHandler with the given attributes.
func (h *NATSHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// TODO: implement
	return h
}

func (h *NATSHandler) Handle(ctx context.Context, rec slog.Record) error {
	dump, err := json.Marshal(rec)

	if err != nil {
		return err
	}

	h.lock.Lock()
	defer h.lock.Unlock()

	if h.nc != nil {
		h.nc.Publish(h.opts.Subject, dump)
	}

	_, err = os.Stdout.Write(dump)
	return err
}
