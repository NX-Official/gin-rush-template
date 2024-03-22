package ping

type ModulePing struct{}

func (p *ModulePing) GetName() string {
	return "Ping"
}

func (p *ModulePing) Init() {}
