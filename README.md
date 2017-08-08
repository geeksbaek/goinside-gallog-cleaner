# goinside-gallog-cleaner

[![GitHub release](https://img.shields.io/github/release/geeksbaek/goinside-gallog-cleaner.svg)](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)
[![Build Status](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner.svg?branch=master)](https://travis-ci.org/geeksbaek/goinside-gallog-cleaner)
[![Github All Releases](https://img.shields.io/github/downloads/geeksbaek/goinside-gallog-cleaner/total.svg)](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)
[![Github Releases](https://img.shields.io/github/downloads/geeksbaek/goinside-gallog-cleaner/latest/total.svg)](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)
[![Donate](https://img.shields.io/badge/donate-paypal-blue.svg)](https://paypal.me/geeksbaek/5)

goinside-gallog-cleaner는 [goinside](https://github.com/geeksbaek/goinside) 기반으로 만든 디시인사이드 갤로그 클리너입니다.

갤로그에서 글과 댓글 정보를 수집한 뒤 실제로 존재하는 글과 댓글, 갤로그에 남아있는 기록을 모두 삭제합니다.

![](https://github.com/geeksbaek/goinside-gallog-cleaner/blob/master/guide.gif?raw=true)

## Notice

현재 이 프로그램은 동작하지만, 장기적인 사용을 보장하기 어렵습니다.

상위 프로젝트인 goinside에 몇 가지 해결되지 못한 [이슈](https://github.com/geeksbaek/goinside/issues)가 있습니다. 이 프로그램을 위해 기여해주신다면 감사하겠습니다.

## Q&A

Q. 어떻게 사용하나요?

>[**여기**](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)에서 사용 환경에 맞는 실행 파일을 다운로드 받은 뒤 실행시키면 됩니다.

Q. 작동하나요?

>**2017년 8월 5일 현재 작동합니다.** 이 프로그램은 지속적으로 관리되고 있는 프로젝트입니다. 작동하지 않는 경우 [**연락처**](#contact)를 통해 문의해주세요.

Q. 갤로그 로그만 지워지나요? 실제 게시물만 지워지나요?

>**갤로그 로그와 실제 게시물이 모두 지워집니다.** 다만 갤로그 로그를 기반으로 삭제하므로 수동으로 로그를 삭제한 게시물의 경우 삭제되지 않습니다.

Q. 마이너 갤러리 게시물도 지워지나요?

>**지워지지 않습니다.** 마이너 갤러리에 작성한 게시물은 현재 갤로그에 나타나지 않습니다. 갤로그를 기반으로 삭제하기 때문에 마이너 갤러리에 작성한 게시물은 삭제되지 않습니다.

Q. 갤로그를 공개로 전환해야 삭제되나요?

>**아닙니다.** 공개든 비공개든 상관없습니다.

Q. 악성 코드가 숨겨져 있거나 제 정보를 수집하는 것은 아닌가요?

>**아닙니다.** 어떠한 정보도 수집하지 않습니다. 누구나 코드를 들여다볼 수 있는 오픈소스 프로그램이므로, 직접 코드를 보시고 확인할 수 있습니다.

Q. 제 컴퓨터에서도 프로그램이 실행될까요?

>**대부분 실행됩니다.** 이 프로그램은 윈도우, 리눅스, macOS를 포함한 대부분 운영체제를 지원합니다. [여기](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/latest)에 미리 컴파일된 실행 파일들이 업로드되어 있습니다.

Q. 어떤 파일을 다운받아야 하죠?

>자신의 운영체제에 맞는 파일을 하나만 다운받으시면 됩니다. 386은 32-bit, amd64는 64-bit 운영체제를 의미합니다. macOS 사용자는 darwin 버전을 다운받으시면 됩니다.

Q. 비밀번호 입력이 안 돼요

>비밀번호가 타이핑되는 동안 보이지 않을 뿐, 입력되는 상태입니다.

Q. Windows Defender SmartScreen에서 프로그램을 차단합니다

[![](http://i.imgur.com/08TjfVx.png)](http://i.imgur.com/08TjfVx.png)

>'추가 정보' 버튼을 눌러 실행시킬 수 있습니다.

## Contact

기타 질문이나 버그 신고는 Issues에 글을 남기실 수 있습니다.

Jongyeol Baek <geeksbaek@gmail.com>
