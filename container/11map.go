package container

import "fmt"

// Map map演示
func Map() {
	/*
		map：映射，是一种专门用于存储键值对的集合。属于引用类型
		存储特点：
			A:存储的是无序的键值对
			B:键不能重复，并且和值一一对应的。
				map 中的 key 不能重复，如果重复，那么新的 value 会覆盖原来的，程序不会报错。
		语法结构：
			1. 创建 map
				var map1 map[key类型]value类型
					nil map，无法直接使用

				var map2 = make(map[key类型])value类型

				var map3 = map[key类型]value类型{key:value,key:value,key:value}

			2. 添加/修改
				map[key] = value
					如果 key 不存在，就是添加数据
					如果 key 存在，就是修改数据

			3. 获取
				map[key] --> value
				value, ok := map[key]
					根据 key 获取对应的 value
					如果 key 存在，value 就是对应的数据，ok 为 true
					如果 key 不存在，value 就是值类型的默认值，ok 为 false

			4. 删除
				delete(map,key)
					如果 key 存在，就可以直接删除
					如果 key 不存在，删除失效，程序不会报错

			5. 长度
				len()
		每种数据类型：
			int:0
			float:0.0 --> 0
			string:""
			array:[00000]

			slice: nil
			map:nil
	*/

	// 1.创建 map
	var map1 map[int]string         // 没有初始化， nil
	var map2 = make(map[int]string) //创建
	var map3 = map[string]int{"GO": 98, "Python": 87, "Java": 79, "Html": 93}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	// 2.nil map
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	// 3.存储键值对到 map 中
	// map1[key] = value
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "memeda"
	map1[4] = "王二狗"

	// 4.获取数据，根据 key 获取对应的 value 值
	// 根据 key 获取对应的 value， 如果 key 存在，获取数值，如果 key 不存在，获取的是 value 值类型的默认值
	fmt.Println(map1)
	fmt.Println(map1[4])  // 王二狗
	fmt.Println(map1[40]) // ""

	v1, ok := map1[40]
	if ok {
		fmt.Println("对应的数值是：", v1)
	} else {
		fmt.Println("操作的 key 不存在获取到的是默认的零值：", v1)
	}

	// 5.修改
	fmt.Println(map1)
	map1[3] = "如花"
	fmt.Println(map1)

	// 6.删除
	delete(map1, 3)
	fmt.Println(map1)
	delete(map1, 30)
	fmt.Println(map1)

	// 7.长度
	fmt.Println(len(map1))
}
