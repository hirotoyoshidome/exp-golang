package main

import (
  "fmt"
  "math"
  "runtime"
)

// 変数はvarで定義する
var greet = "good morning"
// 初期値を代入しないときは型を明示する
var greeting string
// 定数を定義
const NUM = 2

// exec main
func main() {
  fmt.Println(math.Pi)
  fmt.Printf(hello("Hello\n"))
  greeting = "good night"
  fmt.Println(greet)
  fmt.Println(greeting)
  // 関数の中で定義するときは:=で定義できる
  gre := "aaa"
  fmt.Println(gre)
  fmt.Println(NUM)

  // for
  // ()は不要で定義できる
  for i := 0; i < 5; i++ {
    fmt.Println(i+1)
  }

  // if
  if num := 3; num == NUM {
    fmt.Println("2")
  } else {
    fmt.Println("not 2")
  }

  // defer(遅延実行) cc -> bb -> aaの順で出力
  defer fmt.Println("aa")
  defer fmt.Println("bb")
  fmt.Println("cc")

  // pointer(ポインタ（メモリのアドレス情報のこと）)
  // &でアドレスの取得
  // *でポインタ変数の定義
  // ポインタ変数の前に*で参照できる
  lang := "Go"
  fmt.Println(&lang) // 0xc42000e210
  var lang_p *string
  lang_p = &lang
  fmt.Println(lang_p) // 0xc42000e210
  fmt.Println(*lang_p) // Go

  // インスタンスを生成
  bob := Person{"Bob", 22}
  // 継承しているUserインスタンスを生成
  var user1 User
  user1.name = "Bob"

  // 親のメソッドが呼び出される
  fmt.Println(bob.intro("hello"))
  // 継承した子のメソッドが呼び出される
  fmt.Println(user1.intro("1"))

  // Arrayは配列で他の言語と一緒([]で定義)
  // Slices(スライス)は長さの指定がない配列（他にも配列よりリッチな機能あり）
  // Mapsは連想配列（mapで定義）で他の言語と一緒

  // Interfaceはインターフェースでメソッドが定義されているもの
  // オブジェクト指向のそれと同等の機能を提供
  // type interfaceで定義

  // goroutine(ゴルーチン)は並列実行
  // 関数または、メソッドの呼び出しの前にgoを付与することで、並列実行できる
  // runtime.NumGoroutine()を呼び出すことで現在実行中のgoroutineの数を確認できる

  go hello("aa")
  go hello("bb")
  go hello("cc")
  // 個数がわからなかったので1000000をプラスしただけ
  fmt.Println(runtime.NumGoroutine() + 1000000)

  // channel(チャネル)はmake()を使用することで生成可能
  // チャネルオペレータの<-を利用することで値の送受信が可能
  // 送信 channel<-value
  // 受信 <-channel
  // チャネルは値の交換と同期ができるため、ゴルーチンが予期しない状態にならないようにしている

  // chan stringでstring型のチャネルを生成する(chanにも意味があるためこの文字列で生成すること)
  ch := make(chan string)

  // 値の送信
  go func(){
    ch <- "str"
  }()

  // 値の受信
  msg := <- ch
  fmt.Println(msg)
}

// classの概念がないため、type structでclassライクなものを定義する
// structsの定義
type Person struct {
  name string
  age int
}

// Personのメソッドを定義
func(ele Person)intro(str string) string {
  return str + " I am " + ele.name
}

// 継承は無いため、組み込みで定義する(Personを継承したUser)
type User struct {
  Person
}

// Userのメソッド
func(ele User)intro(str string) string {
  return "User no." + str + " : " + ele.name
}

// define function
func hello(str string) string{
  return str
}

// ↑戻り値の型、引数の型も明示する
