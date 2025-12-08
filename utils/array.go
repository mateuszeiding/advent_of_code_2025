package utils

func IndexOf[T comparable](arr []T, searched T) int {
	for i, v := range arr {
		if v == searched {
			return i
		}
	}

	return -1
}

func Dedup[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
