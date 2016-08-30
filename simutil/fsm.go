package simutil

type FiniteStateMachineEvent struct {
	Fsm       *BaseFiniteStateMachine
	Condition interface{}
	Params    interface{}
}

func NewBaseFiniteStateMachineEvent(fsm *BaseFiniteStateMachine, condition interface{}, params interface{}) *FiniteStateMachineEvent {
	var baseFiniteStateMachineEvent = &FiniteStateMachineEvent{
		Fsm:fsm,
		Condition:condition,
		Params:params,
	}

	return baseFiniteStateMachineEvent
}

type EnterStateEvent struct {
	*FiniteStateMachineEvent
}

func NewEnterStateEvent(fsm *BaseFiniteStateMachine, condition interface{}, params interface{}) *EnterStateEvent {
	var enterStateEvent = &EnterStateEvent{
		FiniteStateMachineEvent:NewBaseFiniteStateMachineEvent(fsm, condition, params),
	}

	return enterStateEvent
}

type ExitStateEvent struct {
	*FiniteStateMachineEvent
}

func NewExitStateEvent(fsm *BaseFiniteStateMachine, condition interface{}, params interface{}) *ExitStateEvent {
	var exitStateEvent = &ExitStateEvent{
		FiniteStateMachineEvent:NewBaseFiniteStateMachineEvent(fsm, condition, params),
	}

	return exitStateEvent
}

type FiniteStateMachine interface {
	State() interface{}
	SetState(condition interface{}, params interface{}, state interface{})
}

type BaseFiniteStateMachine struct {
	state                   interface{}
	BlockingEventDispatcher *BlockingEventDispatcher
	settingStates           bool
}

func NewBaseFiniteStateMachine(state interface{}) *BaseFiniteStateMachine {
	var finiteStateMachine = &BaseFiniteStateMachine{
		state:state,
		BlockingEventDispatcher:NewBlockingEventDispatcher(),
	}

	return finiteStateMachine
}

func (finiteStateMachine *BaseFiniteStateMachine) State() interface{} {
	return finiteStateMachine.state
}

func (finiteStateMachine *BaseFiniteStateMachine) SetState(condition interface{}, params interface{}, state interface{}) {
	if finiteStateMachine.settingStates {
		panic("Impossible")
	}

	finiteStateMachine.settingStates = true

	finiteStateMachine.BlockingEventDispatcher.Dispatch(NewExitStateEvent(finiteStateMachine, condition, params))

	finiteStateMachine.state = state

	finiteStateMachine.BlockingEventDispatcher.Dispatch(NewEnterStateEvent(finiteStateMachine, condition, params))

	finiteStateMachine.settingStates = false
}

type StateTransition struct {
	State               interface{}
	Condition           interface{}
	NewState            interface{}
	Action              func(fsm FiniteStateMachine, condition interface{}, params interface{})
	OnCompletedCallback func(fsm FiniteStateMachine)
}

func NewStateTransition(state interface{}, condition interface{}, newState interface{}, action func(fsm FiniteStateMachine, condition interface{}, params interface{}), onCompletedCallback func(fsm FiniteStateMachine)) *StateTransition {
	var stateTransition = &StateTransition{
		State:state,
		Condition:condition,
		NewState:newState,
		Action:action,
		OnCompletedCallback:onCompletedCallback,
	}

	return stateTransition
}

type StateTransitions struct {
	fsmFactory          *FiniteStateMachineFactory
	state               interface{}
	perStateTransitions map[interface{}]*StateTransition
	onCompletedCallback func(fsm FiniteStateMachine)
}

func NewStateTransitions(fsmFactory *FiniteStateMachineFactory, state interface{}) *StateTransitions {
	var stateTransitions = &StateTransitions{
		fsmFactory:fsmFactory,
		state:state,
		perStateTransitions:make(map[interface{}]*StateTransition),
	}

	return stateTransitions
}

func (stateTransitions *StateTransitions) OnCondition(condition interface{}, action func(fsm FiniteStateMachine, condition interface{}, params interface{}), newState interface{}, onCompletedCallback func(fsm FiniteStateMachine)) {
	stateTransitions.perStateTransitions[condition] = NewStateTransition(stateTransitions.state, condition, newState, action, onCompletedCallback)
}

func (stateTransitions *StateTransitions) fireTransition(fsm FiniteStateMachine, condition interface{}, params interface{}) {
	var stateTransition = stateTransitions.perStateTransitions[condition]

	stateTransition.Action(fsm, condition, params)

	var newState = stateTransition.NewState

	fsm.SetState(condition, params, newState)

	if stateTransition.OnCompletedCallback != nil {
		stateTransition.OnCompletedCallback(fsm)
	}

	if stateTransitions.onCompletedCallback != nil {
		stateTransitions.onCompletedCallback(fsm)
	}
}

type FiniteStateMachineFactory struct {
	transitions map[interface{}]*StateTransitions
}

func NewFiniteStateMachineFactory() *FiniteStateMachineFactory {
	var fsmFactory = &FiniteStateMachineFactory{
		transitions:make(map[interface{}]*StateTransitions),
	}

	return fsmFactory
}

func (fsmFactory *FiniteStateMachineFactory) InState(state interface{}) *StateTransitions {
	if _, ok := fsmFactory.transitions[state]; !ok {
		fsmFactory.transitions[state] = NewStateTransitions(fsmFactory, state)
	}

	return fsmFactory.transitions[state]
}

func (fsmFactory *FiniteStateMachineFactory) FireTransition(fsm FiniteStateMachine, condition interface{}, params interface{}) {
	fsmFactory.transitions[fsm.State()].fireTransition(fsm, condition, params)
}