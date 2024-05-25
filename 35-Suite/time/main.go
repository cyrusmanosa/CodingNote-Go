package main

import (
	"fmt"
	"os"
	"time"
)

func init() {
	// 打開文件
	file, err := os.OpenFile("time.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("打開文件失敗:", err)
		return
	}
	defer file.Close()

	// 讀取文件內容
	data := make([]byte, 8192)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("讀取文件失敗:", err)
		return
	}
	fmt.Printf("文件內容：\n%s\n", string(data[:count]))
}

func main() {
	fmt.Println(" ")
	//time.Time代表一個納秒精度的時間點
	var t time.Time

	//返回當前時間：2017-10-23 19:16:25.3599073 +0800 CST
	t = time.Now()
	fmt.Printf("%v\n", t)

	//返回所在時區：Local
	fmt.Printf("%v\n", t.Location())

	//返回UTC時間和UTC時區：2017-10-23 11:16:25.3599073 +0000 UTC UTC
	fmt.Printf("%v %v\n", t.UTC(), t.UTC().Location())

	//同上，In()返回指定時區的時間：2017-10-23 11:16:25.3599073 +0000 UTC UTC
	fmt.Printf("%v %v\n", t.In(time.UTC), t.In(time.UTC).Location())

	//返回當地時區的時間：2017-10-23 19:16:25.3599073 +0800 CST Local
	fmt.Printf("%v %v\n", t.Local(), t.Local().Location())

	//根據時間戳返回本地時間，參數分別表示秒數和納秒數
	//2017-02-23 00:13:30 +0800 CST
	t2 := time.Unix(1487780010, 0)
	fmt.Println(t2)

	//根據指定時間返回time.Time，分別指定年，月，日，時，分，秒，納秒，時區
	//2017-05-26 15:30:20 +0800 CST
	t3 := time.Date(2017, time.Month(5), 26, 15, 30, 20, 0, t.Location())
	fmt.Println(t3)

	//格式化輸出時間：2017-10-23 19:16:25
	t4 := time.Now()
	fmt.Println(t4.Format("2006-01-02 15:04:05"))

	//獲取時間信息：2017 October 23
	t5 := time.Now()

	//返回日期：2017 October 23
	fmt.Println(t5.Date())

	//返回年：2017
	fmt.Println(t5.Year())

	//返回月：October
	fmt.Println(t5.Month())

	//返回日：23
	fmt.Println(t5.Day())

	//返回星期：Monday
	fmt.Println(t5.Weekday())

	//返回ISO 9601標準下的年份和星期編號：2017 43
	fmt.Println(t5.ISOWeek())

	//返回時分秒：19 16 25
	fmt.Println(t5.Clock())

	//返回小時：19
	fmt.Println(t5.Hour())

	//返回分鐘：16
	fmt.Println(t5.Minute())

	//返回秒：25
	fmt.Println(t5.Second())

	//返回納秒：381921400
	fmt.Println(t5.Nanosecond())

	//返回一年中對應的天：296
	fmt.Println(t5.YearDay())

	//返回時區的規範名,時區相對於UTC的時間偏移量(秒)：CST 28800
	fmt.Println(t5.Zone())

	//返回時間戳：1508757385
	fmt.Println(t5.Unix())

	//返回納秒時間戳：1508757385381921400
	fmt.Println(t5.UnixNano())

	//時間的比較與計算
	t6 := time.Now()

	//是否零時時間：false
	fmt.Println(t6.IsZero())

	//t6時間在t5時間之後，返回真：true
	fmt.Println(t6.After(t5))

	//t5時間在t6時間之前，返回真：true
	fmt.Println(t5.Before(t6))

	//時間是否相同：true
	fmt.Println(t6.Equal(t6))

	//返回t6加上納秒的時間：2017-10-23 19:16:25.383933 +0800 CST
	fmt.Println(t6.Add(10000))

	//返回兩個時間之差的納秒數：2.0016ms
	fmt.Println(t6.Sub(t5))

	//返回t6加1年，1月，1天的時間：2018-11-24 19:16:25.383923 +0800 CST
	fmt.Println(t6.AddDate(1, 1, 1))

	//時間的序列化
	t7 := time.Now()

	//序列化二進制
	bin, _ := t7.MarshalBinary()

	//反序列化二進制：2017-10-23 19:16:25.383923 +0800 CST
	t7.UnmarshalBinary(bin)
	fmt.Println(t7)

	//序列化json："2017-10-23T19:16:25.383923+08:00"
	json, _ := t7.MarshalJSON()
	fmt.Println(string(json))

	//反序列化json：2017-10-23 19:16:25.383923 +0800 CST
	t7.UnmarshalJSON(json)
	fmt.Println(t7)

	//序列化文本：2017-10-23T19:16:25.383923+08:00
	txt, _ := t7.MarshalText()
	fmt.Println(string(txt))

	//反序列化文本：2017-10-23 19:16:25.383923 +0800 CST
	t7.UnmarshalText(txt)
	fmt.Println(t7)

	//gob編碼：2017-10-23 19:16:25.383923 +0800 CST
	gob, _ := t7.GobEncode()
	t7.GobDecode(gob)
	fmt.Println(t7)

	//時間段time.Duration
	dur := time.Duration(6666666600000)

	//返回字符串表示：1h51m6.6666s
	fmt.Println(dur.String())

	//返回小時表示：1.8518518333333334
	fmt.Println(dur.Hours())

	//返回分鐘表示：111.11111
	fmt.Println(dur.Minutes())

	//返回秒表示：6666.6666
	fmt.Println(dur.Seconds())

	//返回納秒表示：6666666600000
	fmt.Println(dur.Nanoseconds())

	//時區time.Location，返回時區名：local
	fmt.Println(time.Local.String())

	//通過地點名和時間偏移量返回時區：Shanghai
	fmt.Println(time.FixedZone("Shanghai", 800))

	//通過給定時區名稱，返回時區：Asia/Shanghai
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(loc)

	//阻塞當前進程3秒
	time.Sleep(time.Second * 3)

	//定時器time.Timer，創建一個1秒後觸發定時器
	timer1 := time.NewTimer(time.Second * 1)
	<-timer1.C
	fmt.Println("timer1 end")

	//1秒後運行函數
	time.AfterFunc(time.Second*1, func() {
		fmt.Println("wait 1 second")
	})
	time.Sleep(time.Second * 3)

	//打點器time.Ticker，創建一個打點器，在固定1秒內重覆執行
	ticker := time.NewTicker(time.Second)
	num := 1
	for {
		if num > 5 {
			ticker.Stop()
			break
		}
		select {
		case <-ticker.C:
			num++
			fmt.Println("1 second...")
		}
	}
}
