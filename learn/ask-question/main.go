package main

import (
	"fmt"
	"math"
)

func main() {
	lat1 := 23.1378010917
	lng1 := 113.4022203113
	lat2 := 22.1191433172
	lng2 := 113.5826193044
	fmt.Println(GetDistance(lat1, lng1, lat2, lng2))

}

// https://blog.csdn.net/malimingwq/article/details/114950050
// 计算两经纬度之间距离
// select ST_Distance(ST_GeomFromEWKT('SRID=4326;POINT(lng1 lat1)') :: geography,ST_SetSRID(ST_Point(lng2::double precision, lat2::double precision), 4326) :: geography)
// 计算(lng2, lat2)是否在(lng1, lat1)的X米之内
// select ST_DWithin(ST_SetSRID(ST_Point(lng1::double precision, lat1::double precision), 4326) :: geography,ST_GeomFromEWKT('SRID=4326;POINT(lng2 lat2)') :: geography, X)

// GetDistance 返回单位为：千米
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
	return 2 * dist * radius / 1000
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
