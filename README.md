# go

- 文字列はバイト配列の集まり。

```go
s := "apple"
fmt.Println(s[0]) // 72
fmt.Println(string(s[0])) // H
```

- Int のビット数は環境依存。64bit の PC で `int` を定義したら `int64` になる。ただし、環境依存の int と明示的に定義した int64 はまったくの別タイプな点に注意
- 文字列はアスキーコードで表現されている
- バイト型

```go
func main() {
    byteA := []byte{72, 73}
    fmt.Println(byteA) // [72, 73]
    fmt.Println(string(byteA)) // HI
    c := []byte("HI")
    fmt.Println(c) // [72, 73]
    fmt.Println(string(c)) // HI
}
```

- 配列型は要素数の変更ができない
- スライスは要素数の変更ができる

```go
// 配列
var arr [3]int
var arr1 [3]int = [3]string{"apple", "orange", "banana"}
arr2 := [...]string{"apple", "orange", "banana"} // 要素数を自動定義する
```

- interface 型( `var x interface{}`)の初期値は `nil` 。全ての型と互換性を持つが、演算には利用できない。
- 型変換

```go
var int a = 10
flt64 := float64(a)
// 文字から数値に変換、packageを使う
var s string = "10"
i, _ = strconv.Atoi(s)
// 数値から文字列に変換
i2 := 200
i, _ = strconv.Itoa(i2)
// 文字列からバイト配列に変換
h := "Hello world"
b := []byte(h)
// バイト配列から文字列に変換
h2 := string(b)
```

- 定数はグローバルに書く `const Pi = 3.14` 定数は大文字で定義すると他のパッケージからも呼び出せるが、小文字で定義するとその定義したパッケージ内でしか呼び出せない。定数は型指定不要。

```go
const a = (
  c0 = iota // iotaは連番を生成する
  c1
  c2
)

fmt.Println(a) // 0, 1, 2
```

## 関数

```go
func Plus(x int, y int) int {
  return x + y;
}

// 同じ型の場合は、型を省略できる。
func Div(x, y int) (int, int) {
  q := x / y
  r := x % y
  return q, r
}

// 返り値は全て使うことが前提だが、_で返り値の破棄できる
x, _ = Div(5, 3)

// 返り値の変数を指定
func Double(x int) (result int) {
  result := x * 2
  return
}

// 返り値なし
func NoReturn() {
  fmt.Println("No return")
}

// 無名関数
f := func(x, y int) int {
  return x + y
}
// 上と同じ
i2 := func(x, y int) int {
  return x + y
}(1, 3)

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
    fmt.Print("This is a function")
  }
}

f := RetrunFunc()
f() // This is a function

// 関数を引数に取る関数

func CallFunc(f func()) {
  f()
}

CallFunc(func() {
  fmt.Print("This is a function")
})

// クロージャー
func Later() func(string) string {
  var store string
  return func(next string) string {
    s := store
    store = next
    return s
  }
}

// ジェネレーター
func integers() func() int {
   i := 0
   return func() int {
     i++
     return i
   }
}

ints := integers()
ints() // 1
ints() // 2
ints() // 3
otherInts := integers()
otherInts() // 1
otherInts() // 2
otherInts() // 3
```

- クロージャーとは日本語では関数閉包と呼ばれ、関数と関数の処理に関する関数外の環境をセットして閉じ込めた物です。
- 通常は関数の実行時にローカルの変数は初期化されるが、クロージャーが絡む場合、参照され続ける限り内部的にクロージャーの変数として定義される。
- ジェネレーター何らかのルールにしたがって、連続した値を返し続ける。go ではクロージャを利用して、ジェネレーターを作れる。

## 制御構文

- if, switch

