package main
 
import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    "github.com/PuerkitoBio/goquery"
)

// go mod init이 선행되어야
// go get을 사용할 수 있음.

func main() {
    response, err := http.Get("https://news.naver.com/main/list.naver?mode=LS2D&mid=shm&sid1=100&sid2=264")
    if err != nil {
        fmt.Println("1")
        panic(err)
    }
    defer response.Body.Close()

    if response.StatusCode != 200 {
        fmt.Println("2")
        panic(response.StatusCode)
    }
    // 결과 출력
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    data2, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
    if err != nil{
        panic(err)
    }

    data2.Find(".list_body dt a").Each(func(i int, s *goquery.Selection) {
        titleText := s.Text()
        fmt.Println(strings.Replace(titleText,"\n","",-1))
        value, _ := s.Attr("href")
        fmt.Println(value)
    })


    // // fmt.Printf("%s\n", string(data))

    // document, err := goquery.NewDocumentFromReader(response.Body)
    // if err != nil {
    //     fmt.Println("3")
    //     panic(err)
    // }

}