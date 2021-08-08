package distributed

type Service struct {
	Name string
	Host string
	Port string
	HttpHandleFunctions func()
}
