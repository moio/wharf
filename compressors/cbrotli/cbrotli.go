package cbrotli

import (
	"github.com/andybalholm/brotli"
	"io"

	"github.com/itchio/wharf/pwr"
)

type brotliCompressor struct{}

func (bc *brotliCompressor) Apply(writer io.Writer, quality int32) (io.Writer, error) {
	return brotli.NewWriterLevel(writer, int(quality)), nil
}

func init() {
	pwr.RegisterCompressor(pwr.CompressionAlgorithm_BROTLI, &brotliCompressor{})
}
