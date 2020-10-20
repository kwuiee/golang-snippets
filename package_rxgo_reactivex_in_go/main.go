// https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651442353&idx=3&sn=98e8990b75f6dc3f7e97731dbffa7035&chksm=80bb1243b7cc9b5589ddc6fc051935c9d3f52c64d596f52077407ecc805a9e5092e00fcad362&xtrack=1&scene=90&subscene=93&sessionid=1602840963&clicktime=1602841204&enterid=1602841204&ascene=56&devicetype=android-26&version=2700133f&nettype=3gnet&abtest_cookie=AAACAA%3D%3D&lang=zh_CN&exportkey=AZCeUxbmYPRku%2FYa49VCKr8%3D&pass_ticket=p9B40Z9VjEiL2wxGI%2BsU8%2FzwkD2Ozk%2Be8qAY2fnXg6w2BCd%2FrqqoqqGaoM%2FEoa6s&wx_header=1
package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observer: %d\n", i)
	})

	time.Sleep(3 * time.Second)
	fmt.Println("before subscribe second observer")

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	time.Sleep(3 * time.Second)
}
