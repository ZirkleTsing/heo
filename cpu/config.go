package cpu

import "github.com/mcai/acogo/simutil"

type CPUConfig struct {
	OutputDirectory           string

	ContextMappings           []*ContextMapping

	MaxDynamicInsts           int32

	NumCores                  int32
	NumThreadsPerCore         int32

	PhysicalRegisterFileSize  uint32

	DecodeWidth               uint32
	IssueWidth                uint32
	CommitWidth               uint32

	DecodeBufferSize          uint32
	ReorderBufferSize         uint32
	LoadStoreQueueSize        uint32

	BranchPredictorSize       uint32
	BranchTargetBufferNumSets uint32
	BranchTargetBufferAssoc   uint32
	ReturnAddressStackSize    uint32
}

func NewCPUConfig(outputDirectory string) *CPUConfig {
	var config = &CPUConfig{
		OutputDirectory:outputDirectory,

		MaxDynamicInsts:-1,

		NumCores:2,
		NumThreadsPerCore:2,

		PhysicalRegisterFileSize:128,

		DecodeWidth:4,
		IssueWidth:4,
		CommitWidth:4,

		DecodeBufferSize:96,
		ReorderBufferSize: 96,
		LoadStoreQueueSize:48,

		BranchPredictorSize:2048,
		BranchTargetBufferNumSets:512,
		BranchTargetBufferAssoc:4,
		ReturnAddressStackSize:8,
	}

	return config
}

func (config *CPUConfig) Dump(outputDirectory string) {
	simutil.WriteJsonFile(config, outputDirectory, simutil.CPU_CONFIG_JSON_FILE_NAME)
}