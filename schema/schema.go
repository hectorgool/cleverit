package schema

type (
	Openapi struct {
		Openapi string   `json:"openapi"`
		Servers []string `json:"servers"`
	}
	Info struct {
		Description string  `json:"description"`
		Version     string  `json:"version"`
		Title       string  `json:"title"`
		Contact     Contact `json:"contact"`
	}
	Contact struct {
		Email string `json:"email"`
	}
	BeerItem struct {
		ID       uint    `gorm:"primaryKey"`
		Name     string  `json:"name"`
		Brewery  string  `json:"brewery"`
		Country  string  `json:"country"`
		Price    float64 `json:"price"`
		Currency string  `json:"currency"`
	}
	BeerBox struct {
		ID         uint    `gorm:"primaryKey"`
		PriceTotal float64 `json:"price_total"`
	}
	API struct {
		Openapi
		Info Info
	}
)

const (
	OPENAPIN    string = "3.0.0"
	DESCRIPTION string = "Esta API esta diseñada para ser una prueba para los nuevos candidatos al equipo."
	VERSION     string = "1.0.0"
	TITLE       string = "API Falabella FIF"
	EMAIL       string = "lugaetea@falabella.cl"
)

var (
	Beer = API{
		Openapi{
			Openapi: OPENAPIN,
			Servers: []string{},
		},
		Info{
			Description: DESCRIPTION,
			Version:     VERSION,
			Title:       TITLE,
			Contact: Contact{
				Email: EMAIL,
			},
		},
	}
)
