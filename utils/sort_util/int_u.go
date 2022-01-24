package sort_utils

/**
 *@Author Imy@relaper.com
 *@Date 2021/11/24 11:20
 */

type IntS []int

func (ii IntS) Len() int {
	return len(ii)
}

func (ii IntS) Less(i, j int) bool {
	return ii[i] < ii[j]
}

func (ii IntS) Swap(i, j int) {
	ii[i], ii[j] = ii[j], ii[i]
}
