// 유일한 경로로 구성되는 module 선언으로 시작
module ch9

// 최소 호환 버전
go 1.22.1

// 의존하는 모듈, 각 모듈에 필요한 최소 버전
// 어떤 모듈과도 의존하지 않으면 require 생략 가능
require (
	github.com/shopspring/decimal v1.2.0
)

// replace : 의존성 있는 모듈이 있는 위치 재정의
// exclude : 특정 버전의 모듈 사용 막을 수 있다.