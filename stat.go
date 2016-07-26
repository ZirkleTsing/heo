package acogo

import (
	"fmt"
	"bytes"
	"encoding/json"
)

type Stat struct {
	Key   string
	Value string
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
		Value: fmt.Sprintf("%d", experiment.CycleAccurateEventQueue.CurrentCycle),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "SimulationTime",
		Value: fmt.Sprintf("%v", experiment.EndTime.Sub(experiment.BeginTime)),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "CyclesPerSecond",
		Value: fmt.Sprintf("%f", float64(experiment.CycleAccurateEventQueue.CurrentCycle) / experiment.EndTime.Sub(experiment.BeginTime).Seconds()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "PacketsPerSecond",
		Value: fmt.Sprintf("%f", float64(experiment.Network.NumPacketsTransmitted) / experiment.EndTime.Sub(experiment.BeginTime).Seconds()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPacketsReceived",
		Value: fmt.Sprintf("%d", experiment.Network.NumPacketsReceived),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPacketsTransmitted",
		Value: fmt.Sprintf("%d", experiment.Network.NumPacketsTransmitted),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "Throughput",
		Value: fmt.Sprintf("%f", experiment.Network.Throughput()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePacketDelay",
		Value: fmt.Sprintf("%f", experiment.Network.AveragePacketDelay()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePacketHops",
		Value: fmt.Sprintf("%f", experiment.Network.AveragePacketHops()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPacketDelay",
		Value: fmt.Sprintf("%d", experiment.Network.MaxPacketDelay),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPacketHops",
		Value: fmt.Sprintf("%d", experiment.Network.MaxPacketHops),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPayloadPacketsReceived",
		Value: fmt.Sprintf("%d", experiment.Network.NumPayloadPacketsReceived),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "NumPayloadPacketsTransmitted",
		Value: fmt.Sprintf("%d", experiment.Network.NumPayloadPacketsTransmitted),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "PayloadThroughput",
		Value: fmt.Sprintf("%f", experiment.Network.PayloadThroughput()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePayloadPacketDelay",
		Value: fmt.Sprintf("%f", experiment.Network.AveragePayloadPacketDelay()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "AveragePayloadPacketHops",
		Value: fmt.Sprintf("%f", experiment.Network.AveragePayloadPacketHops()),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPayloadPacketDelay",
		Value: fmt.Sprintf("%d", experiment.Network.MaxPayloadPacketDelay),
	})

	experiment.Stats = append(experiment.Stats, Stat{
		Key: "MaxPayloadPacketHops",
		Value: fmt.Sprintf("%d", experiment.Network.MaxPayloadPacketHops),
	})

	for _, state := range VALID_FLIT_STATES {
		experiment.Stats = append(experiment.Stats, Stat{
			Key: fmt.Sprintf("AverageFlitPerStateDelay::%s", state),
			Value: fmt.Sprintf("%f", experiment.Network.AverageFlitPerStateDelay(state)),
		})
	}

	for _, state := range VALID_FLIT_STATES {
		experiment.Stats = append(experiment.Stats, Stat{
			Key: fmt.Sprintf("MaxFlitPerStateDelay::%s", state),
			Value: fmt.Sprintf("%d", experiment.Network.MaxFlitPerStateDelay[state]),
		})
	}

	fmt.Println("Stats:")
	for _, stat := range experiment.Stats {
		fmt.Printf("  %s: %s\n", stat.Key, stat.Value)
	}

	WriteJsonFile(experiment.Stats, experiment.Config.OutputDirectory, "stats.json")
}
