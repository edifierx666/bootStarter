package bootStarter

var Hubs = NewHub()
var RegisterStarter = Hubs.Register
var GetAllStarter = Hubs.AllStarters

type Hub struct {
	starters []IStarter
}

func NewHub() *Hub {
	return &Hub{
		starters: []IStarter{},
	}
}

func (i *Hub) Register(s IStarter) {
	i.starters = append(i.starters, s)
}

func (i *Hub) AllStarters() []IStarter {
	return i.starters
}
