package main

import (
	"context"
	"log"
	"os"

	greeting "temporal-exp/src/greating/workflow"
	"temporal-exp/src/mock"
	schedule "temporal-exp/src/schedule/workflow"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// 检查命令行参数，决定运行哪个工作流
	if len(os.Args) > 1 && os.Args[1] == "schedule" {
		// 运行班次推荐工作流
		runScheduleWorkflow(c)
	} else {
		// 运行问候工作流
		runGreetingWorkflow(c)
	}
}

func runGreetingWorkflow(c client.Client) {
	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "my-task-queue",
	}

	name := "World"
	if len(os.Args) > 2 {
		name = os.Args[2]
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, greeting.SayHelloWorkflow, name)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}

func runScheduleWorkflow(c client.Client) {
	options := client.StartWorkflowOptions{
		ID:        "schedule-recommendation-workflow",
		TaskQueue: "my-task-queue",
	}

	// 加载Mock用户请求
	request := mock.LoadMockUserRequest()

	we, err := c.ExecuteWorkflow(context.Background(), options, schedule.RecommendSchedulesWorkflow, request)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result interface{}
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
