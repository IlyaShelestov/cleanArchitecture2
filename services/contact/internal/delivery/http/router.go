package http

func (d *Delivery) SetupRoutes() {
	d.Router.GET("/contacts/:id", d.GetContact)
	d.Router.POST("/contacts", d.CreateContact)
	d.Router.PUT("/contacts/:id", d.UpdateContact)
	d.Router.DELETE("/contacts/:id", d.DeleteContact)

	d.Router.GET("/groups/:id", d.GetGroup)
	d.Router.POST("/groups", d.CreateGroup)
	d.Router.PUT("/groups/:id", d.UpdateGroup)
	d.Router.DELETE("/groups/:id", d.DeleteGroup)

	d.Router.POST("/groupContacts/:groupId/contacts", d.CreateContactInGroup)
	d.Router.POST("/groupContacts/:groupId/addContact/:contactId", d.AddContactToGroup)
	d.Router.DELETE("/groupContacts/:groupId/removeContact/:contactId", d.DeleteContactFromGroup)
}
