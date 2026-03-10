package models

import "time"

// UserRequest 用户请求信息
type UserRequest struct {
	City          string    `json:"city"`          // 城市
	SeatCount     int       `json:"seatCount"`     // 座位数
	DepartureTime time.Time `json:"departureTime"` // 乘车时间
	PickupLat     float64   `json:"pickupLat"`     // 上车点纬度
	PickupLng     float64   `json:"pickupLng"`     // 上车点经度
	DropoffLat    float64   `json:"dropoffLat"`    // 下车点纬度
	DropoffLng    float64   `json:"dropoffLng"`    // 下车点经度
	Page          int       `json:"page"`          // 页码
	PageSize      int       `json:"pageSize"`      // 每页大小
}

// Station 站点信息
type Station struct {
	Name          string    `json:"name"`          // 站点名称
	ArrivalTime   time.Time `json:"arrivalTime"`   // 预计到达时间
	DepartureTime time.Time `json:"departureTime"` // 预计出发时间
	Lat           float64   `json:"lat"`           // 纬度
	Lng           float64   `json:"lng"`           // 经度
}

// Schedule 班次信息
type Schedule struct {
	ID             string    `json:"id"`             // 班次ID
	City           string    `json:"city"`           // 城市
	TotalSeats     int       `json:"totalSeats"`     // 库存总数
	RemainingSeats int       `json:"remainingSeats"` // 剩余库存数
	DepartureTime  time.Time `json:"departureTime"`  // 发车时间
	Stations       []Station `json:"stations"`       // 站点列表
}

// RecommendationResult 推荐结果
type RecommendationResult struct {
	Schedules []Schedule `json:"schedules"` // 推荐的班次列表
	Total     int        `json:"total"`     // 总数量
	Page      int        `json:"page"`      // 当前页码
	PageSize  int        `json:"pageSize"`  // 每页大小
}
