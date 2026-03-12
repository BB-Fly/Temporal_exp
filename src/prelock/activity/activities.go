package activity

import (
	"fmt"
	"time"
)

// AddOrderToCarpoolPool 添加订单到carpool池
func AddOrderToCarpoolPool() error {
	fmt.Println("AddOrderToCarpoolPool run")
	time.Sleep(100 * time.Millisecond)
	return nil
}

// AddOrderToPreLockPool 添加订单到pre_lock池
func AddOrderToPreLockPool() error {
	fmt.Println("AddOrderToPreLockPool run")
	time.Sleep(50 * time.Millisecond)
	return nil
}

// InitCtx 初始化上下文
func InitCtx() error {
	fmt.Println("InitCtx run")
	time.Sleep(5 * time.Millisecond)
	return nil
}

// RecallOrderFromCarpoolPool 从carpool池中召回订单
func RecallOrderFromCarpoolPool() error {
	fmt.Println("RecallOrderFromCarpoolPool run")
	time.Sleep(180 * time.Millisecond)
	return nil
}

// RecallOrderFromPreLockPool 从pre_lock池中召回订单
func RecallOrderFromPreLockPool() error {
	fmt.Println("RecallOrderFromPreLockPool run")
	time.Sleep(30 * time.Millisecond)
	return nil
}

// LockShift 锁定班次
func LockShift() error {
	fmt.Println("LockShift run")
	time.Sleep(10 * time.Millisecond)
	return nil
}

// GetForasInfo 获取foras信息
func GetForasInfo() error {
	fmt.Println("GetForasInfo run")
	time.Sleep(10 * time.Millisecond)
	return nil
}

// SetShiftVersion 设置班次版本
func SetShiftVersion() error {
	fmt.Println("SetShiftVersion run")
	time.Sleep(10 * time.Millisecond)
	return nil
}

// GetShiftInventoryFromRedis 从Redis获取班次库存
func GetShiftInventoryFromRedis() error {
	fmt.Println("GetShiftInventoryFromRedis run")
	time.Sleep(100 * time.Millisecond)
	return nil
}

// GetShiftInventoryFromDB 从DB获取班次库存
func GetShiftInventoryFromDB() error {
	fmt.Println("GetShiftInventoryFromDB run")
	time.Sleep(120 * time.Millisecond)
	return nil
}

// GetPrelockInventory 获取预锁库存
func GetPrelockInventory() error {
	fmt.Println("GetPrelockInventory run")
	time.Sleep(10 * time.Millisecond)
	return nil
}

// AddShiftOrderToStgData 添加班次订单到stg数据
func AddShiftOrderToStgData() error {
	fmt.Println("AddShiftOrderToStgData run")
	time.Sleep(100 * time.Millisecond)
	return nil
}

// GetRtFeature 获取实时特征
func GetRtFeature() error {
	fmt.Println("GetRtFeature run")
	time.Sleep(80 * time.Millisecond)
	return nil
}

// TryOccupySeats 尝试占用座位
func TryOccupySeats() error {
	fmt.Println("TryOccupySeats run")
	time.Sleep(5 * time.Millisecond)
	return nil
}

// ApiCheck API检查
func ApiCheck() error {
	fmt.Println("ApiCheck run")
	time.Sleep(50 * time.Millisecond)
	return nil
}

// DelOrderFromPreLockPool 从pre_lock池中删除订单
func DelOrderFromPreLockPool() error {
	fmt.Println("DelOrderFromPreLockPool run")
	time.Sleep(50 * time.Millisecond)
	return nil
}

// DelOrderFromCarpoolPool 从carpool池中删除订单
func DelOrderFromCarpoolPool() error {
	fmt.Println("DelOrderFromCarpoolPool run")
	time.Sleep(50 * time.Millisecond)
	return nil
}

// LockSeats 锁定座位
func LockSeats() error {
	fmt.Println("LockSeats run")
	time.Sleep(30 * time.Millisecond)
	return nil
}

// UnlockShift 解锁班次
func UnlockShift() error {
	fmt.Println("UnlockShift run")
	time.Sleep(10 * time.Millisecond)
	return nil
}

// AlreadyLocked 已经锁定
func AlreadyLocked() error {
	fmt.Println("AlreadyLocked run")
	time.Sleep(10 * time.Millisecond)
	return nil
}
