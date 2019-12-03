package container

import "fmt"

// Map3 map 结合 slice 使用
func Map3() {
	/*
		map 和 slice 的结合使用：
			1.创建 map 用于存储人的信息
				name, age, sex, address

			2.每个 map 存储一个人的信息
			3.将这些 map 存储到 slice 中
			4.遍历打印输出
	*/

	// 1.创建 map 用于存储人的信息
	map1 := make(map[string]string)
	map1["name"] = "王二狗"
	map1["age"] = "30"
	map1["sex"] = "男"
	map1["address"] = "北京市"

	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["name"] = "李小花"
	map2["age"] = "20"
	map2["sex"] = "女"
	map2["address"] = "上海市"

	fmt.Println(map2)

	map3 := make(map[string]string)
	map3["name"] = "ruby"
	map3["age"] = "30"
	map3["sex"] = "女"
	map3["address"] = "杭州市"

	fmt.Println(map3)

	// 将 map 存入到 slice 中
	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)

	// 遍历切片
	for i, val := range s1 {
		//val :map1,map2.map3
		fmt.Printf("第%d个人的信息是\n", i+1)
		fmt.Printf("\t姓名：%s\n", val["name"])
		fmt.Printf("\t年龄：%s\n", val["age"])
		fmt.Printf("\t性别：%s\n", val["sex"])
		fmt.Printf("\t地址：%s\n", val["address"])
	}
}
