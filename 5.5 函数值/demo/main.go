package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		doc, err := html.Parse(strings.NewReader(string(b)))
		if err != nil {
			fmt.Fprint(os.Stderr, "findlinks1:%v\n", err)
			os.Exit(1)
		}

		forEachNode(doc, startElement, endElement)
	}
}

func forEachNode (n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		//%*s中*会在 字符串 之前填充一些空格
		fmt.Printf("%*s<%s>\n",depth*2,"", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n",depth*2, "", n.Data)
	}
}

/*在golang中，函数被看作第一类值(first-class values)：函数像其他值一样，拥有类型，可以被
赋值给其他变量传递给函数，从函数返回，对函数值的调用类似对函数调用，例如：
func square(n int) int {return n * n}
func square(n int) int {return -n}
func product(m, n int) int {return m * n}

f := square
fmt.Println(f(3)) // 9

f = negative
fmt.Println(f(3))
fmt.Printf("%T\n", f)

f = product //compile error: can't assign func(int, int) int to func(int) int
函数类型的零值是nil。调用值为nil的函数值会引起panic错误：
var f func(int) int
f(3) //此处f的值为nil，会引起panic错误
函数值可以与nil比较
var f func(int) int
if f != nil {
	f(3)
}
但是函数值之间是不可以比较的，也不能用函数值map的key*/
