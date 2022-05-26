package adapter

type Validator interface {
	Validate(obj interface{}) error
}
