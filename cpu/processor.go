package cpu

type Processor struct {
	Experiment              *CPUExperiment
	Cores                   []*Core
	ContextToThreadMappings map[*Context]*Thread
}

func NewProcessor(experiment *CPUExperiment) *Processor {
	var processor = &Processor{
		Experiment:experiment,
		ContextToThreadMappings:make(map[*Context]*Thread),
	}

	for i := uint32(0); i < experiment.Config.NumCores; i++ {
		var core = NewCore(processor, i)

		for j := uint32(0); j < experiment.Config.NumThreadsPerCore; j++ {
			var thread = NewThread(core, j)
			core.Threads = append(core.Threads, thread)
		}
	}

	processor.UpdateContextToThreadAssignments()

	return processor
}

func (processor *Processor) UpdateContextToThreadAssignments() {
	var ch = make(chan *Context, len(processor.Experiment.Kernel.Contexts))

	for _, context := range processor.Experiment.Kernel.Contexts {
		ch <- context
	}

	close(ch)

	for _, context := range processor.Experiment.Kernel.Contexts {
		if context.ThreadId != -1 && processor.ContextToThreadMappings[context] == nil {
			context.State = ContextState_RUNNING

			var coreNum = uint32(context.ThreadId) / processor.Experiment.Config.NumThreadsPerCore
			var threadNum = uint32(context.ThreadId) % processor.Experiment.Config.NumThreadsPerCore

			var candidateThread = processor.Cores[coreNum].Threads[threadNum]

			processor.ContextToThreadMappings[context] = candidateThread

			candidateThread.Context = context
			//candidateThread.UpdateFetchNpcAndNnpcFromRegs()
		} else if context.State == ContextState_FINISHED {
			//var thread = processor.ContextToThreadMappings[context]
			//if thread.IsLastDecodedDynamicInstCommitted {
			//	processor.Kill(context)
			//}
			//TODO
		}
	}
}