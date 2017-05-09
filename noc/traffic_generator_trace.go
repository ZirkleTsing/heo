package noc

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
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

	for _, traceFileName := range traceFileNames {
		traceFile, err := os.Open(traceFileName)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(traceFile)
		for scanner.Scan() {
			fmt.Println(scanner.Text())

			var line = scanner.Text()
			var parts = strings.Split(line, ",")

			for _, part := range parts {
				fmt.Println(part)
			}

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		traceFile.Close()
	}

	return generator
}

func (generator *TraceTrafficGenerator) AdvanceOneCycle() {
}