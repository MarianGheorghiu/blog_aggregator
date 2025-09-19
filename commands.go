package main

import "errors"

type command struct {
	// Numele comenzii (ex: "login")
	Name string
	// Argumentele asociate (ex: ["John"])
	Args []string
}

type commands struct {
	// Mapă care leagă un nume de comandă de un handler.
	// Ar fi util să adăugăm și un "help text" aici, pentru autogenerare de usage.
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	// Nu se validează dacă numele comenzii există deja → pot apărea suprascrieri silențioase.
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	// Verificăm dacă există comanda înregistrată
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		// Poate ar fi util să listăm aici comenzile valide
		return errors.New("command not found")
	}
	// Executăm handler-ul specific
	return f(s, cmd)
}
