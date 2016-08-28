# goinside-gallog-cleaner
[![Build Status](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner.svg?branch=master)](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner)

goinside-gallog-cleaner는 [goinside](https://github.com/geeksbaek/goinside) 기반으로 만든 디시인사이드 갤로그 클리너입니다. 

갤로그에서 글과 댓글 정보를 수집한 뒤 실제로 존재하는 글과 댓글, 갤로그에 남아있는 기록을 모두 삭제합니다.

## Install
```
$ go get -u github.com/geeksbaek/goinside-gallog-cleaner
```
`go get` 명령어로 직접 패키지를 인스톨해서 빌드하는 대신,  [여기](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)에서 실행 파일을 직접 다운로드 할 수 있습니다. 직접 다운로드하는 경우 [설명서](http://imgur.com/a/Ei1ok)를 참고하세요.

## Usage
```
$ goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
```

프로그램의 인자로 -id에 디시인사이드 ID를, -pw에 패스워드를 전달해야 하며, 프로그램 실행 시 다음과 같이 동작합니다.
```
$ goinside-gallog-cleaner.exe -id YOUR_ID -pw YOUR_PASSWORD
ㅇㅇ 님. 로그인에 성공하였습니다.
모든 글과 댓글을 불러오는 중입니다. 잠시만 기다려주세요.
글 1894개, 댓글 0개 불러오기를 완료하였습니다.
불러오는 데 걸린 시간 : 6.5880896s
삭제를 시작합니다. 잠시만 기다려주세요.
삭제 중... 1894/1894
삭제가 끝났습니다.
삭제하는 데 걸린 시간 : 3m1.4018302s
```

## Q&A
####Q. 어떻게 사용하나요?
A. 이 프로그램은 명령행 프로그램입니다. 더블 클릭으로 실행하는 것이 아니라 명령행에서 실행해야 합니다. 자세한 설명은 [설명서](http://imgur.com/a/Ei1ok)를 참고하세요.

####Q. 작동하나요?
A. 8월 28일 현재 작동합니다. 작동하지 않는 경우 [연락처](#contact)를 통해 문의해주세요.

####Q. 갤로그 로그만 지워지나요? 실제 게시물만 지워지나요?
A. 갤로그 로그와 실제 게시물이 모두 지워집니다. 다만 갤로그 로그를 기반으로 삭제하기 때문에 수동으로 로그를 삭제한 게시물의 경우 삭제되지 않습니다.

####Q. 마이너 갤러리 게시물도 지워지나요?
A. 마이너 갤러리에 작성한 게시물은 현재 갤로그에 나타나지 않습니다. 갤로그 기반으로 삭제를 하기 때문에 마이너 갤러리에 작성한 게시물은 삭제되지 않습니다.

####Q. 갤로그를 공개로 전환해야 삭제되나요?
A. 공개든 비공개든 상관 없습니다.

####Q. 삭제하는데 얼마나 걸리나요?
A. 테스트하는 시점에서는 약 2000개의 게시물을 삭제하는데 약 3분 가량이 소요되었습니다.

####Q. 악성 코드가 숨겨져 있거나 제 정보를 수집하는 것은 아닌가요?
A. 아닙니다. 이 프로그램은 누구나 코드를 들여다 볼 수 있는 오픈소스 프로그램입니다. 직접 코드를 보시면 확인할 수 있습니다. 

####Q. 제 컴퓨터에서도 프로그램이 실행될까요?
A. 실행됩니다. 이 프로그램은 윈도우, 리눅스, 맥OS를 포함한 대부분의 OS를 지원합니다. [여기](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)에 미리 컴파일된 실행 파일들이 업로드되어 있습니다.

## Contact
기타 질문이나 버그 신고는 Issues에 글을 남기시거나 [카카오톡](https://open.kakao.com/o/s3tYb7m)을 통해 하실 수 있습니다.

## Update

### [1.0.1](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.1)
삭제 상황을 표시하도록 하였습니다.

삭제 속도를 개선하였습니다.

Jongyeol Baek <geeksbaek@gmail.com>
