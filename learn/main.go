package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func getList() []int64 {
	return nil
}

// SubStrEllipsis 截取字符串，如果原字符串超长，返回时拼接"..."
func SubStrEllipsis(s string, length int) string {
	var rs = []rune(s)
	if len(rs) <= length {
		return string(rs)
	}
	return string(rs[:length]) + "..."
}

func main() {
	fmt.Println(SubStrEllipsis("回复评论", 3))
	str := "CommentReply 回复评论"
	fmt.Println(SubStrEllipsis(str, 16))
	fmt.Println(SubStrEllipsis(str, 17))
	fmt.Println(SubStrEllipsis(str, 20))

	//var ids = []int64{1, 2, 3}
	//ids = append(ids, getList()...)
	//fmt.Println(ids)
	//var m = map[string]interface{}{"a": "3"}
	//fmt.Println(m["xx"].(string))

	//type PersistentStatus int
	//const (
	//	PersistentStatusOK PersistentStatus = 10 // 持久化OK
	//)
	//fmt.Println(PersistentStatus(5))

	//eg := errgroup.Group{}
	////var arr = []string{"h1", "h2", "h3"}
	////for _, s := range arr {
	////	var tmp = s
	////	eg.Go(func() error {
	////		return printStr(tmp)
	////	})
	////}
	//s := "h1"
	//eg.Go(func() error { return printStr(s) })
	//s = "h2"
	//eg.Go(func() error { return printStr(s) })
	//s = "h3"
	//eg.Go(func() error { return printStr(s) })
	////for i := 0; i < 1000; i++ {
	////	fmt.Print(RandomInt(1, 100), " ")
	////	if i%30 == 0 {
	////		fmt.Println()
	////	}
	////}
	//eg.Wait()
}
func printStr(s string) error {
	fmt.Println(s)
	return nil
}

// RandomInt64 产生一个随机数,真随机
func RandomInt64(min, max int64) int64 {
	bi := big.NewInt(max - min)
	r, _ := rand.Int(rand.Reader, bi)
	return min + r.Int64()
}

// RandomInt 产生一个随机数,真随机
func RandomInt(min, max int) int {
	return int(RandomInt64(int64(min), int64(max)))
}
