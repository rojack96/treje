package set

const (
	AlreadyExists = "already exists in the set"
	EmptySet      = "empty set"
)

type Operations interface {
	Add()
	Remove()
}