```go
if a == 2 {
  //　条件式
} else if {
  //　条件式
} else {
　//　条件式
}

// 簡易文つきif
if b := 100, b == 100 {
  //　条件式
}

// エラーハンドリング
var s string := "100"
i, err := strconv.Atoi(s)
if err != nil {
  // 処理
}

// switch
n := 3
switch n {
  case 1, 2:
    //　処理
  case 3:
    // 処理
  default:
    // 処理
}

// 型アサーション
func typeAssertion (value interface{}) {
 switch v := x.(type) {
    case bool:
		  // 処理
    case int:
		  // 処理
    case string:
      // 処理
    case default:
      // 処理
  }
}

// 型によるswitch
// // 異なる型で型アサーションしても、エラーにならない
switch x.(type) {
  case int:
    // 処理
  case string:
    // 処理
  default:
    // 処理
}

switch v := x.(type) {
  case bool:
		// 処理
  case int:
		// 処理
  case string
    // 処理
}
```

- for

```go
point := 0
for point < 10 {
  fmt.Print(10)
  point++
}

for i := 0; i < 10; i++ {
  // 処理
}

arr := [3]string{"a", "b", "c"}
for i := 0; i < len(arr); i++ {
  // 処理
}

// iはindex, vは値
for i, v := range arr {
  // 処理
}

//　ラベル付きfor
Loop:
  for {
    for {
      fmt.Println("start")
      break Loop
    }
  }
```

- `defer`: 関数の終了時に実行される処理を登録できる。

```go
func Defer() {
  defer fmt.Println("END")
  fmt.Println("START")
}

Defer() // Start, ENDの順に表示される

// 複数の処理を登録
// 無名関数を使うことで複数の処理を登録できる。最後に()をつけて実行する
defer func() {
  fmt.Println()
  fmt.Println()
  fmt.Println()
}()

// 複数のdeferを登録した場合、関数の終了時に後から登録された順に表示される。
func　RunDefer() {
  defer fmt.Println(1)
  defer fmt.Println(2)
  defer fmt.Println(3)
}()

RunDefer() // 3,2,1の順で表示される

// ex)
file, err := os.Create("test.txt")
if err != nil {
  fmt.Println(err)
}

defer file.Close()
file.Write([]byte("Hello"))
```

- panic&recover(エラーハンドリング)、プログラムを強制的に止める強力なツールのためあまり使わないようにする。
- `goroutine`(並行処理)はスレッドより小さい処理単位

```go
func sub() {
  for {
    fmt.Println("Sub loop")
    time.Sleep(100 * time.Millisecond)
  }
}

func main() {
  go sub()

  for {
    fmt.Println("Main loop")
    time.Sleep(200 * time.Millisecond)
  }
}
```

- パッケージの初期化をする `init`

```go
// mainより先に実行される
// 複数のinitを定義することもできるが、メリットは少ない
func init() {
  fmt.Print("init")
}

func main() {
  fmt.Print("main")
}
```

## 参照方

- スライス

```go
// 宣言
var s string[]
var s2 []int := []int{10, 20}
s3 := []string{"A", "B"}
// スライスを作成する関数 make(データ型[], 要素数)
s4 := make([]int, 5)

// 値の代入
s[0] = 1000
//　値の取り出し
fmt.Println(s[0]) // 1000
ex := []string{"apple", "orange", "grape", "banana"}
fmt.Println(ex[1:2]) //orange, grape
fmt.Println(ex[:2]) // apple, oranage
fmt.Println(ex[2]) // grape, banana
fmt.Println(ex[:]) // apple, orange, grape, banana
fmt.Println(len(ex)-1) // banana
fmt.Println(ex[1: len(ex)-1]) // orange, grape

/***********************************************/

ex2 := []int{1,2,3,4,5}
// append: 要素の追加
ex2 = append(ex2, 6)
ex2 = append(ex2, 7,8,9)

// make: スライスを生成
// 第二引数は要素数、第二引数は要領
ex3 := make([]int, 3, 10)

// len: 要素数を調べる
len(ex2) // 9

// cap: プログラムのメモリを気にする場合は考慮する。
// 過剰にメモリを確保してしまうと実行速度が落ちたりする。
// 要領以上の要素が追加されるとメモリの消費が倍になってしまいます。
// 良質なパフォーマンスを実現するには、要領の管理も気にする。
cap(ex2) // 9

```

