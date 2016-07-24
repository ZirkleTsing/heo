package acogo

type RoundRobinArbiter interface {
	GetResource() interface{}
	GetRequesters() []interface{}
	ResourceAvailable(resource interface{}) bool
	RequesterHasRequests(requester interface{}) bool
}

func Next(arbiter RoundRobinArbiter) interface{}  {
	if !arbiter.ResourceAvailable(arbiter.GetResource()) {
		return nil
	}

	for i:=0; i < len(arbiter.GetRequesters()); i++ {
		var requester = arbiter.GetRequesters()[i] //TODO
		if arbiter.RequesterHasRequests(requester) {
			return requester
		}
	}

	return nil
}
