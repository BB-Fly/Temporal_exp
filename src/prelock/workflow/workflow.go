package workflow

import (
	"temporal-exp/src/prelock/activity"
	"time"

	"go.temporal.io/sdk/workflow"
)

// PreLockSeatsWorkflow 预锁座位工作流
func PreLockSeatsWorkflow(ctx workflow.Context) error {
	// 设置活动选项
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 30,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// 1. 并行添加订单到carpool池和pre_lock池
	future1 := workflow.ExecuteActivity(ctx, activity.AddOrderToCarpoolPool)
	future2 := workflow.ExecuteActivity(ctx, activity.AddOrderToPreLockPool)

	err := future1.Get(ctx, nil)
	if err != nil {
		return err
	}
	err = future2.Get(ctx, nil)
	if err != nil {
		return err
	}

	// 2. 初始化上下文
	err = workflow.ExecuteActivity(ctx, activity.InitCtx).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 3. 并行从carpool池和pre_lock池召回订单
	future3 := workflow.ExecuteActivity(ctx, activity.RecallOrderFromCarpoolPool)
	future4 := workflow.ExecuteActivity(ctx, activity.RecallOrderFromPreLockPool)

	err = future3.Get(ctx, nil)
	if err != nil {
		return err
	}
	err = future4.Get(ctx, nil)
	if err != nil {
		return err
	}

	// 4. 锁定班次
	err = workflow.ExecuteActivity(ctx, activity.LockShift).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 5. 并行执行获取foras信息、设置班次版本和获取预锁库存
	future5 := workflow.ExecuteActivity(ctx, activity.GetForasInfo)
	future6 := workflow.ExecuteActivity(ctx, activity.SetShiftVersion)
	future7 := workflow.ExecuteActivity(ctx, activity.GetPrelockInventory)

	err = future5.Get(ctx, nil)
	if err != nil {
		return err
	}
	err = future6.Get(ctx, nil)
	if err != nil {
		return err
	}
	err = future7.Get(ctx, nil)
	if err != nil {
		return err
	}

	// 6. 获取班次库存（根据CGraph中的条件选择，这里选择从Redis获取）
	err = workflow.ExecuteActivity(ctx, activity.GetShiftInventoryFromRedis).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 7. 添加班次订单到stg数据
	err = workflow.ExecuteActivity(ctx, activity.AddShiftOrderToStgData).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 8. 获取实时特征
	err = workflow.ExecuteActivity(ctx, activity.GetRtFeature).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 9. 尝试占用座位
	err = workflow.ExecuteActivity(ctx, activity.TryOccupySeats).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 10. API检查
	err = workflow.ExecuteActivity(ctx, activity.ApiCheck).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 11. 锁定座位
	err = workflow.ExecuteActivity(ctx, activity.LockSeats).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 12. 并行执行第一次解锁班次和从池中删除订单
	future8 := workflow.ExecuteActivity(ctx, activity.UnlockShift)

	// 并行从carpool池和pre_lock池删除订单
	future9 := workflow.ExecuteActivity(ctx, activity.DelOrderFromCarpoolPool)
	future10 := workflow.ExecuteActivity(ctx, activity.DelOrderFromPreLockPool)

	err = future9.Get(ctx, nil)
	if err != nil {
		return err
	}
	err = future10.Get(ctx, nil)
	if err != nil {
		return err
	}

	err = future8.Get(ctx, nil)
	if err != nil {
		return err
	}

	// 13. 第二次解锁班次
	err = workflow.ExecuteActivity(ctx, activity.UnlockShift).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 14. 已经锁定处理
	err = workflow.ExecuteActivity(ctx, activity.AlreadyLocked).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
