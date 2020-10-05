package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var b []byte

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links{
			fmt.Println(link)
		}
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	//我们必须要确保resp.Body被关闭，释放网络资源。虽然Go的垃圾回收机制会回收不被使用的内存，但是这不包括操作系统层面的资源，比如打开的文件、网络连接。因此我们必须显示的释放这些资源
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("gotting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

//按照惯例，函数的最后一个bool类型的返回值表示函数是否运行成功，error类型的返回值代表函数的错误信息，对于这些返回值，我们无需解释。

/*如果一个函数所有的返回值都有显示的变量名，那么该函数的return语句可以省略操作数，这称之为bare return，bare return虽然可以减少代码重复，但代码的可读性下降*/
