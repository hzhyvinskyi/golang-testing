package adapter

type Target interface {
	Query() string
}

type Adaptee struct {}

func (a *Adaptee) SpecificQuery() string {
	return "Spec req"
}

type Adapter struct {
	* Adaptee
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func (a *Adapter) Query() string {
	return a.SpecificQuery()
}
