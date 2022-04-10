
# easy-sqlite3

## purpose
- sqlite3를 활용하여 csv 파일 내의 데이터를 쉽게 가공하고자 하는 유저에게, 자유도는 낮지만 기존 sqlite3 모듈보다 쉽게 활용가능한 모듈을 만든다.
- 한번 만들어두고 두고두고 재사용하려고 만드는 내가 필요해서 만드는 모듈 

## concept 

1. 유저는 struct만 정의하면 된다. 
   - main 함수 내에서는 db init()을 통해 생성된 DB 포인터와 struct만 가지고 함수에 변수로 넘겨서 사용한다. 
   - 유저가 제시한 struct로 table create sql, insert sql, delete sql, union 연산 등은 모듈 내부에서 생성하며 함수로 제공한다. 
   - 테이블명, DB명은 랜덤명으로 생성되고 유저가 임의로 지정할 수 없다.
  
2. 유저는 struct list 또는 csv파일을 통해 데이터를 DB에 넣을 수 있다. 
   - 합집합 교집합 등의 연산을 지원한다. (먼저 같은 struct끼리 되게하고, 다른 struct끼리도 key value를 기반으로 되게만든다.)
   - csv파일을 읽어들이고 연산해서 csv파일로 내보내기 기능을 지원한다.

## dependency
- go-sqlite3 모듈을 기반으로 제작되었다. 
https://pkg.go.dev/github.com/mattn/go-sqlite3
https://github.com/mattn/go-sqlite3