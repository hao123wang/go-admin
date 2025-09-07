// 时间工具类
package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 使用带双引号的日期模板，以符合JSON标准
var formatTime = `"2006-01-02 15:04:05"`

type HTime struct {
	time.Time
}

// MarshalJOSN 和 UnmarshalJSON 实现 HTime 类型与前端 JSON 交互
func (t HTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(formatTime)), nil
}

func (t *HTime) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(formatTime, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = HTime{Time: now}
	return nil
}

// Value 和 Scan 实现 HTime 类型与数据库交互的序列化/反序列化
func (t HTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *HTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	*t = HTime{Time: value}
	return nil
}
