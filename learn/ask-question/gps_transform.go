package main

import (
	"math"
)

const (
	a  float64 = 6378245.0
	ee float64 = 0.00669342162296594323
)

// 坐标系相互转换的在线工具：https://tool.lu/coordinate/
// 经测试，本代码和网上在线工具的实际运行结果完全相同

// wgs84ToGcj02 84 to 火星坐标系 (GCJ-02) World Geodetic System ==> Mars Geodetic System
func wgs84ToGcj02(lat84, lng84 float64) (lat02, lng02 float64) {
	if outOfChina(lat84, lng84) {
		return lat84, lng84
	}
	dLat := transformLat(lng84-105.0, lat84-35.0)
	dLng := transformLng(lng84-105.0, lat84-35.0)
	radLat := lat84 / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * math.Pi)
	dLng = (dLng * 180.0) / (a / sqrtMagic * math.Cos(radLat) * math.Pi)
	return lat84 + dLat, lng84 + dLng
}

// gcj02ToWgs84 火星坐标系 (GCJ-02) to 84
func gcj02ToWgs84(lat02, lng02 float64) (lat84, lng84 float64) {
	if outOfChina(lat02, lng02) {
		return lat02, lng02
	}
	dLat := transformLat(lng02-105.0, lat02-35.0)
	dLng := transformLng(lng02-105.0, lat02-35.0)
	radLat := lat02 / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * math.Pi)
	dLng = (dLng * 180.0) / (a / sqrtMagic * math.Cos(radLat) * math.Pi)
	lat, lng := lat02+dLat, lng02+dLng
	return lat02*2 - lat, lng02*2 - lng
}

func transformLat(x, y float64) float64 {
	ret := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*math.Pi) + 40.0*math.Sin(y/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*math.Pi) + 320*math.Sin(y*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLng(x, y float64) float64 {
	ret := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*math.Pi) + 40.0*math.Sin(x/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*math.Pi) + 300.0*math.Sin(x/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}

// 中国的经纬度范围
// 经度范围：73°33′E至135°05′E 纬度范围：3°51′N至53°33′N。
func outOfChina(lat, lng float64) bool {
	if lng < 73.33 || lng > 135.05 {
		return true
	}
	if lat < 3.51 || lat > 53.33 {
		return true
	}
	return false
}
