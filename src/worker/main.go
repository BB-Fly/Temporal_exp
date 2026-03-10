package main

import (
	"log"
	greeting_a "temporal-exp/src/greating/activity"
	greeting_w "temporal-exp/src/greating/workflow"
	schedule_a "temporal-exp/src/schedule/activity"
	schedule_w "temporal-exp/src/schedule/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "my-task-queue", worker.Options{})

	// 注册问候相关的工作流和活动
	w.RegisterWorkflow(greeting_w.SayHelloWorkflow)
	w.RegisterActivity(greeting_a.Greet)

	// 注册班次推荐相关的工作流和活动
	w.RegisterWorkflow(schedule_w.RecommendSchedulesWorkflow)
	w.RegisterActivity(schedule_a.RecallSchedules)
	w.RegisterActivity(schedule_a.FilterBySeatCount)
	w.RegisterActivity(schedule_a.FilterByDepartureTime)
	w.RegisterActivity(schedule_a.FilterByDistance)
	w.RegisterActivity(schedule_a.SortByDistance)
	w.RegisterActivity(schedule_a.PaginateSchedules)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
