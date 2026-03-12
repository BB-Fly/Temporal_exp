package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	greeting "temporal-exp/src/greating/workflow"
	mock "temporal-exp/src/mock"
	"temporal-exp/src/models"
	prelock "temporal-exp/src/prelock/workflow"
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
	} else if len(os.Args) > 1 && os.Args[1] == "batch" {
		// 批量运行班次推荐工作流
		runBatchScheduleWorkflows(c)
	} else if len(os.Args) > 1 && os.Args[1] == "prelock" {
		// 运行预锁库存工作流
		runPreLockWorkflow(c)
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

func runBatchScheduleWorkflows(c client.Client) {
	// 加载多个Mock用户请求
	requests := mock.LoadMockUserRequests()

	// 创建等待组
	var wg sync.WaitGroup

	// 执行所有请求
	for i, request := range requests {
		// 添加到等待组
		wg.Add(1)

		// 启动goroutine执行工作流
		go func(idx int, req models.UserRequest) {
			defer wg.Done()

			// 生成唯一的工作流ID
			workflowID := fmt.Sprintf("schedule-recommendation-workflow-%d", idx)

			options := client.StartWorkflowOptions{
				ID:        workflowID,
				TaskQueue: "my-task-queue",
			}

			we, err := c.ExecuteWorkflow(context.Background(), options, schedule.RecommendSchedulesWorkflow, req)
			if err != nil {
				log.Printf("Unable to execute workflow %d: %v", idx, err)
				return
			}
			log.Printf("Started workflow %d, WorkflowID: %s, RunID: %s", idx, we.GetID(), we.GetRunID())

			var result interface{}
			err = we.Get(context.Background(), &result)
			if err != nil {
				log.Printf("Unable get workflow %d result: %v", idx, err)
				return
			}
			log.Printf("Workflow %d result: %v", idx, result)
		}(i, request)

		// 每间隔短暂时间执行一个请求，模拟真实的线上情况
		// 为了模拟并行情况，前两个请求间隔时间较短
		if i == 0 {
			time.Sleep(100 * time.Millisecond)
		} else if i == 1 {
			time.Sleep(50 * time.Millisecond)
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}

	// 等待所有工作流执行完成
	wg.Wait()
	log.Println("All workflow executions completed")
}

func runPreLockWorkflow(c client.Client) {
	options := client.StartWorkflowOptions{
		ID:        "prelock-seats-workflow",
		TaskQueue: "my-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, prelock.PreLockSeatsWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result error
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow completed successfully")
}
