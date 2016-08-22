package cpu

import (
	"fmt"
	"github.com/mcai/acogo/simutil"
)

func (experiment *CPUExperiment) DumpStats() {
	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "TotalCycles",
		Value: experiment.CycleAccurateEventQueue.CurrentCycle,
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "SimulationTime",
		Value: fmt.Sprintf("%v", experiment.EndTime.Sub(experiment.BeginTime)),
	})

	experiment.Stats = append(experiment.Stats, simutil.Stat{
		Key: "CyclesPerSecond",
		Value: float64(experiment.CycleAccurateEventQueue.CurrentCycle) / experiment.EndTime.Sub(experiment.BeginTime).Seconds(),
	})

	simutil.WriteJsonFile(experiment.Stats, experiment.Config.OutputDirectory, simutil.STATS_JSON_FILE_NAME)
}

func (experiment *CPUExperiment) LoadStats() {
	simutil.LoadJsonFile(experiment.Config.OutputDirectory, simutil.STATS_JSON_FILE_NAME, &experiment.statMap)
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
