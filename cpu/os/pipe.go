package os

import "github.com/mcai/acogo/cpu/mem"

type Pipe struct {
	FileDescriptors []int
	Buffer *mem.CircularByteBuffer
}

func NewPipe(fileDescriptors []int) *Pipe {
	var pipe = &Pipe{
		FileDescriptors:fileDescriptors,
		Buffer:mem.NewCircularByteBuffer(1024),
	}

	return pipe
}
