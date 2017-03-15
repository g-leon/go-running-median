package dataprovider

// IntFetcher is an interface to a data provider
// that returns one integer at each Fetch call.
type IntFetcher interface {
	Fetch() (int, error)
}