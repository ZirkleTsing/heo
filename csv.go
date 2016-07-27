package acogo

import (
	"fmt"
	"encoding/csv"
	"os"
)

type CSVField struct {
	Name     string
	Callback func(experiment *Experiment) interface{}
}

func GetCSVFields() []CSVField {
	var csvFields = []CSVField{
		{
			Name: "Data_Packet_Traffic",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.DataPacketTraffic
			},
		},
		{
			Name: "Data_Packet_Injection_Rate_(packets/cycle/node)",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.DataPacketInjectionRate
			},
		},
		{
			Name: "Routing_Algorithm",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.Routing
			},
		},
		{
			Name: "Selection_Policy",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.Selection
			},
		},
		{
			Name: "Ant_Packet_Traffic",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.AntPacketTraffic
			},
		},
		{
			Name: "Ant_Packet_Injection_Rate_(packets/cycle/node)",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.AntPacketInjectionRate
			},
		},
		{
			Name: "Alpha",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.AcoSelectionAlpha
			},
		},
		{
			Name: "Reinforcement_Factor",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.Config.ReinforcementFactor
			},
		},
		{
			Name: "NoC_Routing_Solution",
			Callback: func(experiment *Experiment) interface{} {
				switch experiment.Config.Routing {
				case ROUTING_XY:
					return "XY"
				case ROUTING_ODD_EVEN:
					switch experiment.Config.Selection {
					case SELECTION_BUFFER_LEVEL:
						return "BufferLevel"
					case SELECTION_ACO:
						return fmt.Sprintf("ACO/aj=%f/a=%f/rf=%f", experiment.Config.AntPacketInjectionRate, experiment.Config.AcoSelectionAlpha, experiment.Config.ReinforcementFactor)
					default:
						panic("Impossible")
					}
				default:
					panic("Impossible")
				}
			},
		},
		{
			Name: "Simulation_Time",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["SimulationTime"]
			},
		},
		{
			Name: "Total_Cycles",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["TotalCycles"]
			},
		},
		{
			Name: "Num_Packets_Transmitted",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["NumPacketsTransmitted"]
			},
		},
		{
			Name: "Throughput_(packets/cycle/node)",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["Throughput"]
			},
		},
		{
			Name: "Avg._Packet_Delay_(cycles)",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["AveragePacketDelay"]
			},
		},
		{
			Name: "Avg._Packet_Hops",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["AveragePacketHops"]
			},
		},
		{
			Name: "Num_Payload_Packets_Transmitted",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["NumPayloadPacketsTransmitted"]
			},
		},
		{
			Name: "Payload_Throughput_(packets/cycle/node)",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["PayloadThroughput"]
			},
		},
		{
			Name: "Avg._Payload_Packet_Delay_(cycles)",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["AveragePayloadPacketDelay"]
			},
		},
		{
			Name: "Avg._Payload_Packet_Hops",
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()["AveragePayloadPacketHops"]
			},
		},
	}

	for _, s := range VALID_FLIT_STATES {
		var state = s

		csvFields = append(csvFields, CSVField{
			Name: fmt.Sprintf("Average_Flit_per_State_Delay[%s]", state),
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()[fmt.Sprintf("AverageFlitPerStateDelay[%s]", state)]
			},
		})

		csvFields = append(csvFields, CSVField{
			Name: fmt.Sprintf("Max_Flit_per_State_Delay[%s]", state),
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()[fmt.Sprintf("MaxFlitPerStateDelay[%s]", state)]
			},
		})
	}

	return csvFields
}

func WriteCSVFile(outputDirectory string, outputCSVFileName string, experiments []*Experiment, fields []CSVField) {
	if err := os.MkdirAll(outputDirectory, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Cannot create output directory (%s)", err))
	}

	fp, err := os.Create(outputDirectory + "/" + outputCSVFileName)

	if err != nil {
		panic(fmt.Sprintf("Cannot create CSV file (%s)", err))
	}

	defer fp.Close()

	var w = csv.NewWriter(fp)

	var head []string

	for _, field := range fields {
		head = append(head, field.Name)
	}

	if err := w.Write(head); err != nil {
		panic(fmt.Sprintf("Error writing record to CSV file (%s)", err))
	}

	for _, experiment := range experiments {
		var record []string

		for _, field := range fields {
			record = append(record, fmt.Sprintf("%+v", field.Callback(experiment)))
		}

		if err := w.Write(record); err != nil {
			panic(fmt.Sprintf("Error writing record to CSV file (%s)", err))
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
	}
}
