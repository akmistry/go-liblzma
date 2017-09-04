package xz

import "sync"

var defaultBufferPool = sync.Pool{New: func() interface{} { return make([]byte, DefaultBufsize) }}

func getBuffer(size int) []byte {
	if size == DefaultBufsize {
		buf := defaultBufferPool.Get().([]byte)
		for i := range buf {
			buf[i] = 0
		}
	}
	return make([]byte, size)
}

func putBuffer(buf []byte) {
	if len(buf) == DefaultBufsize {
		defaultBufferPool.Put(buf)
	}
}
