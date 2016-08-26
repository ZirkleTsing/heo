package mem

type CircularByteBuffer struct {
	Data          *SimpleMemory
	ReadPosition  uint32
	WritePosition uint32
}

func NewCircularByteBuffer(capacity uint32) *CircularByteBuffer {
	var circularByteBuffer = &CircularByteBuffer{
		Data:NewSimpleMemory(true, make([]byte, capacity)),
	}

	return circularByteBuffer
}

func (circularByteBuffer *CircularByteBuffer) Reset() {
	circularByteBuffer.ReadPosition = 0
	circularByteBuffer.WritePosition = 0
}

func (circularByteBuffer *CircularByteBuffer) Read(dest *[]byte, count uint32) (uint32) {
	var offset = uint32(0)

	if count >= uint32(len(circularByteBuffer.Data.Data)) {
		panic("Requested read is greater than the buffer")
	}

	if circularByteBuffer.WritePosition == circularByteBuffer.ReadPosition {
		return 0
	}

	circularByteBuffer.Data.ReadPosition = circularByteBuffer.ReadPosition
	if circularByteBuffer.WritePosition < circularByteBuffer.ReadPosition {
		var remainder = uint32(len(circularByteBuffer.Data.Data)) - circularByteBuffer.Data.ReadPosition
		if remainder < count {
			copy((*dest)[offset:remainder], circularByteBuffer.Data.ReadBlock(remainder))

			offset += remainder
			count -= remainder

			circularByteBuffer.ReadPosition = 0
			circularByteBuffer.Data.ReadPosition = circularByteBuffer.ReadPosition

			var space = circularByteBuffer.WritePosition - circularByteBuffer.ReadPosition
			if space <= count {
				count = space
			}

			copy((*dest)[offset:count], circularByteBuffer.Data.ReadBlock(count))
			circularByteBuffer.ReadPosition += count

			return remainder + count
		} else {
			copy((*dest)[offset:remainder], circularByteBuffer.Data.ReadBlock(remainder))
			circularByteBuffer.ReadPosition += remainder
			return remainder
		}
	} else {
		var space = circularByteBuffer.WritePosition - circularByteBuffer.ReadPosition
		if space <= count {
			count = space
		}

		copy((*dest)[offset:count], circularByteBuffer.Data.ReadBlock(count))
		circularByteBuffer.ReadPosition += count
		return count
	}
}

func (circularByteBuffer *CircularByteBuffer) Write(src *[]byte, count uint32) bool {
	var offset = uint32(0)

	if count >= uint32(len(circularByteBuffer.Data.Data)) {
		panic("Requested write is greater than the buffer")
	}

	circularByteBuffer.Data.WritePosition = circularByteBuffer.WritePosition

	if (circularByteBuffer.ReadPosition <= circularByteBuffer.WritePosition &&
		circularByteBuffer.WritePosition + count < uint32(len(circularByteBuffer.Data.Data))) ||
		(circularByteBuffer.WritePosition < circularByteBuffer.ReadPosition &&
			count < circularByteBuffer.ReadPosition - circularByteBuffer.WritePosition) {
		circularByteBuffer.Data.WriteBlock(count, (*src)[offset:count])
		circularByteBuffer.WritePosition += count
		return true
	} else {
		var remainder = uint32(len(circularByteBuffer.Data.Data)) - circularByteBuffer.Data.ReadPosition

		if circularByteBuffer.ReadPosition < circularByteBuffer.WritePosition &&
			count > circularByteBuffer.ReadPosition + remainder {
			return false
		}

		if circularByteBuffer.WritePosition < circularByteBuffer.ReadPosition &&
			count > circularByteBuffer.ReadPosition - circularByteBuffer.WritePosition {
			return false
		}

		circularByteBuffer.Data.WriteBlock(remainder, (*src)[offset:remainder])

		offset += remainder
		count -= remainder

		circularByteBuffer.WritePosition = 0
		circularByteBuffer.Data.WritePosition = circularByteBuffer.WritePosition

		if count >= circularByteBuffer.ReadPosition {
			panic("There is not enough room for this write operation")
		}

		circularByteBuffer.Data.WriteBlock(count, (*src)[offset:count])
		circularByteBuffer.WritePosition += count

		return true
	}
}

func (circularByteBuffer *CircularByteBuffer) IsEmpty() bool {
	return circularByteBuffer.WritePosition == circularByteBuffer.ReadPosition
}

func (circularByteBuffer *CircularByteBuffer) IsFull() bool {
	return circularByteBuffer.WritePosition + 1 <= uint32(len(circularByteBuffer.Data.Data)) &&
		circularByteBuffer.WritePosition + 1 == circularByteBuffer.ReadPosition ||
		circularByteBuffer.WritePosition == uint32(len(circularByteBuffer.Data.Data)) - 1 &&
			circularByteBuffer.ReadPosition == 0
}
