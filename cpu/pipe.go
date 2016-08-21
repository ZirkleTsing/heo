package cpu

type Pipe struct {
	FileDescriptors []int
	Buffer *CircularByteBuffer
}

func NewPipe(fileDescriptors []int) *Pipe {
	var pipe = &Pipe{
		FileDescriptors:fileDescriptors,
		Buffer:NewCircularByteBuffer(1024),
	}

	return pipe
}