- slice の copy

```go
// copy(コピー先, コピー元)
sl := []int{1,2,3}
sl2 := make([]int, 5, 10)
n := copy(sl2, sl)
fmt.Println(n, sl2) // 5(コピーに成功した数), [1,2,3,4,5]
```

- slice の for

```go
sl := []string{"A", "B", "C"}

for _, v := range sl {
  fmt.Println(v)
}

// 可変長引数
func Sum(n ...int) int {
  val := 0
  for _, v := range n {
    val += n
  }
  return val
}

Sum(1,2,3)
```

## Map

- 存在しない要素を取得しようとした場合、0 バリューを取得する

```go
var m = map[string]int{"A": 100, "B": 200}
m2 := map[string]int{"A": 100, "B": 200}

// 空のマップ
m3 := make(map[int]string)
m3[0] := "A"
m3[1] := "B"

// エラーハンドリング
val, ok := m3[0]
if !ok {
  fmt.Println("error")
}

// 要素数の取得
len(m3)

// delete
delete(m3, 2) // 第二引数はキー

// for

k, v := range m3 {
  fmt.Println(k, v)
}

```

## チャネル

- 複数の `goroutine` 間のデータの受け渡しをするために設計された
- キューの性質を持つデータ構造
- バッファサイズを超えた場合、デッドロックになる。

```go
// 宣言(双方向)
var ch1 chan int

// 受信専用
// var ch2 <-chan int

// 送信専用
// var ch3 ->chan int

ch1 = make(chan int)
ch2 := make(chan int)
ch3 := make(chan int, 5) // 第二引数でバッファの指定

// 送信
ch3 <- 1
ch3 <- 2
ch3 <- 3
len(ch3) // 3

// チャネルからデータを受信
// 受信するたびにチャネルの要素は取り出されるためなくなる
// キューの性質を持つデータ構造
// バッファサイズを超えた場合、デッドロックになる。
i := <-ch3 // 最初に入れたデータ1
len(ch3) // 2
i2 := <-ch3 // 2番目に入れた2
len(ch3) // 1
fmt.Println(<-ch3) // 3
len(ch3) // 0
```

## チャネルとゴルーチン

```go
func reciever(ch chan int) {
  for {
    i := <-c
    fmt.Println(i)
  }
}

func main () {
  ch1 := make(chan int)
  ch2 := make(chan int)

  // チャネルにデータが入るのを待っている
  go reciever(ch1)
  go reciever(ch2)

　// チャネルに送信
  for i < 100 {
    ch1 <- i
    ch2 <- i
    time.Sleep(50 * time.Milisecond)
    i ++
  }

}
```

## チャネルのクローズ

- 送受信を終えたチャネルを明示的にクローズできる
- 閉じたチャネルにはデータの送信はできないが、受信はできる

```go
func receiver(name string, ch chan int) {
  for {
    _, ok := <-ch
    if !ok {
      break
    }
  }
  fmt.Println("END")
}

func main () {
  ch1 := make(chan int, 2)

  close(ch1)

　ch1 <- 1 // できない

  i, ok := <-ch1 // 第二引数はチャネルの状態を返す
  fmt.Println(i, ok)

  go receiver("chan1", ch1)
  go receiver("chan2", ch1)
  go receiver("chan3", ch1)
  i := 0
  for i < 100 {
    ch1 <- i
    i++
  }
  close(ch1)
  time.Sleep(3 * time.Second)
}
```

## チャネルの for

- 送信後、チャネルをクローズしてから使う

```go
ch1 := make(chan int, 3)
ch1 <- 1
ch1 <- 2
ch1 <- 3
close(ch1)
for i := range ch1 {
  fmt.Println(i)
}
```

## チャネルの select

- 複数のチャネルに対する送受信を、ゴルーチンを停止させることなく制御することができる
- select 内はチャネルによる処理以外はエラーになる
- select 内の case の実行順序はランダム

