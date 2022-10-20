package main

import (
	qrcode "github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
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

	background := color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
	fground := color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
	_ = qrcode.WriteColorFile("https://h5.duohao.com/?id=ajmkApkZ8w&o=1", qrcode.Medium, 512, background, fground, "qr1.png")

	//drawCirclePic()
}

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
