# 測試Dependency Injection 依賴注入
使用google wire工具, 實現golang Dependency Injection

# 1. wire_gen 產生後因同在main package可以直接使用initializeBaz()依賴注入好的容器

# 2. 測試用go run . , 目錄底下之go file一同執行

# 3. GOBINPATH 怪怪的, 暫時用/Users/{user_name}/go/bin/wire

# 常用名詞解釋, 以防忘記

# IoC => Inversion Of Control（控制翻轉）, 主要為定義好介面, 之後各自實作方法, 程式載入要使用之介面都會有相對應之方法可調用

# 未使用ioc, 有耦合之寫法

```
type send struct {}

func (*send) putMsg(msg string) {
	fmt.Printf("%s\n", msg)
}

func main() {
	sendMsg := new(send)
	sendMsg.putMsg("hihi")
}
```

# 替換輸出msg的方法時, 主程式依賴class
```
type send struct {}

func (*send) putMsg(msg string) {
	fmt.Printf("%s\n", msg)
}

type newSend struct {}

func (*newSend) putMsg(msg string) {
	fmt.Printf("%s\n", msg)
}

func main() {
	// sendMsg := new(send)
    sendMsg := new(newSend)
	sendMsg.putMsg("hihi")
}
```

# 運用ioc優化, 可抽換功能

```
type msgSpec interface {
	putMsg()
}

type send struct{}

func (*send) putMsg(msg string) {
	fmt.Println(msg)
}

type newSend struct{}

func (*newSend) putMsg(msg string) {
	fmt.Printf("%s\n", msg)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "1" {
		sendMsg := new(send)
		sendMsg.putMsg("send:hihi")
	} else {
		sendMsg := new(newSend)
		sendMsg.putMsg("newSend:hihi")
	}
}
```

# DI => Dependency Injection（相依性注入）

# DI 可用Constructor和Property的這兩種方式注入, 當我有一個package可注入物件時, package就可以依注入之物件執行相對應之功能

# DI Container

# DI我們可以用手動的方式達到，但是其實還是不彈性。當我要切換實作的時候，還是需要改code並且重新編譯

# DI Container會記錄我們所有的interface和實作的對應（可以用xml方式設定，因此不需要重新compile程式就可以改 interface的對應實作），並且管理這些實作的instance scope
