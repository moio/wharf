package gbrotli

import (
	"io"

	"github.com/dsnet/compress/brotli"
	"github.com/itchio/wharf/pwr"
)

type brotliDecompressor struct{}

func (bc *brotliDecompressor) Apply(reader io.Reader) (io.Reader, error) {
	br := brotli.NewReader(reader, nil)
	return br, nil
}

func init() {
	pwr.RegisterDecompressor(pwr.CompressionAlgorithm_BROTLI, &brotliDecompressor{})
}
