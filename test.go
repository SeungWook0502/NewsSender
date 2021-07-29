package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
 
func main() {
    // GET 호출
    resp, err := http.Get("https://news.naver.com/main/list.naver?mode=LS2D&mid=shm&sid1=100&sid2=264")
    if err != nil {
        panic(err)
    }
 
    defer resp.Body.Close()
 
    // 결과 출력
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", string(data))
}