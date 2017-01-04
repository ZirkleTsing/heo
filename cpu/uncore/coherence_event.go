package uncore

type GeneralCacheControllerServiceNonblockingRequestEvent struct {
	CacheController *BaseCacheController
	Access          *MemoryHierarchyAccess
	Tag             uint32
	Set             uint32
	Way             uint32
	HitInCache      bool
}

func NewGeneralCacheControllerServiceNonblockingRequestEvent(cacheController *BaseCacheController, access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32, hitInCache bool) *GeneralCacheControllerServiceNonblockingRequestEvent {
	var event = &GeneralCacheControllerServiceNonblockingRequestEvent{
		CacheController:cacheController,
		Access:access,
		Tag:tag,
		Set:set,
		Way:way,
		HitInCache:hitInCache,
	}

	return event
}

type GeneralCacheControllerNonblockingRequestHitToTransientTagEvent struct {
	CacheController *BaseCacheController
	Access          *MemoryHierarchyAccess
	Tag             uint32
	Set             uint32
	Way             uint32
}

func NewGeneralCacheControllerNonblockingRequestHitToTransientTagEvent(cacheController *BaseCacheController, access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32) *GeneralCacheControllerNonblockingRequestHitToTransientTagEvent {
	var event = &GeneralCacheControllerNonblockingRequestHitToTransientTagEvent{
		CacheController:cacheController,
		Access:access,
		Tag:tag,
		Set:set,
		Way:way,
	}

	return event
}

type GeneralCacheControllerLineReplacementEvent struct {
	CacheController *BaseCacheController
	Access          *MemoryHierarchyAccess
	Tag             uint32
	Set             uint32
	Way             uint32
}

func NewGeneralCacheControllerLineReplacementEvent(cacheController *BaseCacheController, access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32) *GeneralCacheControllerLineReplacementEvent {
	var event = &GeneralCacheControllerLineReplacementEvent{
		CacheController:cacheController,
		Access:access,
		Tag:tag,
		Set:set,
		Way:way,
	}

	return event
}

type GeneralCacheControllerLastPutSOrPutMAndDataFromOwnerEvent struct {
	CacheController *BaseCacheController
	Access          *MemoryHierarchyAccess
	Tag             uint32
	Set             uint32
	Way             uint32
}

func NewGeneralCacheControllerLastPutSOrPutMAndDataFromOwnerEvent(cacheController *BaseCacheController, access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32) *GeneralCacheControllerLastPutSOrPutMAndDataFromOwnerEvent {
	var event = &GeneralCacheControllerLastPutSOrPutMAndDataFromOwnerEvent{
		CacheController:cacheController,
		Access:access,
		Tag:tag,
		Set:set,
		Way:way,
	}

	return event
}

type LastLevelCacheControllerLineInsertEvent struct {
	CacheController *BaseCacheController
	Access          *MemoryHierarchyAccess
	Tag             uint32
	Set             uint32
	Way             uint32
	VictimTag       int32
}

func NewLastLevelCacheControllerLineInsertEvent(cacheController *BaseCacheController, access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32, victimTag int32) *LastLevelCacheControllerLineInsertEvent {
	var event = &LastLevelCacheControllerLineInsertEvent{
		CacheController:cacheController,
		Access:access,
		Tag:tag,
		Set:set,
		Way:way,
		VictimTag:victimTag,
	}

	return event
}

func (event *LastLevelCacheControllerLineInsertEvent) Eviction() bool {
	return event.VictimTag != INVALID_TAG
}