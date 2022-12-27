package main

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//var png []byte
	//png, err := qrcode.Encode("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium, 256)
	//q, _ := qrcode.New("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium)
	////qrcode.new

	//err := qrcode.WriteFile("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium, 256, "qr.png")
	//err := qrcode.WriteColorFile("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium, 256, color.Black, color.White, "qr.png")
	//_ = qrcode.WriteColorFile("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium, 512, color.White, color.Black, "qr1.png")
	//_ = qrcode.WriteColorFile("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium, 512, color.Transparent, color.White, "qr2.png")

	logoUrl := "https://tim.duohao.com/b1000364_1669686783_1503443296.webp"
	//logoUrl := "https://tim.duohao.com/default/logo.png"
	fileName := "qr1.png"
	background := color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
	fground := color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
	size := 512
	_ = qrcode.WriteColorFile("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium, size, background, fground, fileName)

	qrImg, err := gg.LoadPNG(fileName)
	if err != nil {
		log.Panic(err)
	}
	ggCtx := gg.NewContextForImage(qrImg)

	bs, err := accessImage(GenQiNiuPngUrl(logoUrl))
	if err != nil {
		log.Panic(err)
	}
	logoImg, _, err := image.Decode(bytes.NewReader(bs))
	if err != nil {
		log.Panic(err)
	}

	//logoImag, err := gg.LoadImage("logo.webp")
	//logoImag, err := gg.LoadImage("test.jpg")
	//if err != nil {
	//	log.Panic(err)
	//}

	//logoGgCtx := gg.NewContextForImage(logoImag)
	//logoGgCtx.Scale(0.01, 0.01)
	//fmt.Println(logoGgCtx.SavePNG("test.png"))
	scaleX, scaleY := size/4, size/4
	left, right := (size-scaleX)/2, (size-scaleY)/2
	ggCtx.DrawImage(DrawCircleImg(ScaleImage(logoImg, scaleX, scaleY)), left, right)
	fmt.Println(ggCtx.SavePNG(fileName))
}

const (
	DuoHaoDomain = ".duohao.com"
)

// GenQiNiuPngUrl 转换图片格式为png
// 如果图片本来并不是七牛url，直接返回
func GenQiNiuPngUrl(url string) string {
	if !strings.Contains(url, DuoHaoDomain) {
		return url
	}
	cmd := "?imageMogr2/format/png"
	if strings.Contains(url, "?") {
		// https://tim.duohao.com/1669603462_3_webp.webp?imageView2/2/w/340/q/50 => https://tim.duohao.com/1669603462_3_webp.webp
		idx := strings.Index(url, "?")
		url = url[idx:]
	}
	return url + cmd
}

// DrawCircleImg 头像圆图
func DrawCircleImg(im image.Image) image.Image {
	loadStart := time.Now()
	defer func() {
		log.Printf("生成圆头像耗时: %v\n", time.Now().Sub(loadStart).Milliseconds())
	}()
	b := im.Bounds()
	w := float64(b.Dx())
	h := float64(b.Dy())
	dc := gg.NewContext(int(w), int(h))
	r := w / 2 // 半径
	dc.DrawRoundedRectangle(0, 0, w, h, r)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	return dc.Image()
}

func ScaleImage(image image.Image, x, y int) image.Image {
	loadStart := time.Now()
	defer func() {
		log.Printf("缩放耗时: %v\n", time.Now().Sub(loadStart).Milliseconds())
	}()
	w := image.Bounds().Size().X
	h := image.Bounds().Size().Y
	dc := gg.NewContext(x, y)
	var ax = float64(x) / float64(w)
	var ay = float64(y) / float64(h)
	dc.Scale(ax, ay)
	dc.DrawImage(image, 0, 0)
	return dc.Image()
}

func accessImage(url string) (bs []byte, err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(url, "access network image error.", err.Error())
		return
	}
	resp, err := client.Do(request) //发送请求
	if err != nil {
		fmt.Println(url, "access network image error.", err.Error())
	}
	defer func() { _ = resp.Body.Close() }() //一定要关闭resp.Body
	bs = make([]byte, 0, resp.ContentLength)
	bs, err = ioutil.ReadAll(resp.Body)
	return
}

//drawCirclePic()

func drawCirclePic() {
	f, err := os.Open("./qr.png")
	if err != nil {
		panic(err)
	}
	gopherImg, _, err := image.Decode(f)
	d := gopherImg.Bounds().Dx()

	//将一个cicle作为蒙层遮罩，圆心为图案中点，半径为边长的一半
	c := circle{p: image.Point{X: d / 2, Y: d / 2}, r: d / 4}
	circleImg := image.NewRGBA(image.Rect(0, 0, d, d))
	draw.DrawMask(circleImg, circleImg.Bounds(), gopherImg, image.Point{}, &c, image.Point{}, draw.Over)
	//SavePng(circleImg)

	f, err = os.Create("qr2.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, circleImg)
	if err != nil {
		panic(err)
	}

}

type circle struct { // 这里需要自己实现一个圆形遮罩，实现接口里的三个方法
	p image.Point // 圆心位置
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}
func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

// At 对每个像素点进行色值设置，在半径以内的图案设成完全不透明
func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{A: 255}
	}
	return color.Alpha{}
}
