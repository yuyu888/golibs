
# DataMultiplex 

 使用bitmap的原理，把多个属性存入到一个bit串（字段）中，减少存储并，方便扩展

### DOC

[https://godoc.org/github.com/yuyu888/golibs/datamultiplex](https://godoc.org/github.com/yuyu888/golibs/datamultiplex)

### 参考阅读

[bitmap原理](https://yuyu888.github.io/posts/2020/12/28/bitmap%E5%8E%9F%E7%90%86.html)

[位运算](https://yuyu888.github.io/posts/2020/12/31/%E4%BD%8D%E8%BF%90%E7%AE%97.html)

[php多态计算](https://yuyu888.github.io/posts/2020/12/22/php%E5%A4%9A%E6%80%81%E8%AE%A1%E7%AE%97.html)

### example

````go
package main

import (
	"fmt"
	"github.com/yuyu888/golibs/datamultiplex"
)

func main() {
	a := datamultiplex.IsAttribute(5, 2)
	b := datamultiplex.AddAttribute(8, 2)
	c := datamultiplex.DelAttribute(7, 2)
	d := datamultiplex.CheckAttribute(10)
	e := datamultiplex.GetAttributeList(53)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
````
