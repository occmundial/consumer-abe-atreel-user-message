package interfaces

type IQuery interface {
	GetDicState() (map[string]string, error)
	CheckHealth() error
}
