package cpu

type PipelineBuffer struct {
	Capacity uint32
	Entries  []interface{}
}

func NewPipelineBuffer(capacity uint32) *PipelineBuffer {
	var pipelineBuffer = &PipelineBuffer{
		Capacity:capacity,
	}

	return pipelineBuffer
}

func (pipelineBuffer *PipelineBuffer) Full() bool {
	return uint32(len(pipelineBuffer.Entries)) >= pipelineBuffer.Capacity
}

func (pipelineBuffer *PipelineBuffer) Empty() bool {
	return len(pipelineBuffer.Entries) == 0
}