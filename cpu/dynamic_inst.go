package cpu

type DynamicInst struct {
	Thread           Thread
	Pc               uint32
	StaticInst       *StaticInst
	EffectiveAddress uint32
}

//TODO...
//func NewDynamicInst(thread Thread, pc uint32, staticInst *StaticInst) *DynamicInst {
//	var dynamicInst = &DynamicInst{
//		Thread:thread,
//		Pc:pc,
//		StaticInst:staticInst,
//	}
//
//	if staticInst.Mnemonic.
//
//	return dynamicInst
//}


