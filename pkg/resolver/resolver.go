package resolver

// A Generic Type Reolver that retiroves enfpoints and exposes channels for refreshing them

type Resolver interface {
	GetEndPoints() []string
	Watch()
	Refresh() <-chan bool
	Close()
	Latest() []string
}
