package noc

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type TraceFileLine struct {
	ThreadId int64
	Pc       int64
	Read     bool
	Ea       int64
}

type TraceTrafficGenerator struct {
	Network              *Network
	PacketInjectionRate  float64
	MaxPackets           int64
	TraceFileNames       []string
	TraceFileLines       [][]*TraceFileLine
	CurrentTraceFileLine []int64
}

func NewTraceTrafficGenerator(network *Network, packetInjectionRate float64, maxPackets int64, traceFileNames []string) *TraceTrafficGenerator {
	var generator = &TraceTrafficGenerator{
		Network:network,
		PacketInjectionRate:packetInjectionRate,
		MaxPackets:maxPackets,
		TraceFileNames:traceFileNames,
	}

	for threadId, traceFileName := range traceFileNames {
		generator.TraceFileLines = append(generator.TraceFileLines, []*TraceFileLine{})

		traceFile, err := os.Open(traceFileName)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(traceFile)
		for scanner.Scan() {
			fmt.Println(scanner.Text())

			var line = scanner.Text()
			var parts = strings.Split(line, ",")

			pc, err := strconv.ParseInt(parts[1], 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			var read = parts[2] == "R"

			ea, err := strconv.ParseInt(parts[3], 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			var traceFileLine = &TraceFileLine{
				ThreadId:int64(threadId),
				Pc:pc,
				Read:read,
				Ea:ea,
			}

			generator.TraceFileLines[threadId] = append(generator.TraceFileLines[threadId], traceFileLine)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		traceFile.Close()
	}

	return generator
}

func (generator *TraceTrafficGenerator) AdvanceOneCycle() {
	//TODO：每隔10个时钟周期（cycle），从每个trace文件中取一行，并创建和发送对应的消息包（packet
	if(generator.Network.Driver.CycleAccurateEventQueue().CurrentCycle % 10 == 0) {
		var packet = NewDataPacket(generator.Network, src, dest, 16, true, func() {})

		generator.Network.Driver.CycleAccurateEventQueue().Schedule(func() {
			generator.Network.Receive(packet)
		}, 1)
	}
}