package helpers

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/xuri/excelize/v2"
	"io"
	mathrand "math/rand"
	"net/http"
	"reflect"
	"time"
)

func GetQueryParams(c *gin.Context) map[string]string {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]string, len(query))
	for k := range query {
		queryMap[k] = c.Query(k)
	}
	return queryMap
}

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	//ReadAtLeast函数用于从指定的读取器“r”读取到指定的缓冲区“buf”，直到它至少读取了指定的最小字节数
	//rand.Reader 共享的密码用强随机数生成器
	n, err := io.ReadAtLeast(rand.Reader, b, length)

	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}

	return string(b)
}

//CreateQuery 创建Sql
func CreateQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s (", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			name := v.Type().Field(i).Name
			if i == 0 {
				query = fmt.Sprintf("%s%s", query, name)
			} else {
				query = fmt.Sprintf("%s, %s", query, name)
			}

		}
		query = fmt.Sprintf("%s) value (", query)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

//ImportXLSX 导入xlsx
func ImportXLSX(file string) ([][]string, error) {
	startTime := time.Now()
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println(row)
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

	return rows, err
}

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//WebsocketHandler 服务
func WebsocketHandler(c *gin.Context) {
	//升级get请求为webSocket协议
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		// 读取客户端的消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// 把消息打印到标准输出
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// 把消息写回客户端，完成回音
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

//TaskDemo 工作池
func TaskDemo() {
	//接受任务
	taskChannel := make(chan int, 20)
	//处理任务
	resChannel := make(chan int, 20)
	//关闭任务
	closeChannel := make(chan bool, 5)

	go func() {
		for i := 0; i < 10; i++ {
			taskChannel <- i
		}
		close(taskChannel)
	}()

	//处理任务
	for i := 0; i < 5; i++ {
		go Task(taskChannel, resChannel, closeChannel)
	}

	go func() {
		//当接收到5个值,说明5个任务完成了
		for i := 0; i < 5; i++ {
			<-closeChannel
		}
		close(resChannel)
		close(closeChannel)
	}()

	//for 循环channel, 当channel关闭以后会退出循环
	for r := range resChannel {
		fmt.Println("res:", r)
	}
}

func Task(taskChannel chan int, resChannel chan int, closeChannel chan bool) {
	for t := range taskChannel {
		resChannel <- t
	}
	closeChannel <- true
}
