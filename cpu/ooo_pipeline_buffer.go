package cpu

type PipelineBuffer struct {
	Size    uint32
	Entries []interface{}
}

func NewPipelineBuffer(size uint32) *PipelineBuffer {
	var pipelineBuffer = &PipelineBuffer{
		Size:size,
	}

	return pipelineBuffer
}

func (pipelineBuffer *PipelineBuffer) Count() uint32 {
	return uint32(len(pipelineBuffer.Entries))
}

func (pipelineBuffer *PipelineBuffer) Full() bool {
	return pipelineBuffer.Count() >= pipelineBuffer.Size
}

func (pipelineBuffer *PipelineBuffer) Empty() bool {
	return pipelineBuffer.Count() == 0
}