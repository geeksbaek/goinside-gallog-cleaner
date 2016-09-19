package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/geeksbaek/goinside/gallog"
)

var (
	flagID     = flag.String("id", "", "id")
	flagPW     = flag.String("pw", "", "password")
	flagMaxCon = flag.Int("c", 10, "maximum concurrent request")
)

func main() {
	flag.Parse()
	s, err := gallog.Login(*flagID, *flagPW)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Logout()
	fmt.Println(s.Name, "님. 로그인에 성공하였습니다.")

	fmt.Println("모든 글과 댓글을 불러오는 중입니다. 잠시만 기다려주세요.")
	start := time.Now()
	data := s.FetchAll(*flagMaxCon)

	fmt.Printf("글 %v개, 댓글 %v개 ", len(data.As), len(data.Cs))
	fmt.Println("불러오기를 완료하였습니다.")
	fmt.Println("불러오는 데 걸린 시간 :", time.Since(start))

	fmt.Println("삭제를 시작합니다. 잠시만 기다려주세요.")
	middle := time.Now()
	s.DeleteAll(*flagMaxCon, data, func(i, n int) {
		fmt.Printf("\r삭제 중... %v/%v", i, n)
	})
	fmt.Println()

	fmt.Println("삭제가 끝났습니다.")
	fmt.Println("삭제하는 데 걸린 시간 :", time.Since(middle))
}
