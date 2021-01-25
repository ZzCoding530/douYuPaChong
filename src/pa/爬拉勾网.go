package pa

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLAGou() {

	webUrl := `https://www.lagou.com/jobs/list_产品助理labelWords=&fromSearch=true&suginput=`
	resp, err := http.Get(webUrl)
	HandleError(err, "http.Get")
	webContent, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(webContent))
	HandleError(err, "ioRead")
	//
	////取每个工作的第二层链接
	//hrefRes := regexp.MustCompile(`<a class="position_link" href=(.*?) target="_blank" data-index=`)
	//hrefSet := hrefRes.FindAllStringSubmatch(string(webContent), -1)
	//
	////取每个工作的名字
	//nameRes := regexp.MustCompile(`<h3 style="max-width: 180px;">(.*?)</h3>`)
	//nameSet := nameRes.FindAllStringSubmatch(string(webContent), -1)
	//
	//for index:=0; index<len(hrefSet);index++  {
	//	fmt.Println(hrefSet[index][1],"\t",nameSet[index][1])
	//}

}
