package user

type ModuleUser struct{}

func (u *ModuleUser) GetName() string {
	return "User"
}

func (u *ModuleUser) Init() {}
