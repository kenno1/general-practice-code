package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-ini/ini"
	"golang.org/x/sync/semaphore"
)

func init() {
	fmt.Println("Init!")
}

func HelloWorld() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
}

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	f, t bool    = true, false
)

const Pi = 3.14
const (
	usename  = "test_user"
	password = "test_password"
)

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func getOsName() string {
	return "Mac"
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func thirdPartyConnectDB() {
	panic("Unable to connect database")
}

func save() {
	defer func() {
		ss := recover()
		fmt.Println(ss)
	}()
	thirdPartyConnectDB()
}

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	//(*v).X = 1000
	v.X = 1000
}

type Vertexx struct {
	X, Y int
}

func (v Vertexx) Area() int {
	return v.X * v.Y
}

func (v *Vertexx) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

type Vertexx3D struct {
	Vertexx
	z int
}

func (v Vertexx3D) Area3D() int {
	return v.X * v.Y * v.z
}

func (v *Vertexx3D) Scale3D(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertexx3D {
	return &Vertexx3D{Vertexx{x, y}, z}
}
func Area(v Vertexx) int {
	return v.X * v.Y
}

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Println("I don't know %T\n", v)
	}
}

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintln("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func producer(ch chan int, wg *sync.WaitGroup) {
	// Something
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
	fmt.Println("#########################")
}

func producer1(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first chan int, second chan int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second chan int, third chan int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func goroutine11(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine22(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

type PersonMarshal struct {
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:nicknames"`
}

var semaph *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcessSemph(ctx context.Context) {
	isAcquire := semaph.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	defer semaph.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}
func main() {
	LoggingSettings("test.log")

	b := []byte(`{"name":"mike", "age":20, "nicknames": ["a","b","c"]}`)
	var p1 PersonMarshal
	if err := json.Unmarshal(b, &p1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1.Name, p1.Age, p1.Nicknames)

	v01, _ := json.Marshal(p1)
	fmt.Println(string(v01))

	ch1 := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch1)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch1:
			fmt.Println("success")
			break CTXLOOP
		}
	}

	fmt.Println(KB, MB, GB)

	t1 := time.Now()
	fmt.Println(t1.Format(time.RFC3339))

	r11 := regexp.MustCompile("a([a-z]+)e")
	ms1 := r11.MatchString("apple")
	fmt.Println(ms1)

	r22 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r22.FindString("/view/test")
	fss := r22.FindStringSubmatch("/view/test")
	fmt.Println(fs, fss[0], fss[1], fss[2])

	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("Key"))

	c11 := make(chan string)
	c22 := make(chan string)
	go goroutine11(c11)
	go goroutine22(c22)

OuterLoop:
	for {
		select {
		case msg1 := <-c11:
			fmt.Println(msg1)
		case msg2 := <-c22:
			fmt.Println(msg2)
		}
		break OuterLoop
	}

	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer1(first)
	go multi2(first, second)
	go multi4(second, third)
	for result := range third {
		fmt.Println(result)
	}

	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	for ch2 := range ch {
		fmt.Println(ch2)
	}

	sgo := []int{1, 2, 3, 4, 5}
	cgo := make(chan int)
	go goroutine1(sgo, cgo)
	go goroutine2(sgo, cgo)
	xgo := <-cgo
	fmt.Println(xgo)
	ygo := <-cgo
	fmt.Println(ygo)

	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	normal("hello")
	wg.Wait()

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

	do(10)
	do("Mike")
	do(true)

	var mike Human = &Person{"Mike"}
	var x Human = &Person{"x"}
	DriveCar(mike)
	DriveCar(x)

	v10 := New(3, 4, 5)
	v10.Scale3D(10)
	fmt.Println(v10.Area3D())

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	v11 := Vertex{1, 2, "test"}
	changeVertex(v11)
	fmt.Println(v11)

	v12 := &Vertex{1, 2, "test"}
	changeVertex2(v12)
	fmt.Println(v12)

	save()
	fmt.Println("ok?")

	file, err := os.Open("./sample.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}
	fmt.Println(count, string(data))
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}

	HelloWorld()
	fmt.Println(i, f64, s, t, f)
	r1, r2 := add(1, 2)
	fmt.Println(r1, r2)
	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	c2 := circleArea(3)
	fmt.Println(c2(2))

	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xf, xt)

	fmt.Println(string("Hello World!"[0]))
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Contains(s, "test"))

	// var x int = 1
	// xx := float64(x)
	// var y float64 = 1.2
	// xy := int(y)
	// var s string = "14"
	// i, _ := strconv.Atoi(s)
	// fmt.Println(xx, xy, i)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2:4])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{2, 3, 4},
		[]int{5, 6, 7},
	}
	fmt.Println(board)

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m["apple"])
	m["new"] = 500

	// v, ok := m["apple"]
	// fmt.Println(v, ok)

	// v2, ok2 := m["nothing"]
	// fmt.Println(v2, ok2)

	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great")
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	for i, v := range l {
		fmt.Println(i, v)
	}
	for _, v := range l {
		fmt.Println(v)
	}
	ms := map[string]int{"apple": 100, "banana": 200}
	for k, v := range ms {
		fmt.Println(k, v)
	}
	for k := range ms {
		fmt.Println(k)
	}
	for _, v := range ms {
		fmt.Println(v)
	}

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}

	// file, _ := os.Open("./sample.go")
	// defer file.Close()
	// data := make([]byte, 100)
	// fmt.Println(string(data))

	xsx, ysx := 11, 12
	if xsx == 10 && ysx == 10 {
		fmt.Println("&&")
	}

	if xsx == 10 || ysx == 10 {
		fmt.Println("||")
	}
}
