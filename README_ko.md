# Port Killer

[English](README.md)

Port Killer는 현재 실행 중인 TCP 포트를 리스닝하는 프로세스를 찾아 손쉽게 종료할 수 있는 간단한 Go CLI 도구입니다.

개발 중 "Address already in use" 에러가 발생했을 때, 어떤 프로세스가 해당 포트를 점유하고 있는지 확인하고 즉시 종료하는 데 유용합니다.

## 기능

- `lsof` 명령어를 사용하여 LISTEN 상태인 TCP 포트 목록 조회
- 사용자 친화적인 테이블 형식으로 프로세스 정보 표시 (Command, PID, User, Address)
- 번호 선택을 통한 대화형 프로세스 종료
- `kill -9` (SIGKILL) 시그널을 사용하여 강제 종료

## 요구 사항

- **OS**: Linux 또는 macOS (Windows는 `lsof` 미지원으로 인해 현재 지원되지 않음)
- **Dependencies**: `lsof`가 설치되어 있어야 합니다.
- **Go**: 소스 코드를 빌드하거나 실행하려면 Go가 필요합니다.

## 설치 및 실행

### 1. 소스 코드 실행

```bash
go run main.go
```

### 2. 빌드하여 실행

```bash
go build -o port-killer main.go
./port-killer
```

전역에서 사용하려면 빌드된 바이너리를 PATH에 포함된 디렉토리로 이동시키세요 (예: `/usr/local/bin`).

## 사용 방법

프로그램을 실행하면 현재 열려있는 포트 목록이 표시됩니다.

```text
 List of currently running ports
No   Command    PID    User     Address
--   -------    ---    ----     -------
1    node       12345  user     *:3000
2    main       67890  user     *:8080

 Select a number to terminate (Cancel: 0):
```

종료하고 싶은 프로세스의 번호를 입력하고 엔터를 누르면 해당 프로세스가 종료됩니다.

> **참고**: 시스템 프로세스나 다른 사용자의 프로세스를 종료하려면 `sudo` 권한이 필요할 수 있습니다.

```bash
sudo ./port-killer
# 또는
sudo go run main.go
```

## 기여

버그를 발견하거나 개선 사항이 있다면 이슈를 등록하거나 PR을 보내주세요.
