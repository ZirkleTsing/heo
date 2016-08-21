package mem

type CircularByteBuffer struct {
	Data          *SimpleMemory
	ReadPosition  uint64
	WritePosition uint64
}

func NewCircularByteBuffer(capacity uint64) *CircularByteBuffer {
	var circularByteBuffer = &CircularByteBuffer{
		Data:NewSimpleMemory(true, make([]byte, capacity)),
	}

	return circularByteBuffer
}

func (circularByteBuffer *CircularByteBuffer) Reset() {
	circularByteBuffer.ReadPosition = 0
	circularByteBuffer.WritePosition = 0
}

func (circularByteBuffer *CircularByteBuffer) Read(dest *[]byte, offset uint64, length uint64) (uint64) {
	if length >= uint64(len(circularByteBuffer.Data.Data)) {
		panic("Requested read is greater than the buffer")
	}

	if circularByteBuffer.WritePosition == circularByteBuffer.ReadPosition {
		return 0
	}

	circularByteBuffer.Data.ReadPosition = circularByteBuffer.ReadPosition
	if circularByteBuffer.WritePosition < circularByteBuffer.ReadPosition {
		var remainder = uint64(len(circularByteBuffer.Data.Data)) - circularByteBuffer.Data.ReadPosition
		if remainder < length {
			copy((*dest)[offset:remainder], circularByteBuffer.Data.ReadBlock(remainder))

			offset += remainder
			length -= remainder

			circularByteBuffer.ReadPosition = 0
			circularByteBuffer.Data.ReadPosition = circularByteBuffer.ReadPosition

			var space = circularByteBuffer.WritePosition - circularByteBuffer.ReadPosition
			if space <= length {
				length = space
			}

			copy((*dest)[offset:length], circularByteBuffer.Data.ReadBlock(length))
			circularByteBuffer.ReadPosition += length

			return remainder + length
		} else {
			copy((*dest)[offset:remainder], circularByteBuffer.Data.ReadBlock(remainder))
			circularByteBuffer.ReadPosition += remainder
			return remainder
		}
	} else {
		var space = circularByteBuffer.WritePosition - circularByteBuffer.ReadPosition
		if space <= length {
			length = space
		}

		copy((*dest)[offset:length], circularByteBuffer.Data.ReadBlock(length))
		circularByteBuffer.ReadPosition += length
		return length
	}
}

func (circularByteBuffer *CircularByteBuffer) Write(src *[]byte, offset uint64, length uint64) bool {
	if length >= uint64(len(circularByteBuffer.Data.Data)) {
		panic("Requested write is greater than the buffer")
	}

	circularByteBuffer.Data.WritePosition = circularByteBuffer.WritePosition

	if (circularByteBuffer.ReadPosition <= circularByteBuffer.WritePosition &&
		circularByteBuffer.WritePosition + length < uint64(len(circularByteBuffer.Data.Data))) ||
		(circularByteBuffer.WritePosition < circularByteBuffer.ReadPosition &&
			length < circularByteBuffer.ReadPosition - circularByteBuffer.WritePosition) {
		circularByteBuffer.Data.WriteBlock(length, (*src)[offset:length])
		circularByteBuffer.WritePosition += length
		return true
	} else {
		var remainder = uint64(len(circularByteBuffer.Data.Data)) - circularByteBuffer.Data.ReadPosition

		if circularByteBuffer.ReadPosition < circularByteBuffer.WritePosition &&
			length > circularByteBuffer.ReadPosition + remainder {
			return false
		}

		if circularByteBuffer.WritePosition < circularByteBuffer.ReadPosition &&
			length > circularByteBuffer.ReadPosition - circularByteBuffer.WritePosition {
			return false
		}

		circularByteBuffer.Data.WriteBlock(remainder, (*src)[offset:remainder])

		offset += remainder
		length -= remainder

		circularByteBuffer.WritePosition = 0
		circularByteBuffer.Data.WritePosition = circularByteBuffer.WritePosition

		if length >= circularByteBuffer.ReadPosition {
			panic("There is not enough room for this write operation")
		}

		circularByteBuffer.Data.WriteBlock(length, (*src)[offset:length])
		circularByteBuffer.WritePosition += length

		return true
	}
}

