package data

import (
	"github.com/robfig/cron/v3"
)

func init() {
	CronDay()
}

// ********** Scheduled tasks ******************//

func CronDay() {
	c := cron.New()
	b := blogRepo{}
	c.AddFunc("0 0 * * *", func() {
		go b.UpdateBlogVisitsCount()
	})
	c.Start()
}
