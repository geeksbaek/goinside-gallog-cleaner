# Change Log

## [1.1.2](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.1.2)

app_id 문제를 임시로 해결하였습니다.

## [1.1.1](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.1.1)

디시인사이드 삭제 방식 변경에 대응하여 갤로그 로그가 삭제되지 않은 문제를 해결하였습니다.

## [1.1.0](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.1.0)

FetchAll 함수도 진행 상황을 알 수 있도록 하였습니다.

프로그레스바를 추가하였습니다.

사용성 개선을 위해 아이디와 패스워드를 flag로 전달하는 대신 표준 입력으로 전달하도록 하였습니다.

기본 동시 요청 횟수를 10에서 50으로 증가시켰습니다.

## [1.0.6](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.6)

동시 삭제요청 횟수를 조절할 수 있도록 하였습니다. 이 값은 동시에 몇 번의 삭제 요청을 보낼 것인지 의미합니다. 현재까지 내부적으로 이 값을 100으로 두었으나, 상황에 따라 변경할 수 있도록 프로그램의 인자로 전달받게 하였습니다. 인자로 값을 전달하지 않을 경우 기본 값은 10 입니다.

아래와 같이 -c 옵션으로 값을 전달할 수 있습니다.
```
$ goinside-gallog-cleaner -id ID -pw PW -c 100
```

디시인사이드 서버가 요청을 300회 이상 거부하더라도 프로그램이 종료되지 않도록 하였습니다. 다만, 이 경우에 해당 게시물이 정상적으로 삭제되지 않았을 수 있습니다.

## [1.0.5](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.5)

안정성을 향상시켰습니다.

디시인사이드 서버가 같은 요청을 300회 이상 거부하면 프로그램을 종료하도록 하였습니다.

## [1.0.4](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.4)

실패 로그를 더 자세하게 나타내도록 하였습니다.

## [1.0.3](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.3)

작성한 게시물이 매우 많은 사용자의 글과 댓글을 불러오는 중에 발생하는 문제를 개선하였습니다.

## [1.0.2](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.2)

일부 사용자가 로그인이 실패하는 문제를 goinside 패키지에서 수정하였습니다. 자세한 사항은 [업데이트](https://github.com/geeksbaek/goinside/blob/master/README.md#2016-08-28)를 확인하세요.

## [1.0.1](https://github.com/geeksbaek/goinside-gallog-cleaner/releases/tag/1.0.1)

삭제 상황을 표시하도록 하였습니다.

삭제 속도를 개선하였습니다.
