package jsonFormat

/*
 *@Author georgezhangjc@163.com
 *@Date 2021/12/29 下午3:45
 */

import (
	"fmt"
	"time"
)

type JsonTime time.Time

// 实现它的json序列化方法
func (j JsonTime) MarshalJSON() ([]byte, error) {
	realTime := time.Time(j).Unix()
	stamp := fmt.Sprintf("%d", realTime)
	return []byte(stamp), nil
}
