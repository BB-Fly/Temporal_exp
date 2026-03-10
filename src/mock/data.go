package mock

import (
	"temporal-exp/src/models"
	"time"
)

// LoadMockSchedules 加载Mock班次数据
func LoadMockSchedules() []models.Schedule {
	// 这里使用硬编码数据模拟，实际项目中可以从文件读取
	baseTime := time.Now()

	schedules := []models.Schedule{
		{
			ID:             "s1",
			City:           "北京",
			TotalSeats:     50,
			RemainingSeats: 20,
			DepartureTime:  baseTime.Add(2 * time.Hour),
			Stations: []models.Station{
				{
					Name:          "北京站",
					ArrivalTime:   baseTime.Add(2 * time.Hour),
					DepartureTime: baseTime.Add(2*time.Hour + 5*time.Minute),
					Lat:           39.9042,
					Lng:           116.4074,
				},
				{
					Name:          "北京南站",
					ArrivalTime:   baseTime.Add(2*time.Hour + 20*time.Minute),
					DepartureTime: baseTime.Add(2*time.Hour + 25*time.Minute),
					Lat:           39.8652,
					Lng:           116.3786,
				},
				{
					Name:          "天津站",
					ArrivalTime:   baseTime.Add(3*time.Hour + 30*time.Minute),
					DepartureTime: baseTime.Add(3*time.Hour + 35*time.Minute),
					Lat:           39.0842,
					Lng:           117.2009,
				},
			},
		},
		{
			ID:             "s2",
			City:           "北京",
			TotalSeats:     40,
			RemainingSeats: 5,
			DepartureTime:  baseTime.Add(3 * time.Hour),
			Stations: []models.Station{
				{
					Name:          "北京站",
					ArrivalTime:   baseTime.Add(3 * time.Hour),
					DepartureTime: baseTime.Add(3*time.Hour + 5*time.Minute),
					Lat:           39.9042,
					Lng:           116.4074,
				},
				{
					Name:          "北京西站",
					ArrivalTime:   baseTime.Add(3*time.Hour + 20*time.Minute),
					DepartureTime: baseTime.Add(3*time.Hour + 25*time.Minute),
					Lat:           39.8942,
					Lng:           116.3229,
				},
				{
					Name:          "石家庄站",
					ArrivalTime:   baseTime.Add(5*time.Hour + 10*time.Minute),
					DepartureTime: baseTime.Add(5*time.Hour + 15*time.Minute),
					Lat:           38.0452,
					Lng:           114.5149,
				},
			},
		},
		{
			ID:             "s3",
			City:           "北京",
			TotalSeats:     60,
			RemainingSeats: 30,
			DepartureTime:  baseTime.Add(4 * time.Hour),
			Stations: []models.Station{
				{
					Name:          "北京南站",
					ArrivalTime:   baseTime.Add(4 * time.Hour),
					DepartureTime: baseTime.Add(4*time.Hour + 5*time.Minute),
					Lat:           39.8652,
					Lng:           116.3786,
				},
				{
					Name:          "济南站",
					ArrivalTime:   baseTime.Add(6*time.Hour + 20*time.Minute),
					DepartureTime: baseTime.Add(6*time.Hour + 25*time.Minute),
					Lat:           36.6683,
					Lng:           117.0207,
				},
				{
					Name:          "青岛站",
					ArrivalTime:   baseTime.Add(8*time.Hour + 40*time.Minute),
					DepartureTime: baseTime.Add(8*time.Hour + 45*time.Minute),
					Lat:           36.0671,
					Lng:           120.3826,
				},
			},
		},
		{
			ID:             "s4",
			City:           "上海",
			TotalSeats:     55,
			RemainingSeats: 15,
			DepartureTime:  baseTime.Add(2 * time.Hour),
			Stations: []models.Station{
				{
					Name:          "上海站",
					ArrivalTime:   baseTime.Add(2 * time.Hour),
					DepartureTime: baseTime.Add(2*time.Hour + 5*time.Minute),
					Lat:           31.2304,
					Lng:           121.4737,
				},
				{
					Name:          "杭州站",
					ArrivalTime:   baseTime.Add(3*time.Hour + 30*time.Minute),
					DepartureTime: baseTime.Add(3*time.Hour + 35*time.Minute),
					Lat:           30.2741,
					Lng:           120.1551,
				},
			},
		},
		{
			ID:             "s5",
			City:           "北京",
			TotalSeats:     45,
			RemainingSeats: 0,
			DepartureTime:  baseTime.Add(1 * time.Hour),
			Stations: []models.Station{
				{
					Name:          "北京站",
					ArrivalTime:   baseTime.Add(1 * time.Hour),
					DepartureTime: baseTime.Add(1*time.Hour + 5*time.Minute),
					Lat:           39.9042,
					Lng:           116.4074,
				},
				{
					Name:          "天津站",
					ArrivalTime:   baseTime.Add(2*time.Hour + 10*time.Minute),
					DepartureTime: baseTime.Add(2*time.Hour + 15*time.Minute),
					Lat:           39.0842,
					Lng:           117.2009,
				},
			},
		},
	}

	return schedules
}

// LoadMockUserRequest 加载Mock用户请求数据
func LoadMockUserRequest() models.UserRequest {
	return models.UserRequest{
		City:          "北京",
		SeatCount:     2,
		DepartureTime: time.Now().Add(2 * time.Hour),
		PickupLat:     39.9042,
		PickupLng:     116.4074,
		DropoffLat:    39.0842,
		DropoffLng:    117.2009,
		Page:          1,
		PageSize:      10,
	}
}
