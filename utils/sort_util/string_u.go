package sort_utils

/**
 *@Author Imy@relaper.com
 *@Date 2021/11/24 11:13
 */

type SS []string

func (s SS) Len() int {
	return len(s)
}

func (s SS) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SS) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
