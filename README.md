# [goinside-gallog-cleaner](https://github.com/geeksbaek/goinside-gallog-cleaner/releases)
goinside-gallog-cleaner는 디시인사이드 갤로그 클리너입니다. 

갤로그에서 글과 댓글 정보를 수집한 뒤, 실제로 존재하는 글과 댓글, 갤로그에 남아있는 기록을 모두 삭제합니다.

```
// install
go get -u github.com/geeksbaek/goinside-gallog-cleaner

// usage
goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
```

프로그램의 인자로 -id에 디시인사이드 ID를, -pw에 패스워드를 전달해야 하며, 프로그램 실행 시 다음과 같이 동작합니다.
```
$ goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
2016/08/12 12:50:29 ㅇㅇ 님. 로그인에 성공하였습니다.
2016/08/12 12:50:29 모든 글과 댓글을 불러오는 중입니다. 잠시만 기다려주세요.
2016/08/12 12:50:38 글 0개, 댓글 0개
2016/08/12 12:50:38 불러오기를 완료하였습니다.
2016/08/12 12:50:38 불러오는 데 걸린 시간 : 9.1673425s
2016/08/12 12:50:38 삭제를 시작합니다. 잠시만 기다려주세요.
2016/08/12 12:50:38 삭제가 끝났습니다.
2016/08/12 12:50:38 삭제하는 데 걸린 시간 : 976.7µs
```

Jongyeol Baek <geeksbaek@gmail.com>
