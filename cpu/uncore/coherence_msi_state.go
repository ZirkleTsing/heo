package uncore

type CacheControllerState string

const (
	CacheControllerState_I = CacheControllerState("I")
	CacheControllerState_IS_D = CacheControllerState("IS_D")
	CacheControllerState_IM_AD = CacheControllerState("IM_AD")
	CacheControllerState_IM_A = CacheControllerState("IM_A")
	CacheControllerState_S = CacheControllerState("S")
	CacheControllerState_SM_AD = CacheControllerState("SM_AD")
	CacheControllerState_SM_A = CacheControllerState("SM_A")
	CacheControllerState_M = CacheControllerState("M")
	CacheControllerState_MI_A = CacheControllerState("MI_A")
	CacheControllerState_SI_A = CacheControllerState("SI_A")
	CacheControllerState_II_A = CacheControllerState("II_A")
)

func (cacheControllerState CacheControllerState) Stable() bool {
	return cacheControllerState == CacheControllerState_I ||
		cacheControllerState == CacheControllerState_S ||
		cacheControllerState == CacheControllerState_M
}

func (cacheControllerState CacheControllerState) Transient() bool {
	return !cacheControllerState.Stable()
}

type DirectoryControllerState string

const (
	DirectoryControllerState_I = DirectoryControllerState("I")
	DirectoryControllerState_IS_D = DirectoryControllerState("IS_D")
	DirectoryControllerState_IM_D = DirectoryControllerState("IM_D")
	DirectoryControllerState_S = DirectoryControllerState("S")
	DirectoryControllerState_M = DirectoryControllerState("M")
	DirectoryControllerState_S_D = DirectoryControllerState("S_D")
	DirectoryControllerState_MI_A = DirectoryControllerState("MI_A")
	DirectoryControllerState_SI_A = DirectoryControllerState("SI_A")
)

func (directoryControllerState DirectoryControllerState) Stable() bool {
	return directoryControllerState == DirectoryControllerState_I ||
		directoryControllerState == DirectoryControllerState_S ||
		directoryControllerState == DirectoryControllerState_M
}

func (directoryControllerState DirectoryControllerState) Transient() bool {
	return !directoryControllerState.Stable()
}