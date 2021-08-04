package main
 
import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    // "reflect"
    "time"
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/PuerkitoBio/goquery"
    "golang.org/x/text/encoding/korean"
    "golang.org/x/text/transform"
)

// go mod init이 선행되어야
// go get을 사용할 수 있음.
type article_sid struct {
    sidnum2 string
    title []string
    url []string
}

func main() {

    sidnum2 := [][]string{ //Sid2 value
            {"264","265","268","266","267","269"}, //정치
            {"259","258","261","771","260","262","310","263"}, //경제
            {"249","250","251","254","252","59b","255","256","276","257"}, //사회
            {"241","239","240","267","238","376","242","243","244","248","245"}, //생활문화
            {"231","232","233","234","322"}, //세계
            {"731","226","227","230","732","283","229","228"}, //IT과학
    }

    article_tot := [6]article_sid{}

    for sid1_loop := 0; sid1_loop < len(sidnum2); sid1_loop++ {
        for sid2_loop := 0; sid2_loop < len(sidnum2[sid1_loop]); sid2_loop++ {

            title, url := get_article(sidnum2[sid1_loop][sid2_loop])
            article_struct := article_sid{sidnum2[sid1_loop][sid2_loop],title,url}
            article_tot[sid1_loop] = article_struct
        }
    }

    // fmt.Println(len(article_tot[0].title))
    // for i:=0;i<20;i++{
    //     fmt.Println(i,"-",len(article_tot[0].title),article_tot[0].title[i])
    //     fmt.Println(i,"-",len(article_tot[0].url),article_tot[0].url[i])

    // }
    set_article(article_tot[0])

}

func set_article(article article_sid){

    db, err := sql.Open("mysql", "root:12341234@tcp(127.0.0.1:3306)/go_article")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    var response string
    article_time := time.Now().Format("2006-01-02 15:04:05")
    text := "INSERT INTO article (article_title. article_url,  article_sidnum, article_day) value (\""+article.title[0]+"\",\""+article.url[0]+"\",\""+article.sidnum2+"\",\""+article_time+"\")"
    
    got, _, err := transform.String(korean.EUCKR.NewEncoder(), article.title[0])
    if err != nil {
        fmt.Println("Encoding euc-kr")
        panic(err)
    }

    fmt.Println(text,got)
    err = db.QueryRow("INSERT INTO article (article_title, article_url,  article_sidnum, article_day) value (\""+got+"\",\""+article.url[0]+"\",\""+article.sidnum2+"\",\""+article_time+"\")").Scan(&response)

    
    if err != nil {
        fmt.Println("query")
        log.Fatal(err)
    }

    fmt.Println(response)

}

func get_article(sidnum2 string)([]string, []string){

    target_url := "https://news.naver.com/main/list.naver?mode=LS2D&mid=shm" +
    "&sid2=" + sidnum2

    response, err := http.Get(target_url)
    if err != nil {
        fmt.Println("http.Get")
        panic(err)
    }
    defer response.Body.Close()

    if response.StatusCode != 200 {
        fmt.Println("Status Code")
        panic(response.StatusCode)
    }
    // 결과 출력
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Read document")
        panic(err)
    }

    data2, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
    if err != nil{
        fmt.Println("Stream document")
        panic(err)
    }

    article_title := make([]string,0,2)
    article_url := make([]string,0,2)

    data2.Find(".list_body dt a").Each(func(i int, s *goquery.Selection) {
        
        replacer := strings.NewReplacer("\t","","\n","","\"","'","···","")

        title := s.Text()
        title = replacer.Replace(title)
        if len(title) != 1 && len(title) != 10 {
            article_title = append(article_title,strings.Replace(title," ","",1))
        }

        url, _ := s.Attr("href")
        article_url_temp := replacer.Replace(url)

        if len(article_url) == 0 {
            article_url = append(article_url,article_url_temp)

        } else {

            if article_url[len(article_url)-1] != article_url_temp{
               article_url = append(article_url,replacer.Replace(article_url_temp))
            }
        }
    })
    
    return article_title, article_url
}