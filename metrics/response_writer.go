package metrics

import rw "github.com/shaardie/gomiddleware/pkg/responsewriter"

type responsewriter struct {
	rw.ResponseWriter
	size int
}

func (rw *responsewriter) Write(b []byte) (int, error) {
	rw.size += len(b)
	return rw.ResponseWriter.Write(b)
}
