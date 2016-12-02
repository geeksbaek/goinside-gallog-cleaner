package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/geeksbaek/goinside/gallog"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/cheggaaa/pb.v1"
)

var (
	flagMaxCon = flag.Int("c", 50, "maximum concurrent request")
)

func main() {
	defer func() {
		// press enter to close console.
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}()

	id, pw := credentials()
	s, err := gallog.Login(id, pw)
	if err != nil {
		fmt.Print("\n\n")
		fmt.Println("로그인에 실패하였습니다. :", err)
		return
	}
	defer s.Logout()

	fmt.Print("\n\n")
	fmt.Println(s.Name, "님. 로그인에 성공하였습니다.")

	fmt.Println("글과 댓글을 불러오는 중입니다. 잠시만 기다려주세요.")
	start := time.Now()
	fetchProgressCh := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		i := 1
		for _ = range fetchProgressCh {
			fmt.Printf("\r%d번 째 갤로그 페이지 읽는 중...", i)
			i++
		}
		wg.Done()
	}()
	data := s.FetchAll(*flagMaxCon, fetchProgressCh)

	wg.Wait()
	fmt.Print("\n\n")
	fmt.Printf("글 %v개, 댓글 %v개 불러오기를 완료하였습니다.\n", len(data.As), len(data.Cs))
	fmt.Println("불러오는 데 걸린 시간 :", time.Since(start))

	fmt.Print("\n")
	fmt.Println("삭제를 시작합니다.")

	bar := pb.StartNew(len(data.As) + len(data.Cs))
	bar.ShowSpeed = true
	s.DeleteAll(*flagMaxCon, data, func(i, n int) {
		bar.Increment()
	})
	bar.Finish()
	fmt.Println("삭제가 끝났습니다.")
}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("아이디를 입력하세요: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("비밀번호를 입력하세요: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}
