package set

const (
	AlreadyExists = "already exists in the set"
	EmptySet      = "set is empty"
	HasDuplicates = "set has duplicates"
	ElemNotExist  = "element does not exist in the set"
)

type Operations interface {
	Add()
	Remove()
}
