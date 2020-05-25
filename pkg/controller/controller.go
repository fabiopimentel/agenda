package controller

import (
	"github.com/fabiopimentel/Agenda/pkg/model"
	"github.com/gofiber/fiber"
	"strconv"
)

type AgendaController struct {
	contactsMap map[int] model.Contact
}

func (a *AgendaController) Init() {
	a.contactsMap = make(map[int]model.Contact)
}
func (a *AgendaController) GetContacts(c *fiber.Ctx) {
	c.JSON(a.contactsMap)
}

func (a *AgendaController) GetContact(c *fiber.Ctx) {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Status(500).Send("No Contact Found with ID")
		return
	}
	if i>=0 && i<= len(a.contactsMap) {
		c.JSON(a.contactsMap[i])
	}else{
		c.Status(404).Send("Contact not Found")
	}

}

func (a *AgendaController) NewContact(c *fiber.Ctx) {
	contact := new(model.Contact)
	if err := c.BodyParser(contact); err != nil {
		c.Status(503).Send(err)
		return
	}
	if _, found := a.contactsMap[1]; found == false {
		a.contactsMap[1] = *contact
	}else{
		a.contactsMap[len(a.contactsMap)+1] =*contact
	}
	c.Status(201).Send("Contact added")
}

func (a *AgendaController) DeleteContact(c *fiber.Ctx) {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Status(500).Send("No Contact Found with ID")
		return
	}
	delete(a.contactsMap, i)
	c.Send("Contact Successfully deleted")
}

//func RemoveIndex(s []model.Contact, index int) []model.Contact {
//	return append(s[:index], s[index+1:]...)
//}
