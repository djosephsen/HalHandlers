package bothandlers

import (
	"github.com/djosephsen/hal"
	"fmt"
	)

var Listrooms = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `(what room)|(list *room)|(room *list)`,
	Run: func(res *hal.Response) error {
		room := res.Message.Room
		reply := fmt.Sprintf("Current room is: %s",room)
		return res.Send(reply)
	},
}
