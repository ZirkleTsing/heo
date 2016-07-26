package acogo

import "fmt"

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
						return fmt.Sprintf("ACO/aj=%s/a=%s/rf=%s", experiment.Config.AntPacketInjectionRate, experiment.Config.AcoSelectionAlpha, experiment.Config.ReinforcementFactor)
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
				return experiment.GetStatMap()["NumPacketTransmitted"]
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

	for _, state := range VALID_FLIT_STATES {
		csvFields = append(csvFields, CSVField{
			Name: fmt.Sprintf("Average_Flit_per_State_Delay::%s", state),
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()[fmt.Sprintf("AverageFlitPerStateDelay[%s]", state)]
			},
		})

		csvFields = append(csvFields, CSVField{
			Name: fmt.Sprintf("Max_Flit_per_State_Delay::%s", state),
			Callback: func(experiment *Experiment) interface{} {
				return experiment.GetStatMap()[fmt.Sprintf("MaxFlitPerStateDelay[%s]", state)]
			},
		})
	}

	return csvFields
}