```go
ch1 := make(chan int, 2)
ch2 := make(chan string, 2)

ch2 <- "A"

select {
  case v1 := <-ch1:
    fmt.Println(v1 + 1000)
  case v2 := <-ch2
    fmt.Println(v2 + "!?")
  default:
    fmt.Println("どちらでもない")
}

/**********************************/

func main () {

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
	  for {
	    i := <-ch3
	    ch4 <- i * 2
	  }
	}()

	go func() {
	  for {
	    i2 := <-ch4
	    ch5 <- i2 - 2
	  }
	}()

	n := 0
L:
	for {
	  select {
	    case ch3 <- n:
	      n++
	    case i3 := <- ch5:
	      fmt.Pritnln("received", i3)
	    default:
	      if n > 100 {
	        break L
	      }
	  }
	}
}

```

## ポインタ

- ポインタとは、値型に分類されるデータ構造(基本型、参照型、構造体）のメモリ上のアドレスと型の情報。
- 参照型(スライス、マップ、チャネル)は参照渡しの機能を持っているため、ポインタ型を考慮する必要はない。

```go
func Double (n int) {
  n := n * 2
}

func DoubleV2(i *int) {
  *i = *i * 2
}

func main () {
	var n int
	n := 100
	fmt.Println(n) // 100
  Double(n)
  fmt.Println(n) // 100
	fmt.Println(&n) // メモリアドレスの参照

	// ポインタ型(参照渡し)
	var p *int = &n
	fmt.Println(p) // メモリアドレスを表示
	fmt.Println(*p) // 100(実態を表示)

	*p = 300
	fmt.Println(n) // 300
	n = 200
	fmt.Println(*p) // 200
	DoubleV2(&n)
	fmt.Println(n) // 400
}
```

## 構造体(Struct)

- 複数の任意の型をまとめたもの
- Struct 名もフィールド名も大文字で定義するのが慣例
- フィールドの順番に定義すれば、初期化される
- Struct は値方のため、更新には pointer を使う

```go
type User struct {
  Name string
  Age int
}

func UpdateUser(u *User) {
  u.Name := "Someone"
  u.Age := 50
}

func main() {
  tom := User{Name: "tom", Age: 20}
  fmt.Println(tom.Name)
  fmt.Println(tom.Age)

  // 上書き
  tom.Nmae = "Tom"
  // 順番に定義しないとエラーになる
  bob := {"Bob", 25}
  // 推奨されているの下記
  user := &User{} // { , 0}
  UpdateUser(user)
  fmt.Println(user) // {Name: "Someone", Age: 50}

  user2 := new(User)
  UpdateUser(user)
  fmt.Println(user) // {Name: "Someone", Age: 50}
}
```

- メソッドは任意の型に特化した関数を定義する仕組み。構造体専用の関数。
- 基本的にポインタ型でレシーバーを渡す

```go
func Set(p *Point, i int) {
  p.A = i
}

func (p *Point) Set(i int) {
  p.A = i
}

func main() {
  p1 := &Point{A: 1}
  Set(p1, 2)
  fmt.Println(p1) // 2

  p1.Set(3)
  fmt.Println(p1) // 3
}
```

- 構造体の埋め込み

```go
type Point struct {
  A int
  B string
  C float64
}

type BigPoint {
  Point
}

func main() {
  bp := BigPoint{}
  fmt.Println(bp)
  bp.Point.A = 100
  bp.Point.B = "apple"
  bp.Point.C = 10.5

  fmt.Println(bp.A) // 100

  bp2 := BigPoint {
    Point: Point {
			A: 100,
      B: "apple",
    },
  }
  fmt.Println(bp2) // {{100 apple}}

  bp2.Point.Set(2000)
  fmt.Println(bp2) // {{ 2000 apple}}
}
```

- struct コンストラクタ

```go
type Point struct {
  A int
  B string
  C float64
}

// 型のコンストラクタは」一般的にNewを先頭につける
func NewPoint(a int, b string, c float64) *Point {
  return &Point{a, b, c}
}

func main() {
  p1 := Point{1, "A", 1.1}
  p2 := NewPoint(2, "B", 2.2)
}
```

- struct 構造体とスライス(このパターンは頻出)

```go
type Person struct {
  Name string
}

type Persons struct {
  Persons []*Person
}

func main() {
  ps := make([]Person, 5)
  fmt.Println(ps) // [{} {} {} {} {}]
  ps[0].Name = "Mike"
  ps[1].Name = "Jon"
  ps[2].Name = "Tom"
  ps[3].Name = "Nina"
  ps[4].Name = "Bob"

  fmt.Println(ps) // [{Mike} {Jon} {Tom} {Nina} {Bob}]

  p1 := Person{"George"}
  p2 := Person{"Rikky"}
  p3 := Person{"Mina"}
  p4 := Person{"Nancy"}
  p5 := Person{"Taro"}

  ps := Persons{}
  ps.Persons = append(ps.Persons. &p1, &p2, &p3, &p4, &p5)
  for _, p := range ps.Persons {
    fmt.Println(p)
  }
}
```

## Interface

```go
type Stringfy interface {
  ToString() string
}

type Person {
  Name string
  Age int
}

func (p *Person) Tostring() string {
  return fmt.Printf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car {
  Number string
  Model string
}

func (c *Car) Tostring() string {
  return fmt.Printf("Number=%v, Modele=%v", c.Number, c.Model)
}

func main() {
  vs := []Stringfy{
      &Person{Name: "Taro", Age: 10},
      &Car{Number: "123-456", Model: "AB-123"}

  for _, v := range vs {
    fmt.Println(v.ToString())
  }
}

/*************************/
// カスタムエラー
type MyError struct {
  Message string
  ErrCode int
}

func (e *Error) Error() string {
  return e.Message
}

func RaiseError() error {
  return &MyError{Message: "カスタムエラー", ErrCode* 123}
}
```

## スコープ(パブリックとプライベート、パッケージ分割)

- 変数・定数・関数を大文字で定義すると public になる
- 変数・定数・関数を小文字にすると private になる
- パッケージ名を省略できるが、省略しないの推奨されている
- パッケージはアルファベット順に並べる
- 関数内で定義された変数は、関数ないでのみ参照可能

## Test

- `パッケージ名_test.go` で定義する
- 関数名の先頭に Test をつける
- 引数に `

```go
// main.go

func IsOne(i int) bool {
  if == 1 {
    return true
  } else {
    return false
  }
}

func main() {
  IsOne(1)
}

/***********************/
// main_test.go

 var Degug bool = false

func TestIsOne(t *testing.T) {
  i := 1
  if Degug {
    t.Skip("スキップさせる")
  }
  v := IsOne(i)

  if !v {
    t.Errorf("%v != %v", i, 1)
  }
}
```

## 標準パッケージ

`os` : ファイル処理とうに使う

`time` : 時間に関するパッケージ

- `time.Now()` : 現在時刻を取得
- `time.Date(2020, 6, 10)` : 日付作成
- `time.Now().Year()` :年
- `time.Now().YearDay()` :その年の何日目
- `time.Now().Month()` :月
- `time.Now().WeekDay()`:曜日
- `time.Now().Day()`:日付
- `time.Now().Hour()`:時
- `time.Now().Minute()`:分
- `time.Now().Second()`:秒
- `time.Now().Nanosecond()`:ナノセカンド
- `time.Now().Zone()`:タイムゾーン

時刻の感覚を表現する `time.Duration` 型

- `time.Hour`
- `time.Second`

時間の比較

- `time.Now().Sub(比較時刻)`
- `time.Now().Before(比較対象)`
- `time.Now().After(比較対象)`

指定時間のスリープ

- `time.Sleep(5)`

`math` : 数学に関数パッケージ

- `math.Max(1, 2)`
- `math.Min(1, 2)`
- `math.Trunc`:　数値の正負を問わず、小数点以下を単純に切り捨てる
- `math.Floor` : 引数の数値より小さい最大の整数を返す
- `math.Ceil` :引数の数値より大きい最少の整数を返す

`rand`: 擬似乱数を生成するパッケージ

- `rand.Seed()`
- `rand.Float64()`
- `rand.Seed(time.Now().UnixNano())` : 現在の時刻をシードに使った擬似乱数の生成

`flag` : コマンドラインのオプション処理

`fmt` : フォーマット関連のパッケージ。（[関連](https://www.udemy.com/course/golang-webgosql/learn/lecture/23533228#overview) ）

`log` : ログの出力(エラーハンドリングで使われる)

- `log.SetOutPut(出力先)`
- `log.Print()`
- `log.Fatal()` : プログラムの終了を含む
- `log.Panic()` : プログラムを強制終了する。
- `log.New(出力先, プリフィックス, ログのフォーマット)` : ロガーの生成

`strconv` : 文字列に変換するパッケージ

`strings` : 文字列の処理をするパッケージ

- `strings.Join()`
- `strings.Index()`
- `strings.LastIndex()`
- `strings.LastIndexAny()`
- `strings.HasPrefix()`
- `strings.HasSuffix()`
- `strings.Contains()`
- `strings.ContainsAny()`
- `strings.Count()`
- `strings.Repeat()`
- `strings.Replace()`
- `strings.Split()`
- `strings.SplitAfter()`
- `strings.SplitN()`
- `strings.SplitAfterN()`
- `strings.ToLower()`
- `strings.ToUpper()`
- `strings.TrimSpace()`
- `strings.Fields()`

`bufio` : 入出力処理にバッファ機能を追加したパッケージ

`ioutil` : 簡易的な入出力を管理するパッケージ

`regexp` : 正規表現による文字列処理のパッケージ

- `regexp.MatchString()` : 大量の処理には向いていない
- `regexp.Compile()`
- `regexp.MustCompile`

```go
func main() {

    re1 := regexp.MustCompile(`\s+`)

    //正規表現による文字列の置換
    //正規表現にマッチした部分を別の文字列に置き換える。
    //regexp.Regexp型メソッドReplaceAllString
    //対象の文字列に正規表現のパターンにマッチする部分がない場合は、元の文字列がそのまま返される。
    //スペースを,に置き換える
    fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ","))

    //また置換する文字列に空文字を指定することで、正規表現にマッチした部分を文字列から取り除ける
    re1 = regexp.MustCompile(`、|。`)
    fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", ""))
}

