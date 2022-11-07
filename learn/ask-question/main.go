package main

import (
	"fmt"
	"math"
)

func main() {
	//lat1 := 23.1378010917
	//lng1 := 113.4022203113
	//lat2 := 22.1191433172
	//lng2 := 113.5826193044
	//fmt.Println(GetDistance(lat1, lng1, lat2, lng2))

	lat02 := 23.1378010917
	lng02 := 113.4022203113
	fmt.Println("02：", lat02, lng02)
	lat84, lng84 := gcj02ToWgs84(lat02, lng02)
	fmt.Println("84：", lat84, lng84)
	lat02_2, lng02_2 := wgs84ToGcj02(lat84, lng84)
	fmt.Println("02-2：", lat02_2, lng02_2)
	lat84_2, lng84_2 := gcj02ToWgs84(lat02_2, lng02_2)
	fmt.Println("84-2：", lat84_2, lng84_2)

	fmt.Println(GetDistance(lat02, lng02, lat02_2, lng02_2))
	fmt.Println(GetDistance(lat84, lng84, lat84_2, lng84_2))

	// 经度范围：73°33′E至135°05′E。
	// 纬度范围：3°51′N至53°33′N。
	fmt.Println(gcj02ToWgs84(3.51, 73.33))
	fmt.Println(gcj02ToWgs84(53.33, 135.05))
}

// https://blog.csdn.net/malimingwq/article/details/114950050
// 计算两经纬度之间距离
// select ST_Distance(ST_GeomFromEWKT('SRID=4326;POINT(lng1 lat1)') :: geography,ST_SetSRID(ST_Point(lng2::double precision, lat2::double precision), 4326) :: geography)
// 计算(lng2, lat2)是否在(lng1, lat1)的X米之内
// select ST_DWithin(ST_SetSRID(ST_Point(lng1::double precision, lat1::double precision), 4326) :: geography,ST_GeomFromEWKT('SRID=4326;POINT(lng2 lat2)') :: geography, X)

// GetDistance 返回单位为：米
// lat 维度  lng 经度
// 自己参考网上公式写的代码 https://zhuanlan.zhihu.com/p/99338702 114.90138253718534
func GetDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378137.0 // 6378137.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	deltaLat := lat2 - lat1
	deltaLng := lng2 - lng1
	dist := math.Asin(math.Pow(math.Pow(math.Sin(deltaLat/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(deltaLng/2), 2), 0.5))
	return 2 * dist * radius
}

//// 网上抄的代码 114.901382537215
//
//// GetDistance 返回单位为：千米
//// lat 维度  lng 经度
//func GetDistance(lat1, lng1, lat2, lng2 float64) float64 {
//	radius := 6378137.0 // 6378137.0
//	rad := math.Pi / 180.0
//	lat1 = lat1 * rad
//	lng1 = lng1 * rad
//	lat2 = lat2 * rad
//	lng2 = lng2 * rad
//	theta := lng2 - lng1
//	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
//	return dist * radius / 1000
//}
