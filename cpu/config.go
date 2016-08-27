package cpu

import "github.com/mcai/acogo/simutil"

type CPUConfig struct {
	OutputDirectory   string
	ContextMappings   []*ContextMapping
	MaxDynamicInsts   int32
	NumCores          int32
	NumThreadsPerCore int32
}

func NewCPUConfig(outputDirectory string) *CPUConfig {
	var cpuConfig = &CPUConfig{
		OutputDirectory:outputDirectory,
		MaxDynamicInsts:-1,
		NumCores:2,
		NumThreadsPerCore:2,
	}

	return cpuConfig
}

func (experiment *CPUExperiment) DumpConfig() {
	simutil.WriteJsonFile(experiment.Config, experiment.Config.OutputDirectory, simutil.CONFIG_JSON_FILE_NAME)
}