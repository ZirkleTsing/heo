package uncore

type MemoryController struct {
	*BaseController
	NumReads int32
	NumWrites int32
}

func NewMemoryController(memoryHierarchy *MemoryHierarchy) *MemoryController {
	var memoryController = &MemoryController{
	}

	//TODO

	return memoryController
}