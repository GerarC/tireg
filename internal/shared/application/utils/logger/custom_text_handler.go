package logger

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"sync"
)

type CustomTextHandler struct {
	opts  slog.HandlerOptions
	out   io.Writer
	mu    *sync.Mutex
	attrs []slog.Attr
	group string
}

func NewCustomTextHandler(out io.Writer, opts *slog.HandlerOptions) *CustomTextHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	return &CustomTextHandler{
		opts: *opts,
		out:  out,
		mu:   &sync.Mutex{},
	}
}

func (h *CustomTextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *CustomTextHandler) Handle(ctx context.Context, r slog.Record) error {
	var buf bytes.Buffer

	timeAttr := slog.Time(slog.TimeKey, r.Time)
	if h.opts.ReplaceAttr != nil {
		timeAttr = h.opts.ReplaceAttr(nil, timeAttr)
	}

	if timeAttr.Key != "" {
		buf.WriteString(timeAttr.Value.String())
	}

	buf.WriteString(" [")
	buf.WriteString(r.Level.String())
	buf.WriteString("]: ")
	buf.WriteString(r.Message)

	for _, attr := range h.attrs {
		h.appendAttr(&buf, attr)
	}
	r.Attrs(func(a slog.Attr) bool {
		h.appendAttr(&buf, a)
		return true
	})

	buf.WriteString("\n")

	h.mu.Lock()
	defer h.mu.Unlock()

	_, err := h.out.Write(buf.Bytes())
	return err
}

func (h *CustomTextHandler) appendAttr(buf *bytes.Buffer, a slog.Attr) {
	buf.WriteString(" ")
	if h.group != "" {
		buf.WriteString(h.group)
		buf.WriteString(".")
	}
	buf.WriteString(a.Key)
	buf.WriteString("=")
	buf.WriteString(a.Value.String())
}

func (h *CustomTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := *h
	newHandler.attrs = append(h.attrs, attrs...)
	return &newHandler
}

func (h *CustomTextHandler) WithGroup(name string) slog.Handler {
	newHandler := *h
	if newHandler.group == "" {
		newHandler.group = name
	} else {
		newHandler.group = newHandler.group + "." + name
	}

	return &newHandler
}
