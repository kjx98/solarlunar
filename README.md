# solarlunar
[![Build Status](https://travis-ci.org/kjx98/solarlunar.svg?branch=master)](https://travis-ci.org/kjx98/solarlunar)

forked from https://github.com/nosixtools/solarlunar ， 修正time.ParseInLocation带来的润秒偏差，利用儒略日表加速

##### 1.阳历和阴历相互转化（支持1900~2049年）
##### 2.节假日计算

## 快速开始
#### 下载和安装
	go get -u github.com/kjx98/solarlunar
#### 创建 solarlunar.go  阳历和阴历转化
```
package main 


import (
	"github.com/kjx98/solarlunar" 
	"fmt"
)


func main() {
	solarDate := "1990-05-06"
	fmt.Println(solarlunar.SolarToChineseLuanr(solarDate))
	fmt.Println(solarlunar.SolarToSimpleLuanr(solarDate))
	
	lunarDate := "1990-04-12"
	fmt.Println(solarlunar.LunarToSolar(lunarDate, false))
}

```
#### 创建 festival.go 节假日计算
```
package main


import (
"fmt"
"github.com/kjx98/solarlunar/festival"
)

func main() {
	festival := festival.NewFestival("./festival.json")
	fmt.Println(festival.GetFestivals("2017-08-28"))
	fmt.Println(festival.GetFestivals("2017-05-01"))
	fmt.Println(festival.GetFestivals("2017-04-05"))
	fmt.Println(festival.GetFestivals("2017-10-01"))
	fmt.Println(festival.GetFestivals("2018-02-15"))
	fmt.Println(festival.GetFestivals("2018-02-16"))
}
```

### benchmark, go test -bench=.
```
庚午马年四月十二日
1990年04月12日
2019年01月01日
Lunar:  2019 1 1 false
1990-05-06
2019-02-05
Solar:  2019 2 5
goos: linux
goarch: amd64
pkg: github.com/nosixtools/solarlunar
BenchmarkLunarToSolar-4   	  200000	      6113 ns/op
BenchmarkSolarToLunar-4   	  200000	      6370 ns/op
BenchmarkLunar2Solar-4    	20000000	        61.4 ns/op
BenchmarkSolar2Lunar-4    	50000000	        30.8 ns/op
PASS
```

