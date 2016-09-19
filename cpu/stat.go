package cpu

import (
	"fmt"
	"github.com/mcai/acogo/simutil"
)

func (experiment *CPUExperiment) DumpStats() {
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "SimulationTime",
		Value: fmt.Sprintf("%v", experiment.SimulationTime()),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "SimulationTimeInSeconds",
		Value: experiment.SimulationTime().Seconds(),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "TotalCycles",
		Value: experiment.CycleAccurateEventQueue().CurrentCycle,
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"NumDynamicInsts",
		Value:experiment.Processor.NumDynamicInsts(),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "CyclesPerSecond",
		Value: experiment.CyclesPerSecond(),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "InstructionsPerSecond",
		Value: experiment.InstructionsPerSecond(),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"InstructionsPerCycle",
		Value:experiment.Processor.InstructionsPerCycle(),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"CyclesPerInstructions",
		Value:experiment.Processor.CyclesPerInstructions(),
	})

	for _, core := range experiment.Processor.Cores {
		for _, thread := range core.Threads() {
			experiment.Stats = append(experiment.Stats, simutil.Stat{
				Key:fmt.Sprintf("thread_%d.NumDynamicInsts", thread.Id()),
				Value:thread.NumDynamicInsts(),
			})

			experiment.Stats = append(experiment.Stats, simutil.Stat{
				Key:fmt.Sprintf("thread_%d.InstructionsPerCycle", thread.Id()),
				Value:thread.InstructionsPerCycle(),
			})

			experiment.Stats = append(experiment.Stats, simutil.Stat{
				Key:fmt.Sprintf("thread_%d.CyclesPerInstructions", thread.Id()),
				Value:thread.CyclesPerInstructions(),
			})

			if oooThread := thread.(*OoOThread); oooThread != nil {
				experiment.Stats = append(experiment.Stats, simutil.Stat{
					Key:fmt.Sprintf("thread_%d.BranchPredictor.HitRatio", thread.Id()),
					Value:oooThread.BranchPredictor.HitRatio(),
				})
				experiment.Stats = append(experiment.Stats, simutil.Stat{
					Key:fmt.Sprintf("thread_%d.BranchPredictor.NumAccesses", thread.Id()),
					Value:oooThread.BranchPredictor.NumAccesses(),
				})
				experiment.Stats = append(experiment.Stats, simutil.Stat{
					Key:fmt.Sprintf("thread_%d.BranchPredictor.NumHits", thread.Id()),
					Value:oooThread.BranchPredictor.NumHits(),
				})
				experiment.Stats = append(experiment.Stats, simutil.Stat{
					Key:fmt.Sprintf("thread_%d.BranchPredictor.NumMisses", thread.Id()),
					Value:oooThread.BranchPredictor.NumMisses(),
				})
			}
		}
	}

	for i, itlb := range experiment.MemoryHierarchy.ITlbs() {
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("itlb_%d.HitRatio", i),
			Value:itlb.HitRatio(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("itlb_%d.NumAccesses", i),
			Value:itlb.NumAccesses(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("itlb_%d.NumHits", i),
			Value:itlb.NumHits,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("itlb_%d.NumMisses", i),
			Value:itlb.NumMisses,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("itlb_%d.NumEvictions", i),
			Value:itlb.NumEvictions,
		})
	}

	for i, dtlb := range experiment.MemoryHierarchy.DTlbs() {
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dtlb_%d.HitRatio", i),
			Value:dtlb.HitRatio(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dtlb_%d.NumAccesses", i),
			Value:dtlb.NumAccesses(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dtlb_%d.NumHits", i),
			Value:dtlb.NumHits,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dtlb_%d.NumMisses", i),
			Value:dtlb.NumMisses,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dtlb_%d.NumEvictions", i),
			Value:dtlb.NumEvictions,
		})
	}

	for i, cacheController := range experiment.MemoryHierarchy.L1IControllers() {
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.HitRatio", i),
			Value:cacheController.HitRatio(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumDownwardAccesses", i),
			Value:cacheController.NumDownwardAccesses(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumDownwardHits", i),
			Value:cacheController.NumDownwardHits(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumDownwardMisses", i),
			Value:cacheController.NumDownwardMisses(),
		})

		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumDownwardReadHits", i),
			Value:cacheController.NumDownwardReadHits,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumDownwardReadMisses", i),
			Value:cacheController.NumDownwardReadMisses,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumDownwardWriteHits", i),
			Value:cacheController.NumDownwardWriteHits,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumDownwardWriteMisses", i),
			Value:cacheController.NumDownwardWriteMisses,
		})

		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("icache_%d.NumEvictions", i),
			Value:cacheController.NumEvictions,
		})
	}

	for i, cacheController := range experiment.MemoryHierarchy.L1DControllers() {
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.HitRatio", i),
			Value:cacheController.HitRatio(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumDownwardAccesses", i),
			Value:cacheController.NumDownwardAccesses(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumDownwardHits", i),
			Value:cacheController.NumDownwardHits(),
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumDownwardMisses", i),
			Value:cacheController.NumDownwardMisses(),
		})

		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumDownwardReadHits", i),
			Value:cacheController.NumDownwardReadHits,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumDownwardReadMisses", i),
			Value:cacheController.NumDownwardReadMisses,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumDownwardWriteHits", i),
			Value:cacheController.NumDownwardWriteHits,
		})
		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumDownwardWriteMisses", i),
			Value:cacheController.NumDownwardWriteMisses,
		})

		experiment.Stats = append(experiment.Stats, simutil.Stat{
			Key:fmt.Sprintf("dcache_%d.NumEvictions", i),
			Value:cacheController.NumEvictions,
		})
	}

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.HitRatio",
		Value:experiment.MemoryHierarchy.L2Controller().HitRatio(),
	})
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumDownwardAccesses",
		Value:experiment.MemoryHierarchy.L2Controller().NumDownwardAccesses(),
	})
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumDownwardHits",
		Value:experiment.MemoryHierarchy.L2Controller().NumDownwardHits(),
	})
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumDownwardMisses",
		Value:experiment.MemoryHierarchy.L2Controller().NumDownwardMisses(),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumDownwardReadHits",
		Value:experiment.MemoryHierarchy.L2Controller().NumDownwardReadHits,
	})
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumDownwardReadMisses",
		Value:experiment.MemoryHierarchy.L2Controller().NumDownwardReadMisses,
	})
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumDownwardWriteHits",
		Value:experiment.MemoryHierarchy.L2Controller().NumDownwardWriteHits,
	})
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumDownwardWriteMisses",
		Value:experiment.MemoryHierarchy.L2Controller().NumDownwardWriteMisses,
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"l2cache.NumEvictions",
		Value:experiment.MemoryHierarchy.L2Controller().NumEvictions,
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"mem.NumReads",
		Value:experiment.MemoryHierarchy.MemoryController().NumReads,
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key:"mem.NumWrites",
		Value:experiment.MemoryHierarchy.MemoryController().NumWrites,
	})

	simutil.WriteJsonFile(experiment.Stats, experiment.CPUConfig.OutputDirectory, simutil.STATS_JSON_FILE_NAME)
}

func (experiment *CPUExperiment) LoadStats() {
	simutil.LoadJsonFile(experiment.CPUConfig.OutputDirectory, simutil.STATS_JSON_FILE_NAME, &experiment.statMap)
}

func (experiment *CPUExperiment) GetStatMap() map[string]interface{} {
	if experiment.statMap == nil {
		experiment.statMap = make(map[string]interface{})

		if experiment.Stats == nil {
			experiment.LoadStats()
		}

		for _, stat := range experiment.Stats {
			experiment.statMap[stat.Key] = stat.Value
		}
	}

	return experiment.statMap
}