func main() {
    re1 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
    //正規表現による文字列の分割
    //正規表現にマッチした部分で文字列を分割する場合は、regexp.Regexp型のメソッドSplitを使う
    //第二引数　分割する最大数を指定。-1でマッチした全ての箇所で分割する。[]stringで返す。
    fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))

    re1 = regexp.MustCompile(`\s+`)
    //スペースやタブなどの空白にマッチ。空白で分割する。
    fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1))
}
```

`sync` : GO の非同期処理における排他作業や同期処理をサポートする処理をまとめたパッケージ

```go
// ミューテックスによる同期処理
var st struct {A, B, C int}
// ミューテックスを保持するパッケージ
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
  // ロック
　　mutex.Lock()
   st.A = n
   time.Sleep(time.Microsecond)
   st.B = n
   time.Sleep(time.Microsecond)
   st.B = n
   time.Sleep(time.Microsecond)

  // アンロック
   mutex.Unlock()
}

func main() {
  mutex = new(sync.Mutex)

  for i := 0; i < 5; i++ {
    go func() {
      for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
      }
    }
  }
  for {

  }
}

/*******************************/
// 任意のゴルーチンの処理の完了を待ち受ける

func main () {
  /* sync.WaitGroupを生成 */
  wg := new(sync.WaitGroup)
  /* 待ち受けるゴルーチンの数を指定 */
  wg.Add(3)

  go func() {
    for i := 0; i < 1000; i++ {
		  fmt.Println(i)
    }
    wg.Done() // 完了
  }()

  go func() {
    for i := 0; i < 1000; i++ {
		  fmt.Println(i)
    }
    wg.Done() // 完了
  }()

  go func() {
    for i := 0; i < 1000; i++ {
		  fmt.Println(i)
    }
    wg.Done() // 完了
  }()

  // ゴルーチンの完了を待ち受ける
  // ここではDoneが3つ完了するまで待つ
  wg.Wait()

}
```

`cript` : ハッシュ値の生成を補助するパッケージ

```go
package main

