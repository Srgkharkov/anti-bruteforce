package cli

type appCLI struct {
	addr, command, login, pass, ip, subnet *string
	port                                   *int
}

func NewAppCLI(command *string, addr *string, port *int, login *string, pass *string, ip *string, subnet *string) *appCLI {

	return &appCLI{command: command, addr: addr, port: port, login: login, pass: pass, ip: ip, subnet: subnet}
}

func (a *appCLI) Run() {
	switch *a.command {
	case "request":
		//todo

	}
}
