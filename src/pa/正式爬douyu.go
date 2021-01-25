package pa

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Anchor struct {
	Name      string
	Hot       string
	Section   string
	RoomTitle string
}

var AnchorSlice []Anchor
var douYuDb *sql.DB

/*
	所有分区的后缀
*/
var SectionSuffix = []string{"g_jdqs", "g_DOTA2", "g_CF", "g_DNF", "g_How", "g_CSGO", "g_NZ",
	"g_ydzhy", "g_mszb", "g_wowclassic", "g_dota2rpg", "g_VALORANT",
	"g_Overwatch", "g_DOTA", "g_WOW", "g_JX3", "g_LRS", "g_TVgame",
	"g_EFT", "g_Survival", "g_Hitman", "g_tlbj", "g_SANGO", "g_gwlrsj",
	"g_classic", "g_FTG", "g_wzry", "g_hpjy", "g_hyrz", "g_smzhsy", "g_tdsy",
	"g_hlddz", "g_wzrpg", "g_yuanshen", "g_CFSY", "g_qipai", "g_tysy",
	"g_xyzx", "g_qqfcsy", "g_yys", "g_phone", "g_mdnf", "g_yxwjdlj",
	"g_wxsh", "g_yz", "g_music", "g_dance", "g_ecy", "g_HW", "g_ms",
	"g_jyhd", "g_yqk", "g_QSH", "g_smkj", "g_wh", "g_kepu", "g_zjzg",
	"g_car", "g_DR", "g_DYGW", "directory/sport/cate", "g_jiaoyou",
	"g_DIANT", "g_kaihei", "g_yyzs", "g_znl", "g_jdy"}

/**
 * @Description 斗鱼爬虫的总开始在这
	计划是半小时一次
 **/
func AllDouYuStartPa() {
	rand.Seed(time.Now().UnixNano())

	//【1】首先拿到各个分区的URL
	sectionUrl := GetSectionUrl(SectionSuffix)

	//【2】拿到每个分区的页面以后，我们想要的是具体的房间号，各种数据的爬取都是在房间内进行的
	//	 所以需要把所有分区的所有房间号都存到一个slice里，这里还是用string储存，以便以后拼接具体地址
	//allRoomNums := GetRoomNumsFromSection(sectionUrl)

	//allRoomNums 里面是  https://www.douyu.com/24422
	//	                 https://www.douyu.com/5720533
	//    	             https://www.douyu.com/296059
	//       	         https://www.douyu.com/709710
	//                   https://www.douyu.com/5384600
	//                   https://www.douyu.com/1679347
	//                   https://www.douyu.com/759905
	//这种格式，可以直接遍历爬取
	//fmt.Println(allRoomNums)
	fmt.Println("分区url拼接完成，准备开始爬取数据！")
	GetAnchorInfos(sectionUrl)
}

/*
	输入：分区后缀
	返回：每个分区的URL
*/
func GetSectionUrl(suffix []string) []string {

	var SectionUrl []string
	//前缀不变
	originWebUrl := `https://www.douyu.com/`
	//加上后缀
	for _, suffix := range SectionSuffix {

		//生成每个分区的URL
		totalWebUrl := originWebUrl + suffix
		/*
			https://www.douyu.com/g_jdqs
			https://www.douyu.com/g_DOTA2
			https://www.douyu.com/g_CF
			https://www.douyu.com/g_DNF
			https://www.douyu.com/g_How
			https://www.douyu.com/g_CSGO
			https://www.douyu.com/g_NZ
			https://www.douyu.com/g_ydzhy
			https://www.douyu.com/g_mszb
			https://www.douyu.com/g_wowclassic
			https://www.douyu.com/g_dota2rpg
			https://www.douyu.com/g_VALORANT
			https://www.douyu.com/g_Overwatch
			https://www.douyu.com/g_DOTA
			https://www.douyu.com/g_WOW
			https://www.douyu.com/g_JX3
			https://www.douyu.com/g_LRS
			https://www.douyu.com/g_TVgame
			https://www.douyu.com/g_EFT
			https://www.douyu.com/g_Survival
			https://www.douyu.com/g_Hitman
			https://www.douyu.com/g_tlbj
			https://www.douyu.com/g_SANGO
			https://www.douyu.com/g_gwlrsj
			https://www.douyu.com/g_classic
			https://www.douyu.com/g_FTG
			https://www.douyu.com/g_wzry
			https://www.douyu.com/g_hpjy
			https://www.douyu.com/g_hyrz
			https://www.douyu.com/g_smzhsy
			https://www.douyu.com/g_tdsy
			https://www.douyu.com/g_hlddz
			https://www.douyu.com/g_wzrpg
			https://www.douyu.com/g_yuanshen
			https://www.douyu.com/g_CFSY
			https://www.douyu.com/g_qipai
			https://www.douyu.com/g_tysy
			https://www.douyu.com/g_xyzx
			https://www.douyu.com/g_qqfcsy
			https://www.douyu.com/g_yys
			https://www.douyu.com/g_phone
			https://www.douyu.com/g_mdnf
			https://www.douyu.com/g_yxwjdlj
			https://www.douyu.com/g_wxsh
			https://www.douyu.com/g_yz
			https://www.douyu.com/g_music
			https://www.douyu.com/g_dance
			https://www.douyu.com/g_ecy
			https://www.douyu.com/g_HW
			https://www.douyu.com/g_ms
			https://www.douyu.com/g_jyhd
			https://www.douyu.com/g_yqk
			https://www.douyu.com/g_QSH
			https://www.douyu.com/g_smkj
			https://www.douyu.com/g_wh
			https://www.douyu.com/g_kepu
			https://www.douyu.com/g_zjzg
			https://www.douyu.com/g_car
			https://www.douyu.com/g_DR
			https://www.douyu.com/g_DYGW
			https://www.douyu.com/directory/sport/cate
			https://www.douyu.com/g_jiaoyou
			https://www.douyu.com/g_DIANT
			https://www.douyu.com/g_kaihei
			https://www.douyu.com/g_yyzs
			https://www.douyu.com/g_znl
			https://www.douyu.com/g_jdy
		*/
		//println(totalWebUrl)
		//把生成的存到切片里
		SectionUrl = append(SectionUrl, totalWebUrl)

	}

	fmt.Println("各个分区的URL拼接结束！")
	return SectionUrl
}

