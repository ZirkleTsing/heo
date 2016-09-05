package cpu

type Processor struct {
	Experiment              *CPUExperiment
	Cores                   []Core
	ContextToThreadMappings map[*Context]Thread
	Mnemonics               []*Mnemonic
}

func NewProcessor(experiment *CPUExperiment) *Processor {
	var processor = &Processor{
		Experiment:experiment,
		ContextToThreadMappings:make(map[*Context]Thread),
	}

	processor.addMnemonics()

	for i := int32(0); i < experiment.CPUConfig.NumCores; i++ {
		var core = NewOoOCore(processor, i)

		for j := int32(0); j < experiment.CPUConfig.NumThreadsPerCore; j++ {
			var thread = NewMemoryHierarchyThread(core, j)
			core.AddThread(thread)
		}

		processor.Cores = append(processor.Cores, core)
	}

	return processor
}

func (processor *Processor) UpdateContextToThreadAssignments() {
	var contextsToReserve []*Context

	for _, context := range processor.Experiment.Kernel.Contexts {
		if context.ThreadId != -1 && processor.ContextToThreadMappings[context] == nil {
			context.State = ContextState_RUNNING

			var coreNum = context.ThreadId / processor.Experiment.CPUConfig.NumThreadsPerCore
			var threadNum = context.ThreadId % processor.Experiment.CPUConfig.NumThreadsPerCore

			var candidateThread = processor.Cores[coreNum].Threads()[threadNum]

			processor.ContextToThreadMappings[context] = candidateThread

			candidateThread.SetContext(context)

			contextsToReserve = append(contextsToReserve, context)
		} else if context.State == ContextState_FINISHED {
			processor.kill(context)
		} else {
			contextsToReserve = append(contextsToReserve, context)
		}
	}

	processor.Experiment.Kernel.Contexts = contextsToReserve
}

func (processor *Processor) kill(context *Context) {
	if context.State != ContextState_FINISHED {
		panic("Impossible")
	}

	for _, c := range processor.Experiment.Kernel.Contexts {
		if c.Parent == context {
			processor.kill(c)
		}
	}

	if context.Parent == nil {
		context.Process.CloseProgram()
	}

	processor.ContextToThreadMappings[context].SetContext(nil)

	context.ThreadId = -1

	processor.Experiment.BlockingEventDispatcher().Dispatch(NewContextKilledEvent(context))
}