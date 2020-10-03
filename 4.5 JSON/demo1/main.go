package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//JavaScript对象表示法(4.5 JSON)是一种用于发送和接收结构化信息的标准协议。
	//基本的JSON类型有数字(十进制或者科学计数法)、布尔值(true或false)、字符串(以双引号包含的Unicode字符序列)，支持反斜杠转义特性。

	/*一个JSON数组是一个有序的值序列，写在一个方括号中并以逗号分割；一个JSON数组可以用于编码Go语言中的数组和slice。
	一个JSON对象是一个字符串到值的映射，写成一系列的name:value对形式，用括号包含并以逗号分割；JSON的对象类型可以用于编码golang中的map类型(key类型是字符串)和结构体。*/

	//在编码时，默认使用golang结构体成员的名字作为JSON的对象。只有导出的结构体成员才会被编码，如果加入了结构体成员Tag，那么一个结构体成员Tag是和编译阶段关联到该成员的元信息字符串
	type Movie struct {
		Title string
		/*因为值中含有双引号字符，因此成员Tag一般用原生字符串面值的形式书写。json开头键名对应的值用于控制encoding/json包的编码和解码行为，
		  成员tag中json对应值的第一部分用于指定json对象的名字。
		  Color成员还带了一个额外的omitempty选项，表示当golang语言结构体成员为空或者零值时不生成该json对象(这里false为零值)*/
		Year   int  `json:"year"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{
			Title:  "摩登时代",
			Year:   1943,
			Color:  false,
			Actors: []string{"卓别林"},
		},
		{
			Title:  "战狼2",
			Year:   2018,
			Color:  true,
			Actors: []string{"吴京", "达康书记"},
		},
		{
			Title:  "我不是药神",
			Year:   2018,
			Color:  true,
			Actors: []string{"徐峥", "王传君", "曹宇"},
		},
	}
	//上面这样的数据结构特别适合JSON格式，并且在两者之间相互转换也很容易。将一个golang中类似movies的结构体slice转为json的过程叫编组。编组通过json.Marshal函数完成

	//Marshal函数返还一个编码后的字节slice，包含很长的字符串，但并没有空白缩进，不便于阅读
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("4.5 JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//MarshalIndent函数将产生整齐缩进的输出
	data, err = json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("4.5 JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n%T\n", data, data)
	//输出的形式如下，需要注意的一点是在最后一个成员或元素后面并没有逗号分隔符
	_ = `[
        {
                "Title": "摩登时代",
                "year": 1943,
                "Actors": [
                        "卓别林"
                ]
        },
        {
                "Title": "战狼2",
                "year": 2018,
                "color": true,
                "Actors": [
                        "吴京",
                        "达康书记"
                ]
        },
        {
                "Title": "我不是药神",
                "year": 2018,
                "color": true,
                "Actors": [
                        "徐峥",
                        "王传君",
                        "曹宇"
                ]
        }
]
`
	type Movie2 struct {
		Title string
		Color bool `json:"color,omitempty"`
	}
	var movie2 = []Movie2{}
	if err := json.Unmarshal(data, &movie2); err != nil {
		log.Fatalf("4.5 JSON marshaling failed: %s", err)
	}
	fmt.Println(movie2)

}