import (
    "fmt"
    "io"

    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
)

func main() {
    //------------------------------
    /* SHA-1 */
    s1 := sha1.New()
    io.WriteString(s1, "ABCDE")
    fmt.Printf("%x\n", s1.Sum(nil))
    // => "7be07aaf460d593a323d0db33da05b64bfdcb3a5"

    //------------------------------

    /* SHA-256 */
    s256 := sha256.New()
    io.WriteString(s256, "ABCDE")
    fmt.Printf("%x\n", s256.Sum(nil))
    // => "f0393febe8baaa55e32f7be2a7cc180bf34e52137d99e056c817a9c07b8f239a"

    //------------------------------

    /* SHA-512 */
    s512 := sha512.New()
    io.WriteString(s512, "ABCDE")
    fmt.Printf("%x\n", s512.Sum(nil))
    // => "9989a8fcbc29044b5883a0a36c146fe7415b1439e995b4d806ea0af7da9ca4390eb92@<dtp>{lb}a604b3ecfa3d75f9911c768fbe2aecc59eff1e48dcaeca1957bdde01dfb"
}
```

`json` : エンコーダーとデコーダーの機能を提供するパッケージ

```go
// 構造体からJSONテキストへの変換
type A struct {}

type User struct {
  // `json:""`は小文字に変換している
  ID       int        `json:"id",string` // stringはjson変後後に文字列に変更
  Name     string     `json:"name",omitempty` // 初期値を隠したい場合にomitemptyを追記
  Email    string     `json:"email"`
  Created  time.Time  `json:"created"`
  A        A          `json:"A"`
}

