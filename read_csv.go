package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	str := "curl 'https://loan.ttco.cn/assets/mutableFee/modify' -H 'Cookie: JSESSIONID=FB9429DF746E6E61753C6CFD90522E10' -H 'Origin: https://loan.ttco.cn' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: zh-CN,zh;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://loan.ttco.cn/v' -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' --data 'editting=true&repaymentPlanNo=@&reasonCode=2&assetInterestValue=#&serviceCharge=&maintenanceCharge=&otherCharge=&comment=' --compressed\n"
	tmp := "echo '@成功'\n"
	outFile := "/Users/louguanyang/Downloads/20190315_xiaomi_curl.sh"
	file, e := os.Open("/Users/louguanyang/Downloads/20190315_xiaomi.csv")
	if e != nil {
		log.Println("csv 读取失败", e.Error())
		return
	}
	defer file.Close()

	out, e := os.OpenFile(outFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if e != nil {
		log.Println("写入文件失败", e.Error())
		return
	}
	defer out.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("csv 读取失败", e.Error())
			return
		}

		str1 := strings.Replace(str, "@", record[0], -1)
		str1 = strings.Replace(str1, "#", record[1], -1)

		_, err = out.WriteString(str1)

		echoStr := strings.Replace(tmp, "@", record[0], -1)

		_, err = out.WriteString(echoStr)

	}

}
