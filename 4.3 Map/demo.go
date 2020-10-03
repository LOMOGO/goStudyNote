package main

import (
	"fmt"
	"sort"
)

func main()  {
	/*哈希表是一个无序的key/value对的集合，其中所有的key都是不同的，通过给定的key可以在检索、更新、或删除对应的value
	在golang中，一个map就是一个哈希表的引用，map类型可以写成map[K]V其中K和V分别对应着key和value。map中所以的key
	都有相同的类型，但是key和value之间可以是不同的数据类型。其中key必须是支持==比较运算符的数据类型，因为map需要根据
	key的值是否相等来判断元素是否存在。虽然浮点数类型也支持相等运算符比较，但是不要使用。V则没有任何限制*/

	//内置的make函数可以创建一个map：
	m1 := make(map[string]int)
	m1["alice"] = 31
	m1["pop"] = 13
	m1["bob"] = 34
	//我们也可以通过map字面值的语法来创建map，同时还可以指定一些最初的key/value，这等价于上面的写法。
	_ = map[string]int{
		"alice": 31,
		"bob": 34,
	}
	//另一种创建空的map的表达式是map[string]int{}
	_ = map[string]int{}

	//map中的元素可以通过key对应的下标语法访问：
	fmt.Println(m1["bob"])

	//通过内置的函数delete()可以删除元素：
	delete(m1, "bob")

	//即使删除的元素不在map中也没关系，依旧安全。
	delete(m1, "top")

	//如果查找失败(不存在这个键)将返回value类型对应的零值。
	m1["bob"]++
	fmt.Println(m1["bob"])//1
	//但是map中的元素不是一个变量，因此我们不能对map的元素进行取地址操作(常量也是不可以取值的)。禁止对map元素取地址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效
	//error： _ = &m1["bob"]

	//map的遍历可以通过range风格的for循环实现。但是map的迭代顺序是不确定的。
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	//如果想要顺序的遍历key/value，我们必须显式的对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序，以下是常见的处理方式
	names := make([]string, 0, len(m1))

	//map类型的range循环可以只用一个值接收，这个时候这个值是map的key
	for name := range m1{
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name, m1[name])
	}

	//map类型的零值是nil，也就是没有引用任何哈希表
	var m2 map[string]int//m2 := map[string]int{}创建的map的值并不为nil
	fmt.Println(m2)
	fmt.Println(m2 == nil)//true

	//向一个nil值的map存入数据会导致一个panic异常：以下代码就会产生panic异常
	//m2["bob"] = 34 panic: assignment to entry in nil map

	//通过key作为索引下标来访问map将产生一个value。如果key在map中是存在的，那么将得到与key对应的value；如果key不存在，那么将得到value对应类型的零值，但如果一元素的开始就存在零的话，这种方法就会失效
	//或者可以另一中方式测试：(在这种场景下map的下标语法将产生两个值；第二个是一个布尔值，用来报告元素是否真的存在。)
	if value, ok := m1["clk"]; !ok {
		fmt.Println("m1[\"clk\"]不存在")
		fmt.Println(value)
	} else {
		fmt.Println(value)
	}
	fmt.Println(equal(map[string]int{"a":0}, map[string]int{"b":42}))

	//测试修改map
	fixMap(m1)
	fmt.Println(m1)

	//测试修改[]slice
	s := make([]string, 0, 20)
	s = []string{"432", "gsjl"}
	fmt.Println(s)
	fixSlice(s)
	fmt.Println(s)

}

//和slice一样，map之间也不能进行相等比较；唯一的例外就是和nil进行比较。但我们可以通过以下函数来判断两个map是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
		//不可以简单使用xv != y[k] 来判断两个map是否相等，因为这在比较map[string]int{"a":0}, map[string]int{"b":42}这样的两个map时会出错。
		/*if xv != y[k] {
			return false
		}*/
	}
	return true
}

func fixMap(m map[string]int) map[string]int {
	m["pop"] = 14
	return m
}

func fixSlice(s []string) []string {
	s[0] = "sdfjs"
	//fixSlice传递的是s的指针变量，因此第一个s[0]修改操作是成功的，append操作没有成功的原因是append返回的可能是一个新的地址的指针，因此失效，正确的写法是使用指针
	s = append(s, "ertw")
	return s
}
