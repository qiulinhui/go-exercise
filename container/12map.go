package container

import "fmt"

import "sort"

// Map2 遍历 map
func Map2() {

	/*
		map 的遍历：
			使用：for range
			数组，切片：index, value
			map：key, value
	*/

	map1 := make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "小钻风"
	map1[1] = "白骨精"
	map1[1] = "白素贞"
	map1[1] = "金角大王"
	map1[1] = "王二狗"

	// 1.遍历 map，map 是无序的
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println("------------------------------")
	// 顺序输出，如果 map 的 key 是连续的数字的话
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, "--->", map1[i])
	}

	/*
		如果 key 不连续的话
		1.获取所有的 key，-->切片/数组
		2.进行排序
		3.遍历 key，--->map[key]
	*/

	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	// 冒泡排序或者使用 sort 包下的排序方法
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	// 字符串排序
	s1 := []string{"Apple", "Windows", "Orange", "abc", "中文", "acd", "acc"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
