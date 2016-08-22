package cpu

type MemoryHierarchy struct {
	Experiment *CPUExperiment
}

func NewMemoryHierarchy(experiment *CPUExperiment) *MemoryHierarchy {
	var memoryHierarchy = &MemoryHierarchy{
		Experiment:experiment,
	}

	return memoryHierarchy
}
