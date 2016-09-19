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

	for i, core := range experiment.Processor.Cores {
		for j, thread := range core.Threads() {
			experiment.Stats = append(experiment.Stats, simutil.Stat{
				Key:fmt.Sprintf("thread_%d-%d.NumDynamicInsts", i, j),
				Value:thread.NumDynamicInsts(),
			})

			experiment.Stats = append(experiment.Stats, simutil.Stat{
				Key:fmt.Sprintf("thread_%d-%d.InstructionsPerCycle", i, j),
				Value:thread.InstructionsPerCycle(),
			})

			experiment.Stats = append(experiment.Stats, simutil.Stat{
				Key:fmt.Sprintf("thread_%d-%d.CyclesPerInstructions", i, j),
				Value:thread.CyclesPerInstructions(),
			})
		}
	}

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
