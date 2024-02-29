package http

import (
	useCase "clean_architecture/services/contact/internal/useCase"

	"github.com/julienschmidt/httprouter"
)

type Delivery struct {
	ucContact useCase.Contact
	ucGroup   useCase.Group
	Router    *httprouter.Router
}

func New(ucContact useCase.Contact, ucGroup useCase.Group) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
		Router:    httprouter.New(),
	}

	d.SetupRoutes()

	return d
}
