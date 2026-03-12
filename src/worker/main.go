package main

import (
	"log"
	greeting_a "temporal-exp/src/greating/activity"
	greeting_w "temporal-exp/src/greating/workflow"
	schedule_a "temporal-exp/src/schedule/activity"
	schedule_w "temporal-exp/src/schedule/workflow"
	prelock_a "temporal-exp/src/prelock/activity"
	prelock_w "temporal-exp/src/prelock/workflow"

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

	// 注册预锁库存相关的工作流和活动
	w.RegisterWorkflow(prelock_w.PreLockSeatsWorkflow)
	w.RegisterActivity(prelock_a.AddOrderToCarpoolPool)
	w.RegisterActivity(prelock_a.AddOrderToPreLockPool)
	w.RegisterActivity(prelock_a.InitCtx)
	w.RegisterActivity(prelock_a.RecallOrderFromCarpoolPool)
	w.RegisterActivity(prelock_a.RecallOrderFromPreLockPool)
	w.RegisterActivity(prelock_a.LockShift)
	w.RegisterActivity(prelock_a.GetForasInfo)
	w.RegisterActivity(prelock_a.SetShiftVersion)
	w.RegisterActivity(prelock_a.GetShiftInventoryFromRedis)
	w.RegisterActivity(prelock_a.GetShiftInventoryFromDB)
	w.RegisterActivity(prelock_a.GetPrelockInventory)
	w.RegisterActivity(prelock_a.AddShiftOrderToStgData)
	w.RegisterActivity(prelock_a.GetRtFeature)
	w.RegisterActivity(prelock_a.TryOccupySeats)
	w.RegisterActivity(prelock_a.ApiCheck)
	w.RegisterActivity(prelock_a.DelOrderFromPreLockPool)
	w.RegisterActivity(prelock_a.DelOrderFromCarpoolPool)
	w.RegisterActivity(prelock_a.LockSeats)
	w.RegisterActivity(prelock_a.UnlockShift)
	w.RegisterActivity(prelock_a.AlreadyLocked)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
