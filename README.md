# STUDY WITH HEADFIRST GO

### error
- errors.New("meesage")로 error를 만들어 반환
- 메세지에 포맷이 필요한 경우 fmt.Errorf("%s",message)로 에러 반환
- 함수에 에러가 없는 경우 보통 에러값으로 nil을 반환
- [Example](https://play.golang.org/p/-gm_6NbX5fT)
### pointer
- 포인터 또한 변수이며 **주소값을 저장하는 특정한 변수**라고 생각하면됨, 포인터도 주소를 가지고 있음
- 변수명 앞에 "주소" 연산자인 &를 사용하면 변수의 주소값을 가져올 수 있음 (&myVariable)
- 포인터 타입은 포인터가 가리키고 있는 변수의 타입앞에 '*'를 붙여 나타냄 (*int, *bool 등등)
- [Example](https://play.golang.org/p/GbRGTnYOm1B)

### pakage
- 유사한 기능을 수행한느 코드들의 집합
- - 패키지 이름에는 소문자만 사용
- - 가능하면 한 단어만 사용, **두 단어 사용시에 밑줄이나 대문자 사용하지 않음** (strcov)
- 문서 주석 규칙
- - 주석은 완전한 문장이어야 함
- - 패키지 주석은 "Package"로 시작하고 그 뒤에 패키지 이름이 와야함
- - 함수 주석은 함수의 이름으로 시작해야 함

### 배열
- 배열의 번호를 index라고 
- 배열을 만들 떄 배열에 포함된 모든 값은 배열이 가진 타입의 제로값으로 초기화
- 배열에 벗어난 인덱스에 접근하면 패닉(프로그램이 실행하는 동안 발생하는 에러, 컴파일 에러와 반대)
- fmt.__f함수에 %#v 동사를 사용하면 Go코드에서 보이는 그대로 출력 ([2]string{"Hello", "world"})
- 파일 읽기 기능 (/array/readfile참고)

### 슬라이스
- 슬라이스는 내부 배열(underlying array)을 기반으로 구현
- 내부 배열을 변경하면 슬라이스도 변경됨
- slice연산자 mySlice[s:e)
- append함수를 통해 요소를 추가한후 새로운 slice 반환

### 맵
- 모든 값에 라벨을 붙일 수 있는 구조 맵!(사전처럼 생겼다고해서 딕션어리라고도 불림)
- 선언
- - var myMap map[string]int
- - myMap = make(map[string]int)
- 맵 리터럴을 이용한 선언
- - myMap := map[string]int{"a":1, "b":2, "c":3}
- 할당된 값과 제로 값 구분 방법
- - value,ok := myMap("key") // key에 값이 할당되어있다면 ok는 true or false
- - if문에 활용: if _,ok := myMap("key"); ok {}
- delete함수를 사용해 키/값 쌍 삭제
- - delete(myMap,"key")
- for range 문
- - for key, value := range myMap{}

### 구조체
- 여러 타입으로 구성된 값 sliice나 map은 하나의 타입의 값만 저장할 수 있지만 구조체는 여러 타입의 값들을 한데 묶어 저장할 수 있음
- Go는 매개변수 값으로 인자의 복사본을 사용하는 언어 구조체도 마찬가지로 복사본이 이동됨 
- - 큰 구조체인 경우 인자로 전달 할때 메모리가 많이 사용될 수 있음..
- - 왠만하면 포인터 타입으로 전달하는게 좋음
- 구조체의 필드에 접근하는 . 연산자는 구조체 포인터에서도 사용가능
- - struct 2/main.go 참고
- 외부 패키지에서 이용하려면 첫문자를 대문자로 변경해야함 struct/4참고
- 구조체 리터럴을 이용한 초깃값 할당
- - s := T{f1 : v1, f2 : v2} // 필드값 생략가능 제로값으로 초기화
- 익명구조체 필드와 임베딩
- - 익명 필드는 구조체 정의에서 필드명을 생략할 수 있음 
- - 또한 익명 필드로 선언된 내부 구조체를 외부 구조체안에 임베딩 되었다고 하는데 이렇게되면 내부 구조체의 필드를 마치 외부 구조체의 필드인것처럼 사용할 수 있음
- - struct/6 참고