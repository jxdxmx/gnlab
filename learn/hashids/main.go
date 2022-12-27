package main

import (
	"fmt"
	"git.gnlab.com/duohao/share.git/crypto_helper"
	"git.gnlab.com/duohao/share.git/hashids"
)

func main() {
	noteId, err := hashids.DecodeNoteID("q48wPeJqQp")
	fmt.Println("note id :", noteId, err)
	//fmt.Println(hashids.EncodeNoteID(1009177))

	//brandID, err := hashids.DecodeUserID("EZYxdVxqYR")
	//fmt.Println("user id", brandID, err)

	brandID, err := hashids.DecodeBrandID("O5YO6byx7e")
	fmt.Println("brand id", brandID, err)

	brandInfoID, err := hashids.DecodeBrandInfoID("RoqnO0bMZY")
	fmt.Println("brand info id", brandInfoID, err)

	//commentID, err := hashids.DecodeCommentID("DLx3JYVGOp")
	//fmt.Println("comment id", commentID, err)

	// genSessionKey 构造session，redis中的值
	// u:1060592 {"nick":"云惜客服佩佩（9:30～18:30）","wechat_ex":1670405449}
	//fmt.Println("token key:", genSessionKey(1060592, 1670405449))

	fmt.Println("encode brand_info_id :")
	fmt.Println(hashids.EncodeBrandInfoID(420734))
}

func test() {
	//输入：s = "the sky is blue"
	//输出："blue is sky the"
	//示例 2：
	//
	//输入：s = "  hello world  "
	//输出："world hello"
	//解释：反转后的字符串中不能存在前导空格和尾随空格。
	//示例 3：
	//
	//输入：s = "a good   example"
	//输出："example good a"
	//解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
	//

}

//示例 1：
//
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//示例 2:
//
//输入： s = "God Ding"
//输出："doG gniD"

//func reverseWords(s string) string {
//	var ans string
//	start := -1
//	i := 0
//	for ; i < len(s); i++ {
//
//	}
//}

func reverse(bs []byte) []byte {
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return bs
}

func genSessionKey(userId int64, expireAt int64) string {
	encrypt := fmt.Sprintf("%d:%d", userId, expireAt)
	key, _ := crypto_helper.EncryptUserKey(encrypt)
	return "w:" + key
}

//fmt.Println("hello world1")
////hashids.EncodeBrandID()
//infoID, err := hashids.DecodeBrandInfoID("85enz5vGlA") // brand_info_id:"85enz5vGlA"
//fmt.Println("brand_info_id:", infoID, err)
//
//infoID, err = hashids.DecodeBrandID("gwE2wxP6p8")
//fmt.Println("band_id:", infoID, err)

//fmt.Println("brand_id:")
//fmt.Println(hashids.EncodeBrandID(1000006))

//fmt.Println(hashids.DecodeShareID("7AlWeOqAwq"))
//fmt.Println(hashids.DecodeUserID("6LkPln32kD"))

//fmt.Println(hashids.EncodeLinkID(3))
//fmt.Println(hashids.DecodeLinkID("eyqKzGk2nM"))
//
//fmt.Println(hashids.EncodeQrCodeID(3))
//fmt.Println(hashids.DecodeQrCodeID("n0d9wxqBX2"))

//ses, err := session.BrandAdmin("b:2b5cb2ee40a257dadf8e4744130d82d531d5dbc9a3164f8634547b0bf5106abd")
//fmt.Println(ses.BrandID(), ses.OperatorID(), err)

//fmt.Println(hashids.EncodeLinkID(5))
//fmt.Println(hashids.EncodeQrCodeID(5))
//fmt.Println(hashids.EncodeMiniPageID(1000070))
//
//noteId, _ := hashids.DecodeNoteID("En8ewzEnmq")
//fmt.Println("note id :", noteId)
//
//brandID, _ := hashids.DecodeBrandID("gwE2wxP6p8")
//fmt.Println("band_id:", brandID)
//
//userID, _ := hashids.DecodeUserID("GaYlvVG904")
//fmt.Println("user id:", userID)
//
//pageID, _ := hashids.DecodeMiniPageID("jVNrG79aMG")
//fmt.Println("page id:", pageID)

//shareID, _ := hashids.DecodeShareID("XD9nMZ77wM")
//fmt.Println("share id:", shareID)
