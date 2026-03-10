package workflow

import (
	"temporal-exp/src/models"
	"temporal-exp/src/schedule/activity"
	"time"

	"go.temporal.io/sdk/workflow"
)

// RecommendSchedulesWorkflow 班次推荐工作流
func RecommendSchedulesWorkflow(ctx workflow.Context, request models.UserRequest) (models.RecommendationResult, error) {
	// 设置活动选项
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 30,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// 1. 召回对应城市下的所有班次
	var recalledSchedules []models.Schedule
	err := workflow.ExecuteActivity(ctx, activity.RecallSchedules, request.City).Get(ctx, &recalledSchedules)
	if err != nil {
		return models.RecommendationResult{}, err
	}

	// 检查是否有班次
	if len(recalledSchedules) == 0 {
		return models.RecommendationResult{
			Schedules: []models.Schedule{},
			Total:     0,
			Page:      request.Page,
			PageSize:  request.PageSize,
		}, nil
	}

	// 2. 根据座位数过滤
	var filteredBySeat []models.Schedule
	err = workflow.ExecuteActivity(ctx, activity.FilterBySeatCount, struct {
		Schedules []models.Schedule
		SeatCount int
	}{
		Schedules: recalledSchedules,
		SeatCount: request.SeatCount,
	}).Get(ctx, &filteredBySeat)
	if err != nil {
		return models.RecommendationResult{}, err
	}

	// 检查是否有符合座位数要求的班次
	if len(filteredBySeat) == 0 {
		return models.RecommendationResult{
			Schedules: []models.Schedule{},
			Total:     0,
			Page:      request.Page,
			PageSize:  request.PageSize,
		}, nil
	}

	// 3. 根据乘车时间过滤
	var filteredByTime []models.Schedule
	err = workflow.ExecuteActivity(ctx, activity.FilterByDepartureTime, struct {
		Schedules     []models.Schedule
		DepartureTime time.Time
		TimeThreshold time.Duration
	}{
		Schedules:     filteredBySeat,
		DepartureTime: request.DepartureTime,
		TimeThreshold: 2 * time.Hour, // 时间阈值设为2小时
	}).Get(ctx, &filteredByTime)
	if err != nil {
		return models.RecommendationResult{}, err
	}

	// 检查是否有符合时间要求的班次
	if len(filteredByTime) == 0 {
		return models.RecommendationResult{
			Schedules: []models.Schedule{},
			Total:     0,
			Page:      request.Page,
			PageSize:  request.PageSize,
		}, nil
	}

	// 4. 根据坐标距离过滤
	var filteredByDistance []models.Schedule
	err = workflow.ExecuteActivity(ctx, activity.FilterByDistance, struct {
		Schedules         []models.Schedule
		PickupLat         float64
		PickupLng         float64
		DropoffLat        float64
		DropoffLng        float64
		DistanceThreshold float64
	}{
		Schedules:         filteredByTime,
		PickupLat:         request.PickupLat,
		PickupLng:         request.PickupLng,
		DropoffLat:        request.DropoffLat,
		DropoffLng:        request.DropoffLng,
		DistanceThreshold: 5, // 距离阈值设为5公里
	}).Get(ctx, &filteredByDistance)
	if err != nil {
		return models.RecommendationResult{}, err
	}

	// 检查是否有符合距离要求的班次
	if len(filteredByDistance) == 0 {
		return models.RecommendationResult{
			Schedules: []models.Schedule{},
			Total:     0,
			Page:      request.Page,
			PageSize:  request.PageSize,
		}, nil
	}

	// 5. 根据距离排序
	var sortedSchedules []models.Schedule
	err = workflow.ExecuteActivity(ctx, activity.SortByDistance, struct {
		Schedules  []models.Schedule
		PickupLat  float64
		PickupLng  float64
		DropoffLat float64
		DropoffLng float64
	}{
		Schedules:  filteredByDistance,
		PickupLat:  request.PickupLat,
		PickupLng:  request.PickupLng,
		DropoffLat: request.DropoffLat,
		DropoffLng: request.DropoffLng,
	}).Get(ctx, &sortedSchedules)
	if err != nil {
		return models.RecommendationResult{}, err
	}

	// 6. 分页截断
	var result models.RecommendationResult
	err = workflow.ExecuteActivity(ctx, activity.PaginateSchedules, struct {
		Schedules []models.Schedule
		Page      int
		PageSize  int
	}{
		Schedules: sortedSchedules,
		Page:      request.Page,
		PageSize:  request.PageSize,
	}).Get(ctx, &result)
	if err != nil {
		return models.RecommendationResult{}, err
	}

	return result, nil
}
