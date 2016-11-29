# [goinside-gallog-cleaner v1.1.1](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)
[![Build Status](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner.svg?branch=master)](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner)
[![Github All Releases](https://img.shields.io/github/downloads/geeksbaek/goinside-gallog-cleaner/total.svg)]()
[![Github Releases](https://img.shields.io/github/downloads/geeksbaek/goinside-gallog-cleaner/latest/total.svg)]()

goinside-gallog-cleaner는 [goinside](https://github.com/geeksbaek/goinside) 기반으로 만든 디시인사이드 갤로그 클리너입니다. 

갤로그에서 글과 댓글 정보를 수집한 뒤 실제로 존재하는 글과 댓글, 갤로그에 남아있는 기록을 모두 삭제합니다.

![](https://raw.githubusercontent.com/geeksbaek/goinside-gallog-cleaner/cf54f8b42e0100dd157f6d15f46ba71433ed9435/guide.gif)

## Latest Update
[2016-11-19 v1.1.1] 디시인사이드 삭제 방식 변경에 대응하여 갤로그 로그가 삭제되지 않은 문제를 해결하였습니다.

## Q&A
#### Q. 어떻게 사용하나요?
> 2016년 10월 30일 (1.1.0 버전) 부터 사용성이 개선되었습니다. [**여기**](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)에서 다운로드받은 파일을 실행시키기만 하면 됩니다.

#### Q. 작동하나요?
> 11월 19일 현재 작동합니다. 이 프로그램은 지속적으로 관리되고 있는 프로젝트입니다. 작동하지 않는 경우 [**연락처**](#contact)를 통해 문의해주세요.

#### Q. 갤로그 로그만 지워지나요? 실제 게시물만 지워지나요?
> 갤로그 로그와 실제 게시물이 모두 지워집니다.
>
> <sub>다만 갤로그 로그를 기반으로 삭제하기 때문에 수동으로 로그를 삭제한 게시물의 경우 삭제되지 않습니다.</sub>

#### Q. 마이너 갤러리 게시물도 지워지나요?
> 지워지지 않습니다. 
>
> <sub>마이너 갤러리에 작성한 게시물은 현재 갤로그에 나타나지 않습니다. 갤로그 기반으로 삭제를 하기 때문에 마이너 갤러리에 작성한 게시물은 삭제되지 않습니다.</sub>

#### Q. 갤로그를 공개로 전환해야 삭제되나요?
> 공개든 비공개든 상관 없습니다.

#### Q. 삭제하는데 얼마나 걸리나요?
> 테스트하는 시점에서는 약 2000개의 게시물을 삭제하는데 약 3분 가량이 소요되었습니다.

#### Q. 악성 코드가 숨겨져 있거나 제 정보를 수집하는 것은 아닌가요?
> 아닙니다. 
>
> <sub>이 프로그램은 누구나 코드를 들여다 볼 수 있는 오픈소스 프로그램입니다. 직접 코드를 보시면 확인할 수 있습니다.</sub> 

#### Q. 제 컴퓨터에서도 프로그램이 실행될까요?
> 실행됩니다.
>
> <sub>이 프로그램은 윈도우, 리눅스, OSX를 포함한 대부분의 운영체제를 지원합니다. [**여기**](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)에 미리 컴파일된 실행 파일들이 업로드되어 있습니다.</sub>

#### Q. 어떤 파일을 다운받아야하죠?
> 자신의 운영체제에 맞는 파일을 하나만 다운받으시면 됩니다. 386은 32bit, amd64는 64비트 운영체제를 의미합니다.
>
> <sub>OSX 사용자는 darwin 버전을 다운받으시면 됩니다.</sub>
>
> <sub>간혹 다운로드 링크가 깨지는 경우, 링크에 포함된 mirror 주소에서 다운받으시기 바랍니다.</sub>

#### Q. 1.1.0 버전 이후로 비밀번호 입력이 안돼요.
> 비밀번호가 보이지 않을 뿐, 입력이 되는 상태입니다.

## Contact
기타 질문이나 버그 신고는 Issues에 글을 남기시거나 [**카카오톡**](https://open.kakao.com/o/s3tYb7m)을 통해 하실 수 있습니다.

Jongyeol Baek <geeksbaek@gmail.com>