func main() {
  u := new(User)
  u.ID = 1
  u.Name = "test"
  u.Email = "example@exaple.com"
  u.Created = time.Now()

  // Marshal JSONに変換
  bs, err := json.Marshal(u)
  if err != {
    log.Fatal(err)
  }

  fmt.Println(string(bs))

  u2 := new(User)
　
  // Unmarshal JSONを構造体に変換
	if err := json.Unmarshal(bs, &u2); err != nil {
    fmt.Println(err)
  }

  fmt.Println(u2)
}
```

`sort` : slice などをソートする機能を提供

```go
func main() {
  i := []int{5, 3, 8, 0, 2, 10}
  s := []string{"a", "z", "p"}

  sort.Ints(i)
  sort.Strings(s)

  // 昇順・Sliceは安定ソートを保証しない
  sort.Slice(el, func(i, j int) bool {return el[i].Name < el[j].Name})
  // 昇順・SliceTableは安定ソートを保証する
  sort.SliceTable(el, func(i, j int) bool {return el[i].Name < el[j].Name})
}
```

`context`: API のサーバーやクライアントを使用する際に、コンテキストを提供してキャンセルやタイムアウトする仕組み。

```go
ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
defer cancel()

```

`net/url` : url 文字列を処理するパッケージ

```go
u, _ := url.Parse("http://example.com/search?a=1")
u.Schema
u.Host
u.Path
u.RawQuery
u.Fragment

// 生成
url := &url.URL{}
url.Schema = "https:"
u.Host = "google.com"
q := url.Query()
q.Set("q", "Go言語")
u.RawQuery = q.Encode()

/*************************/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Get

