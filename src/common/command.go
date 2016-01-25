package common

// Command se utiliza para mandar comandos a este módulo
type Command struct {
	Cmd  string
	Args map[string]string
}
