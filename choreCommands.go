package bothandlers

import (
	"fmt"
	"github.com/djosephsen/hal"
	"math/rand"
	"time"
	"strings"
)

var ListChores = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `(list chores)|(chore list)`,
	Run: func(res *hal.Response) error {
		var reply string
		for _,c := range res.Robot.chores{
			reply = fmt.Sprintf("%s\n%s:%s:%s:%s",reply,c.Name, c.Schedule, c.Next, c.State)
		}
		return res.Reply(reply)
	},
}
