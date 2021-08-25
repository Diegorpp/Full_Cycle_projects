package application


//Definindo uma interface
type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	Getname() string
	GetStatus() string
	GetPrice() float64
}

// Definido um conjunto de variaveis constantes
const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)
//export GOPATH=/usr/local/go


// Definindo uma struct
type Product struct {
	ID string
	Name string
	Price float64
	Status string
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return error.New( text: "Teh price must be greater than zero to enable the product")
}