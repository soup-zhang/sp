package maths

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

//DecimalDigits 保留N位小数
func DecimalDigits(value float64, d int) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%." + strconv.Itoa(d) + "f", value), 64)
	return value
}

//GrowthRate 相对增涨率
func GrowthRateFloat64(v1, v2 float64) float64 {
	if v1 == float64(0) || v2 == float64(0) {
		return float64(0)
	}
	ddAmount := v1 - v2
	decimal.DivisionPrecision = 2 //保持2位精度
	v, _ := decimal.NewFromFloat(ddAmount).Mul(decimal.NewFromFloat(100)).Div(decimal.NewFromFloat(v2)).Float64()
	return v
}

//GrowthRateInt64
func GrowthRateInt64(v1, v2 int64) float64 {
	if v1 == int64(0) || v2 == int64(0) {
		return float64(0)
	}
	ddAmount := v1 - v2
	decimal.DivisionPrecision = 2 //保持2位精度
	v, _ := decimal.NewFromFloat(float64(ddAmount)).Mul(decimal.NewFromFloat(100)).Div(decimal.NewFromFloat(float64(v2))).Float64()
	return v
}

//RatioFloat64 比率,保留N位数
func RatioFloat64(v1, v2 float64, p int) float64 {
	if v1 == float64(0) || v2 == float64(0) {
		return float64(0)
	}
	decimal.DivisionPrecision = p //保持2位精度
	v, _ := decimal.NewFromFloat(v1).Mul(decimal.NewFromFloat(100)).Div(decimal.NewFromFloat(v2)).Float64()
	return v
}

//RatioInt64 比率,保留N位数
func RatioInt64(v1, v2 int64, p int) float64 {
	if v1 == int64(0) || v2 == int64(0) {
		return float64(0)
	}
	decimal.DivisionPrecision = p //保持2位精度
	v, _ := decimal.NewFromFloat(float64(v1)).Mul(decimal.NewFromFloat(100)).Div(decimal.NewFromFloat(float64(v2))).Float64()
	return v
}