func main() {
	//応用
	//ヘッダーをつけたり、クエリをつけたり
	//Parse 正しいURLか確認
	base, _ := url.Parse("https://example.com/")

	//クエリ の確認 URLの後につく
	reference, _ := url.Parse("index/lists?id=1")

	//ResolveReference
	//クエリをくっつけたURLを生成する。
	//相対パスから絶対URLに変換する。
	//baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	//GET ver
	//リクエストを作成 nil部はPOST時のみ設定（バイトを入れる）
	//まだリクエストはしていない。structを作っただけ。
	req, _ := http.NewRequest("GET", endpoint, nil)

	//requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)

	//URLのクエリを確認
	q := req.URL.Query()

	//クエリを追加
	q.Add("name", "test")

	//クエリを表示
	fmt.Println(q)

	//&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode())

	//encodeしてからURLに戻す
	//日本語などを変換する
	req.URL.RawQuery = q.Encode()

	//実際にアクセスする
	//クライアントを作る
	var client *http.Client = &http.Client{}

	//結果 アクセスする
	resp, _ := client.Do(req)

	//読み込み
	body, _ := ioutil.ReadAll(resp.Body)

	//出力
	fmt.Println(string(body))

}

/*****************************/

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

//応用
type Account struct {
    ID       string
    PassWord string
}

//Post

func main() {
    //応用
    //ヘッダーをつけたり、クエリをつけたり
    //Parse 正しいURLか確認
    base, _ := url.Parse("https://example.com/")

    //クエリ の確認 URLの後につく
    reference, _ := url.Parse("index/lists?id=1")

    //Postの時のデータ
    AccountData := &Account{ID: "ABC-DEF", PassWord: "testtest"}
    data, _ := json.Marshal(AccountData)

    //ResolveReference
    //クエリをくっつけたURLを生成する。
    //相対パスから絶対URLに変換する。
    //baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
    endpoint := base.ResolveReference(reference).String()
    fmt.Println(endpoint)

    //POST ver
    //bytes.NewBuffer([]byte("password")でリクエストの領域を作成
    //POSTの場合は、Bodyにデータを入れる。例えばパスワード。見られたらダメな情報はbodyに
    req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))

    //requestにheaderをつける。cash情報など
    req.Header.Add("Content-Type", `application/json"`)

    //URLのクエリを確認
    q := req.URL.Query()

    //クエリを追加
    q.Add("name", "test")

    //クエリを表示
    fmt.Println(q)

    //&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
    fmt.Println(q.Encode())

    //encodeしてからURLに戻す
    //日本語などを変換する
    req.URL.RawQuery = q.Encode()

    //実際にアクセスする
    //クライアントを作る
    var client *http.Client = &http.Client{}

    //結果 アクセスする
    resp, _ := client.Do(req)

    //読み込み
    body, _ := ioutil.ReadAll(resp.Body)

    //出力
    fmt.Println(string(body))

}
```

`net/http` : http クライアントとサーバーがまとめられてたパッケージ

```go
// http client
// GET

res, _ := http.Get("https://exmpale.com")
res.StatusCode
res.Proto
res.Header["Date"]
res.Header["Content-Type"]

res.Request.Method
res.Request.URL

defer res.Body.Close()
body, _ := ioutil.ReadAll(res.Body)

// POST
vs := url.Values{}
vs.Add("id", 1)
vs.Add("message", "Hi")
res, err := http.PostFomr("https://example.com/", vs)
if err != nil {
  log.Fatal(err)
}

defer res.Body.Close()
body, _ := ioutil.ReadAll(res.Body)

/******************************/
// net/http server

type MyHandler struct{}

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *htto,Request) {
//   fmt.Fprintf(w, "Hello world")
// }

func top(w http.ResponseWriter, r *htto,Request) {
  t, err := template.ParseFiles("temp.html")
  if err != nil {
		log.Println(err)
  }
  t.Execute(w, "Hello world!")
}

func main() {
  http.HandleFunc("/top", top)
  http.ListenAndServer(":8080", nil)
}
```

## サードパーティ

`go-ini` : ini ファイルを読み込むためのパッケージ、コンフィグなどの設定情報を読み込む。

```go
// config.ini
[web]
port = 8000

[db]
name = webapp.sql
driver = splite3

```
