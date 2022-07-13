package application

 import "errors"

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
	return errors.New(text: "Teh price must be greater than zero to enable the product")
}

func (p *Product) IsValid() (bool, error){

}

func (p *Product) Disable() error {

}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) Getname() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
