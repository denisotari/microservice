package handler

type Repository interface {
	VersionerRepository
}

type Publisher interface {
	GreeterPublisher
}

type Handler struct {
	DB  Repository
	Pub Publisher
}
