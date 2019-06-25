package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fishjerky/jobrunner-example/jobs"

	"gopkg.in/robfig/cron.v2"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

// Example of GIN micro framework
func main() {
	//Part 1. Job Setting
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)

	//1.every 5s, mintues
	jobrunner.Every(5*time.Second, jobs.ReminderEmails{})
	//=jobrunner.Schedule("@every 5s", tasks.ReminderEmails{})

	//2.every mintue, hourly, daily,
	jobrunner.Every(1*time.Minute, jobs.Job1{"Every mintue"})
	//=jobrunner.Schedule("1 * * * * *", tasks.Job1{"Every mintue"})
	jobrunner.Schedule("@hourly", jobs.Job1{"Every hour"}) //run at xx:00
	jobrunner.Schedule("@daily", jobs.Job1{"Every day"})   //run at 00:00

	//3.fixed time
	jobrunner.Schedule("TZ=Asia/Taipei 21 20 16 * * *", jobs.Job1{"Runs at 04:30 Taipei time every day"})

	//4. others
	jobrunner.In(10*time.Second, jobs.Job1{"Welcome in 10 seconds"}) // one time job. starts after 10sec
	jobrunner.Now(jobs.Job1{"do the job as soon as it's triggered"}) // do the job as soon as it's triggered

	//Part 2. Routing
	routes := gin.Default()

	routes.GET("/", JobHtml)

	// Resource to return the JSON data
	routes.GET("/jobrunner/json", JobJson)

	// Load template file location relative to the current working directory
	routes.LoadHTMLGlob("views/*.html") //Copy from jobrunner/views

	//Run job once
	routes.GET("/jobrunner/run/:job-id", func(c *gin.Context) {
		jobId, err := strconv.Atoi(c.Param("job-id"))
		if err != nil {
			return
		}

		id := cron.EntryID(jobId)

		jobrunner.Schedule("@every 7s", jobs.ReminderEmails{})

		c.String(http.StatusOK, "add  %s", id)
	})

	//Stop job
	routes.GET("/jobrunner/stop/:job-id", func(c *gin.Context) {
		jobId, err := strconv.Atoi(c.Param("job-id"))
		if err != nil {
			return
		}

		id := cron.EntryID(jobId)

		jobrunner.Remove(id)
		c.String(http.StatusOK, "REMOVE %s", id)
	})

	routes.Run(":8080")

}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")

}

func JobJson(c *gin.Context) {
	// returns a map[string]interface{} that can be marshalled as JSON
	//c.String(http.StatusOK, "hello, json")
	c.JSON(200, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {
	// Returns the template data pre-parsed
	//s := jobrunner.StatusPage()
	//fmt.Printf(s)
	//c.HTML(200, "", jobrunner.StatusPage())
	c.HTML(200, "Status.html", jobrunner.StatusPage())
}
