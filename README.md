# goinside-gallog-cleaner
[![Build Status](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner.svg?branch=master)](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner)
[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg?maxAge=2592000)](https://gitter.im/goinside-gallog-cleaner/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)

goinside-gallog-cleaner는 [goinside](https://github.com/geeksbaek/goinside) 기반으로 만든 디시인사이드 갤로그 클리너입니다. 

갤로그에서 글과 댓글 정보를 수집한 뒤 실제로 존재하는 글과 댓글, 갤로그에 남아있는 기록을 모두 삭제합니다.

## Install
```
$ go get -u github.com/geeksbaek/goinside-gallog-cleaner
```
`go get` 명령어로 직접 패키지를 인스톨해서 빌드하는 대신,  [여기](https://github.com/geeksbaek/goinside-gallog-cleaner/releases)에서 실행 파일을 직접 다운로드 할 수 있습니다. 직접 다운로드하는 경우 [설명서](http://imgur.com/a/Ei1ok)를 참고하세요.

## Usage
```
$ goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
```

프로그램의 인자로 -id에 디시인사이드 ID를, -pw에 패스워드를 전달해야 하며, 프로그램 실행 시 다음과 같이 동작합니다.
```
$ goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
ㅇㅇ 님. 로그인에 성공하였습니다.                                                                       
모든 글과 댓글을 불러오는 중입니다. 잠시만 기다려주세요.                                                
글 356개, 댓글 0개 불러오기를 완료하였습니다.                                                           
불러오는 데 걸린 시간 : 3.1939822s                                                                      
삭제를 시작합니다. 잠시만 기다려주세요.                                                                 
삭제 중... 356/356                                                                                      
삭제가 끝났습니다.                                                                                      
삭제하는 데 걸린 시간 : 35.4421305s 
```

## Update

### [1.0.1](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.1)
삭제 상황을 표시하도록 하였습니다.

삭제 속도를 개선하였습니다.

Jongyeol Baek <geeksbaek@gmail.com>
