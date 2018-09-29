package impl

type Datastore interface {
	Name() string
	FindUserById(id int64) (string, error)
}
