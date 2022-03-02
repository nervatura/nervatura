package service

// APIService Nervatura API interface
type APIService interface {
	StartService() error
	Results() string
	ConnectApp(interface{})
	StopService(interface{}) error
}

type ContextKey struct {
	name string
}

func (k *ContextKey) String() string {
	return k.name
}

var NstoreCtxKey = &ContextKey{"nstore"}
var TokenCtxKey = &ContextKey{"token"}
