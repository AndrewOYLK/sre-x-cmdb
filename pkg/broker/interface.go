package broker

type IBroker interface {
	produce()
	consume()
}
