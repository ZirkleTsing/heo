package noc

import (
	"fmt"
	"bytes"
	"encoding/json"
)

const STATS_JSON_FILE_NAME = "stats.json"

type Stat struct {
	Key   string
	Value interface{}
}

type Stats []Stat

func (stats Stats) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")

	for i, stat := range stats {
		if i != 0 {
			buf.WriteString(",")
		}

		key, err := json.Marshal(stat.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)

		buf.WriteString(":")

		val, err := json.Marshal(stat.Value)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteString("}")

	return buf.Bytes(), nil
}

func (experiment *Experiment) DumpStats() {
	experiment.Stats = append(experiment.Stats, Stat{
		Key: "TotalCycles",
		Value: experiment.CycleAccurateEventQueue.CurrentCycle,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "SimulationTime",
		Value: fmt.Sprintf("%v", experiment.EndTime.Sub(experiment.BeginTime)),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "CyclesPerSecond",
		Value: float64(experiment.CycleAccurateEventQueue.CurrentCycle) / experiment.EndTime.Sub(experiment.BeginTime).Seconds(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "PacketsPerSecond",
		Value: float64(experiment.Network.NumPacketsTransmitted) / experiment.EndTime.Sub(experiment.BeginTime).Seconds(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPacketsReceived",
		Value: experiment.Network.NumPacketsReceived,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPacketsTransmitted",
		Value: experiment.Network.NumPacketsTransmitted,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "Throughput",
		Value: experiment.Network.Throughput(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePacketDelay",
		Value: experiment.Network.AveragePacketDelay(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePacketHops",
		Value: experiment.Network.AveragePacketHops(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPacketDelay",
		Value: experiment.Network.MaxPacketDelay,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPacketHops",
		Value: experiment.Network.MaxPacketHops,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPayloadPacketsReceived",
		Value: experiment.Network.NumPayloadPacketsReceived,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPayloadPacketsTransmitted",
		Value: experiment.Network.NumPayloadPacketsTransmitted,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "PayloadThroughput",
		Value: experiment.Network.PayloadThroughput(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePayloadPacketDelay",
		Value: experiment.Network.AveragePayloadPacketDelay(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePayloadPacketHops",
		Value: experiment.Network.AveragePayloadPacketHops(),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPayloadPacketDelay",
		Value: experiment.Network.MaxPayloadPacketDelay,
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPayloadPacketHops",
		Value: experiment.Network.MaxPayloadPacketHops,
	})

	for _, state := range VALID_FLIT_STATES {
		experiment.Stats = append(experiment.Stats, Stat{
			Key: fmt.Sprintf("AverageFlitPerStateDelay[%s]", state),
			Value: experiment.Network.AverageFlitPerStateDelay(state),
		})
	}

	for _, state := range VALID_FLIT_STATES {
		experiment.Stats = append(experiment.Stats, Stat{
			Key: fmt.Sprintf("MaxFlitPerStateDelay[%s]", state),
			Value: experiment.Network.MaxFlitPerStateDelay[state],
		})
	}

	WriteJsonFile(experiment.Stats, experiment.Config.OutputDirectory, STATS_JSON_FILE_NAME)
}

func (experiment *Experiment) LoadStats() {
	LoadJsonFile(experiment.Config.OutputDirectory, STATS_JSON_FILE_NAME, &experiment.statMap)
}

func (experiment *Experiment) GetStatMap() map[string]interface{} {
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
