package resolver

// A Generic Type Resolver that retrieves enfdpoints and exposes channels for refreshing them
type Resolver interface {
	GetEndPoints() []string
	Watch()
	Refresh() <-chan bool
	Close()
	Latest() []string
}
