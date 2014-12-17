package bothandlers

import (
	"github.com/djosephsen/hal"
	"fmt"
	"reflect"
	)

var Help = &hal.Handler{
	Method:  hal.HEAR,
	Pattern: `help (.*)`,
	Usage: `help: prints this message when you type "botname help"`,
	Run: func(res *hal.Response) error {
		var reply string
		handlers:=hal.Handlers
		for _,h:=range handlers{
			hval:=reflect.ValueOf(h)
			usage:=hval.FieldByName(`usage`)
			reply=fmt.Sprintf("%s\n%s",reply,usage)
			}
		return res.Send(reply)
	},
}
