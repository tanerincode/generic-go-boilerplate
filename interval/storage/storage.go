package storage

type Storage interface {
	CleanUp() error
	Disconnect() error
}
