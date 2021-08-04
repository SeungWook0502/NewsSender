package main
 
import (
    "fmt"
    "golang.org/x/text/encoding/korean"
    "golang.org/x/text/transform"
)
 
func main() {
    src := "아름다운 우리말"
    exp := "\xbe\xc6\xb8\xa7\xb4\xd9\xbf\xee \xbf\xec\xb8\xae\xb8\xbb"

    got, n, err := transform.String(korean.EUCKR.NewEncoder(), src)
    if err != nil {
        panic(err)
    }

    if got != exp {
        panic("no match")
    }

    fmt.Println([]byte(got), n) // 22 <= 3(UTF-8) * 7(chars) + 1(a space char)

}