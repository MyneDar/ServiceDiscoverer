package interfaces

type parser interface {
	execute(token []string, commandMap *map[string]interface{})
}
