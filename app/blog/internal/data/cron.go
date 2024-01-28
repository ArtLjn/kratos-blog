package data

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func init() {
	CronDay()
}

// ********** Scheduled tasks ******************//

func CronDay() {
	c := cron.New()
	b := blogRepo{}
	e, err := c.AddFunc("0 0 * * *", func() {
		go b.UpdateBlogVisitsCount()
	})
	fmt.Println(e, err)
	c.Start()
}
