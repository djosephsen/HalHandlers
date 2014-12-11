package bothandlers

import (
	"fmt"
	"github.com/djosephsen/hal"
	"math/rand"
	"time"
	"strings"
)

var Secondtest = &hal.Chore{
	Name:  `Second Test`,
	Schedule: `* * * * * *`,
	Run: func(res *hal.Response) error {

		return res.Reply(reply)
	},
}
