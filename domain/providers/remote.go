package providers

type RemoteProvider interface {
	Get(url string)
	Post(url string)
}
