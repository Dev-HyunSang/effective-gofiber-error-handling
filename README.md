# gofiber-error-handling
gofiber/fiber를 이용하여서 효율적으로 오류를 해결하는 방법에 대해서 연구합니다.  
Go언어를 이용하여 코드를 작성할 때 다양한 오류에 마주치게 되는데 이러한 다양한 오류를 효율적으로 해결하는 방법에 대해서 연구합니다.  
추후에 더 좋은 방법을 알게 되거나 좋은 방법을 

## 방법론 및 문제점
- 패키지 내에서 작성된 Error Message 외에 개발자가 사용자가에게 쉽게 이해할 수 있도록 Error Message를 변경함.
    - 하지만 이런 생각이 듭니다. 과연 이런 방법을 통해서 사용자 측에서는 이해하기 쉽지만 과연 코드를 만지는 개발자에게 "오류를 쉽게 이해할 수 있는가?"라는 질문을 던지게 됩니다.  
    - 또한 이러한 오류 해결 방법을 통해서 코드량이 많아지고 이로 인해서 코드 복잡도가 상습하여 추후 리펙토링 또는 오류 해결 시 복잡해 질 수 있다는 단점이 있습니다.

```golang
func RunHelloWorld() error {
    ErrFailedHelloWorld := errors.New("HelloWorld 실행이 되지 않습니다.")
    err := HelloWorld()
    if err != nil {
        return ErrFailedHelloWorld
    }
}
```

## 레퍼런스
- [Error는 검사만 하지말고, 우아하게 처리하세요.](http://cloudrain21.com/golang-graceful-error-handling)
- [Golang으로 백엔드 개발하기 - 5. Error Handling. 에러 잘 처리하기 (feat. fiber)](https://umi0410.github.io/blog/golang/how-to-backend-in-go-errorhandle/)