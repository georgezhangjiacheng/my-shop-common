package sort_utils

import (
	"errors"
	"reflect"
	"sort"
)

/*
 *@Author georgezhangjc@163.com
 *@Date 2022/1/6 下午5:17
 */

// 根据map的key进行排序，返回排序后的keys

func SortStringMapKeys(m interface{}) (ret []string, err error) {
	rv := reflect.ValueOf(m)
	if rv.Kind() != reflect.Map {
		return ret, errors.New("type error, must map")
	}
	if rv.Type().Key().Kind() != reflect.String {
		return ret, errors.New("map key type error, must string")
	}
	var sortList []string
	keys := reflect.ValueOf(m).MapKeys()
	for _, key := range keys {
		a := key.Interface().(string)
		sortList = append(sortList, a)

	}
	sort.Slice(sortList, func(i, j int) bool {
		return sortList[i] < sortList[j]
	})
	return sortList, nil
}

func SortIntMapKeys(m interface{}) (ret []int, err error) {
	rv := reflect.ValueOf(m)
	if rv.Kind() != reflect.Map {
		return ret, errors.New("type error, must map")
	}
	if rv.Type().Key().Kind() != reflect.Int {
		return ret, errors.New("map key type error, must int")
	}
	var sortList []int
	keys := reflect.ValueOf(m).MapKeys()
	for _, key := range keys {
		a := key.Interface().(int)
		sortList = append(sortList, a)

	}
	sort.Slice(sortList, func(i, j int) bool {
		return sortList[i] < sortList[j]
	})
	return sortList, nil
}
