package main

import (
	"fmt"
	"git.gnlab.com/duohao/share.git/hashids"
	"google.golang.org/grpc/resolver"
)

func main() {
	addr := resolver.Address{Addr: "1.117.8.225:9101"}
	fmt.Println(addr.Addr)

	//fmt.Println("hello world1")
	////hashids.EncodeBrandID()
	//infoID, err := hashids.DecodeBrandInfoID("85enz5vGlA") // brand_info_id:"85enz5vGlA"
	//fmt.Println("brand_info_id:", infoID, err)
	//
	//infoID, err = hashids.DecodeBrandID("gwE2wxP6p8")
	//fmt.Println("band_id:", infoID, err)
	//
	//fmt.Println(hashids.EncodeBrandID(0))
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

	shareID, _ := hashids.DecodeShareID("XD9nMZ77wM")
	fmt.Println("share id:", shareID)
}
