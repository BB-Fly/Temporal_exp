package activity

import (
	"context"
	"math"
	"sort"
	"temporal-exp/src/mock"
	"temporal-exp/src/models"
	"time"
)

// RecallSchedules 召回对应城市下的所有班次
func RecallSchedules(ctx context.Context, city string) ([]models.Schedule, error) {
	schedules := mock.LoadMockSchedules()
	var result []models.Schedule

	for _, schedule := range schedules {
		if schedule.City == city {
			result = append(result, schedule)
		}
	}

	return result, nil
}

// FilterBySeatCount 根据座位数过滤班次
func FilterBySeatCount(ctx context.Context, params struct {
	Schedules []models.Schedule
	SeatCount int
}) ([]models.Schedule, error) {
	var result []models.Schedule

	for _, schedule := range params.Schedules {
		if schedule.RemainingSeats >= params.SeatCount {
			result = append(result, schedule)
		}
	}

	return result, nil
}

// FilterByDepartureTime 根据乘车时间过滤班次
func FilterByDepartureTime(ctx context.Context, params struct {
	Schedules     []models.Schedule
	DepartureTime time.Time
	TimeThreshold time.Duration
}) ([]models.Schedule, error) {
	var result []models.Schedule

	for _, schedule := range params.Schedules {
		// 计算时间差的绝对值
		timeDiff := schedule.DepartureTime.Sub(params.DepartureTime)
		if timeDiff < 0 {
			timeDiff = -timeDiff
		}

		if timeDiff <= params.TimeThreshold {
			result = append(result, schedule)
		}
	}

	return result, nil
}

// CalculateDistance 计算两点之间的距离（公里）
func CalculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371 // 地球半径（公里）

	lat1Rad := lat1 * math.Pi / 180
	lng1Rad := lng1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lng2Rad := lng2 * math.Pi / 180

	dLat := lat2Rad - lat1Rad
	dLng := lng2Rad - lng1Rad

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLng/2)*math.Sin(dLng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c

	return distance
}

// FilterByDistance 根据坐标距离阈值过滤班次
func FilterByDistance(ctx context.Context, params struct {
	Schedules         []models.Schedule
	PickupLat         float64
	PickupLng         float64
	DropoffLat        float64
	DropoffLng        float64
	DistanceThreshold float64
}) ([]models.Schedule, error) {
	var result []models.Schedule

	for _, schedule := range params.Schedules {
		// 检查是否有站点在距离阈值内
		valid := false
		for _, station := range schedule.Stations {
			// 计算上车点到站点的距离
			pickupDistance := CalculateDistance(params.PickupLat, params.PickupLng, station.Lat, station.Lng)
			// 计算下车点到站点的距离
			dropoffDistance := CalculateDistance(params.DropoffLat, params.DropoffLng, station.Lat, station.Lng)

			if pickupDistance <= params.DistanceThreshold || dropoffDistance <= params.DistanceThreshold {
				valid = true
				break
			}
		}

		if valid {
			result = append(result, schedule)
		}
	}

	return result, nil
}

// SortByDistance 根据距离排序班次
func SortByDistance(ctx context.Context, params struct {
	Schedules  []models.Schedule
	PickupLat  float64
	PickupLng  float64
	DropoffLat float64
	DropoffLng float64
}) ([]models.Schedule, error) {
	// 创建一个带有距离信息的结构体
	type ScheduleWithDistance struct {
		Schedule      models.Schedule
		TotalDistance float64
	}

	var schedulesWithDistance []ScheduleWithDistance

	for _, schedule := range params.Schedules {
		// 计算班次的平均距离
		totalDistance := 0.0
		stationCount := len(schedule.Stations)

		for _, station := range schedule.Stations {
			pickupDistance := CalculateDistance(params.PickupLat, params.PickupLng, station.Lat, station.Lng)
			dropoffDistance := CalculateDistance(params.DropoffLat, params.DropoffLng, station.Lat, station.Lng)
			totalDistance += pickupDistance + dropoffDistance
		}

		averageDistance := totalDistance / float64(stationCount)
		schedulesWithDistance = append(schedulesWithDistance, ScheduleWithDistance{
			Schedule:      schedule,
			TotalDistance: averageDistance,
		})
	}

	// 根据距离排序
	sort.Slice(schedulesWithDistance, func(i, j int) bool {
		return schedulesWithDistance[i].TotalDistance < schedulesWithDistance[j].TotalDistance
	})

	// 提取排序后的班次
	var result []models.Schedule
	for _, swd := range schedulesWithDistance {
		result = append(result, swd.Schedule)
	}

	return result, nil
}

// PaginateSchedules 对班次进行分页截断
func PaginateSchedules(ctx context.Context, params struct {
	Schedules []models.Schedule
	Page      int
	PageSize  int
}) (models.RecommendationResult, error) {
	total := len(params.Schedules)
	start := (params.Page - 1) * params.PageSize
	end := start + params.PageSize

	if start >= total {
		return models.RecommendationResult{
			Schedules: []models.Schedule{},
			Total:     total,
			Page:      params.Page,
			PageSize:  params.PageSize,
		}, nil
	}

	if end > total {
		end = total
	}

	return models.RecommendationResult{
		Schedules: params.Schedules[start:end],
		Total:     total,
		Page:      params.Page,
		PageSize:  params.PageSize,
	}, nil
}
