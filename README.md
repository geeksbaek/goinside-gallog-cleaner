# goinside-gallog-cleaner
[![Build Status](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner.svg?branch=master)](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner)

goinside-gallog-cleaner는 [goinside](https://github.com/geeksbaek/goinside) 기반으로 만든 디시인사이드 갤로그 클리너입니다. 

갤로그에서 글과 댓글 정보를 수집한 뒤 실제로 존재하는 글과 댓글, 갤로그에 남아있는 기록을 모두 삭제합니다.

## Install
```
$ go get -u github.com/geeksbaek/goinside-gallog-cleaner
```
`go get` 명령어로 직접 패키지를 인스톨해서 빌드하는 대신,  [여기](https://github.com/geeksbaek/goinside-gallog-cleaner/releases)에서 실행 파일을 직접 다운로드 할 수 있습니다.

## Usage
```
$ goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
```

프로그램의 인자로 -id에 디시인사이드 ID를, -pw에 패스워드를 전달해야 하며, 프로그램 실행 시 다음과 같이 동작합니다.
```
$ goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
ㅇㅇ 님. 로그인에 성공하였습니다.                                                        
모든 글과 댓글을 불러오는 중입니다. 잠시만 기다려주세요.                                 
글 0개, 댓글 1개 불러오기를 완료하였습니다.                                              
불러오는 데 걸린 시간 : 3.1934747s                                                       
삭제를 시작합니다. 잠시만 기다려주세요.                                                  
삭제 중... 1/1                                                                           
삭제가 끝났습니다.                                                                       
삭제하는 데 걸린 시간 : 114.3776ms
```

Jongyeol Baek <geeksbaek@gmail.com>
