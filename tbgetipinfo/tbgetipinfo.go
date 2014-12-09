package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type TbApi struct {
	Code float64
	Data DataObj
}

type DataObj struct {
	Country    string
	Country_id string
	Area       string
	Area_id    float64
	Region     string
	Region_id  float64
	City       string
	City_id    float64
	County     string
	County_id  float64
	Isp        string
	Isp_id     float64
	Ip         string
}

func getipinfo(ipaddr string) {
	res, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ipaddr)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var tbapiinfo TbApi
	json.Unmarshal(body, &tbapiinfo)
	fmt.Printf("IP地址:%s 状态:%v 国家:%s 省:%s 城市:%s 运营商:%s\n", tbapiinfo.Data.Ip, tbapiinfo.Code, tbapiinfo.Data.County, tbapiinfo.Data.Region, tbapiinfo.Data.City, tbapiinfo.Data.Isp)
}

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	//fmt.Printf("%v \n",os.Args[1])
	var para_num int = len(os.Args)
	//fmt.Println(para_num)
	if para_num > 1 {
		if IsIP(os.Args[1]) {
			getipinfo(os.Args[1])
		} else {
			fmt.Println("您输入的不是正确的ip地址")
		}
	} else {
		fmt.Println("使用方法：tbgetipinfo 参数")
	}
}