/**
 * @Description 从每个分区爬取信息
 * @Param  切片 存着所有的分区地址
 * @return 	不返回 直接存入mysql
 **/
func GetAnchorInfos(SectionUrl []string) {
	err := InitDB()
	HandleError(err, "数据库初始化")

	for secIndex, sectionurl := range SectionUrl {
		response, err := http.Get(sectionurl)
		HandleError(err, "http.Get(SectionUrl[0])")

		bytes, err := ioutil.ReadAll(response.Body)
		HandleError(err, "ioutil.ReadAll(response.Body)")

		htmlContent := string(bytes)

		//爬主播名字
		reZhuBoName := regexp.MustCompile(`"></use></svg><div class="DyListCover-userName is-template">(.*?)</div></h2></div></div></a><a href="`)
		nameRes := reZhuBoName.FindAllStringSubmatch(htmlContent, -1)

		//爬热度
		reHot := regexp.MustCompile(`</h3></div><div class="DyListCover-info"><span class="DyListCover-hot is-template"><svg><use xlink:href="#icon-hot_8a57f0b"></use></svg>(.*?)</span><h2 class="DyListCover-user is-template">`)
		hotRes := reHot.FindAllStringSubmatch(htmlContent, -1)

		//爬分区
		reSectionName := regexp.MustCompile(`<span class="DyListCover-zone">(.*?)</span><h3 class="DyListCover-intro" title="`)
		sectionNameRes := reSectionName.FindAllStringSubmatch(htmlContent, -1)

		//爬直播间标题
		resRoomName := regexp.MustCompile(`<h3 class="DyListCover-intro" title="(.*?)">.*?</h3>`)
		roomNameRes := resRoomName.FindAllStringSubmatch(htmlContent, -1)

		for index := 0; index < len(nameRes)-2; index++ {

			intHot := DealWan(hotRes[index][1])

			_, err = douYuDb.Exec("insert into douyu (name,hot,section,roomtitle) values (?,?,?,?);", nameRes[index][1], intHot, sectionNameRes[index][1], roomNameRes[index][1])
			HandleError(err, "插入语句")
			fmt.Println(nameRes[index][1], "\t", intHot, "\t", sectionNameRes[index][1])
			fmt.Println("第", secIndex, "区的第", index, "名主播数据储存完毕")
		}

		time.Sleep(time.Second * 3)
		fmt.Println("===============================")
	}

	defer douYuDb.Close()

}

/**
 * @Description 初始化数据库
 * @Param
 * @return
 **/

func InitDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/Mybatis?charset=utf8mb4&parseTime=True&loc=Local"
	douYuDb, err = sql.Open("mysql", dsn)
	fmt.Println("1111111111")
	HandleError(err, "database")
	fmt.Println("2222222222")
	fmt.Println("3333333333")
	fmt.Println("数据库初始化成功！")
	return err
}

func DealWan(hot string) int {
	wan := 10000.00
	if strings.Contains(hot, "万") {

		temp := strings.Trim(hot, "万")
		//temp是没有万字的 1.4
		float, _ := strconv.ParseFloat(temp, 64)
		result := float * wan

		return int(result)
	}

	strHot, err := strconv.Atoi(hot)
	HandleError(err, " strconv.Atoi(hot)")
	return strHot
}

/*
	输入：分区的URL
	返回：分区的所有在直播主播的房间URL
*/
func GetRoomNumsFromSection(sectionUrls []string) []string {

	var allRoomUrls []string
	douyuWeb := `https://www.douyu.com/`
	//每个分区都要进行的操作
	for index, everySectionUrl := range sectionUrls {

		//对网址进行请求
		response, err := http.Get(everySectionUrl)
		HandleError(err, " http.Get(everySectionUrl)")

		//读出来返回的响应体，但目前是bytes
		readAllBytes, err := ioutil.ReadAll(response.Body)
		HandleError(err, "ioutil.ReadAll(response.Body)")

		//把网页的内容转为字符串
		htmlContent := string(readAllBytes)

		//爬直播间的房间号
		resRoomNums := regexp.MustCompile(`</div></h2></div></div></a><a href="/(.*?)" target="_blank"></a></div></li><li class="layout-Cover-item">`)
		roomNumsSets := resRoomNums.FindAllStringSubmatch(htmlContent, -1)

		for i, num := range roomNumsSets {
			roomUrl := douyuWeb + num[1]
			allRoomUrls = append(allRoomUrls, roomUrl)
			fmt.Println("第", index, "个分区结束啦的第", i, "个房间号储存完毕！")
		}

		randSecond := rand.Intn(10)
		for i := 0; i < randSecond; i++ {
			time.Sleep(1 * time.Second)
		}
		if index == 10 {
			break
		}

		fmt.Println("=========第", index, "个分区结束啦")

	}

	return allRoomUrls

}
