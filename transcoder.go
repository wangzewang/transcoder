package transcoder

import (
	"io"
	"os"
)

// Transcoder ...
type Transcoder interface {
	Start(opts Options) (<-chan Progress, *os.Process, error)
	Input(i string) Transcoder
	InputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	Output(o string) Transcoder
	OutputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	WithOptions(opts Options) Transcoder
	WithAdditionalOptions(opts Options) Transcoder
	GetMetadata() (Metadata, error)
}
