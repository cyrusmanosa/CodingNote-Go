package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func init() {
	// 打開文件
	file, err := os.OpenFile("string.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("打開文件失敗:", err)
		return
	}
	defer file.Close()

	// 讀取文件內容
	data := make([]byte, 8192)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("讀取文件失敗:", err)
		return
	}
	fmt.Printf("文件內容：\n%s\n", string(data[:count]))
}

func main() {
	fmt.Println(" ")
	s := "Hello,World!!!!!"

	//Count計算字符串sep在s中的非重疊個數：2
	//func Count(s, sep string) int
	strings.Count(s, "!!")

	//Contains判斷字符串s中是否包含子串substr：true
	//func Contains(s, substr string) bool
	strings.Contains(s, "Hello")

	//ContainsAny判斷字符串s中是否包含chars中的任何一個字符：true
	//func ContainsAny(s, chars string) bool
	strings.ContainsAny(s, "HJK")

	//ContainsRune判斷字符串s中是否包含字符r：true
	//Go語言的單引號一般用來表示「rune literal」，即碼點字面量
	//func ContainsRune(s string, r rune) bool
	strings.ContainsRune(s, 'H')

	//Index返回子串sep在字符串s中第一次出現的位置,如果找不到，則返回-1：6
	//func Index(s, sep string) int
	strings.Index(s, "W")

	//LastIndex返回子串sep在字符串s中最後一次出現的位置,如果找不到，則返回-1：9
	//func LastIndex(s, sep string) int
	strings.LastIndex(s, "l")

	//IndexRune返回字符r在字符串s中第一次出現的位置,如果找不到，則返回-1：6
	//func IndexRune(s string, r rune) int
	strings.IndexRune(s, 'W')

	//IndexAny返回字符串chars中的任何一個字符在字符串s中第一次出現的位置：8
	//如果找不到，則返回-1
	//func IndexAny(s, chars string) int
	strings.IndexAny(s, "rd")

	//LastIndexAny返回字符串chars中的任何一個字符在字符串s中最後一次出現的位置：10
	//如果找不到，則返回-1
	//func LastIndexAny(s, chars string) int
	strings.LastIndexAny(s, "rd")

	//Split以sep為分隔符，將s切分成多個子切片：[Hello World!!!!!]
	//func Split(s, sep string) []string
	strings.Split(s, ",")

	//Join將a中的子串連接成一個字符串，用sep分隔：Monday|Tuesday|Wednesday
	//func Join(a []string, sep string) string
	ss := []string{"Monday", "Tuesday", "Wednesday"}
	strings.Join(ss, "|")

	//HasPrefix判斷字符串s是否以prefix開頭：true
	//func HasPrefix(s, prefix string) bool
	strings.HasPrefix(s, "He")

	//HasSuffix判斷字符串s是否以prefix結尾：true
	//func HasSuffix(s, suffix string) bool
	strings.HasSuffix(s, "!!")

	//Repeat將count個字符串s連接成一個新的字符串
	//func Repeat(s string, count int) string
	strings.Repeat(s, 2)

	//ToUpper將s中的所有字符修改為其大寫格式：HELLO,WORLD!!!!!
	//func ToUpper(s string) string
	strings.ToUpper(s)

	//ToLower將s中的所有字符修改為其小寫格式：hello,world!!!!!
	//func ToLower(s string) string
	strings.ToLower(s)

	//Trim將刪除s首尾連續的包含在cutset中的字符：世界
	//func Trim(s string, cutset string) string
	sss := " Hello 世界! "
	strings.Trim(sss, " Helo!")

	//TrimLeft將刪除s頭部連續的包含在cutset中的字符：世界！
	//func TrimLeft(s string, cutset string) string
	strings.TrimLeft(sss, " Helo")

	//TrimRight將刪除s尾部連續的包含在cutset中的字符： Hello
	//func TrimRight(s string, cutset string) string
	strings.TrimRight(sss, " 世界!")

	//TrimSpace將刪除s首尾連續的的空白字符：Hello 世界!
	//func TrimSpace(s string) string
	strings.TrimSpace(sss)

	//TrimPrefix刪除s頭部的prefix字符串：,World!!!!!
	//func TrimPrefix(s, prefix string) string
	strings.TrimPrefix(s, "Hello")

	//TrimSuffix刪除s尾部的suffix字符串：Hello,World!
	//func TrimSuffix(s, suffix string) string
	strings.TrimSuffix(s, "!!!!")

	//Replace返回s的副本，並將副本中的old字符串替換為new字符串
	//替換次數為 n 次，如果 n 為 -1，則全部替換
	//func Replace(s, old, new string, n int) string
	strings.Replace(s, "!", "|", -1)

	//EqualFold判斷s和t是否相等。忽略大小寫，同時它還會對特殊字符進行轉換：true
	//比如將“ϕ”轉換為“Φ”、將“Ǆ”轉換為“ǅ”等，然後再進行比較
	//func EqualFold(s, t string) bool
	s1 := "Hello 世界! ϕ Ǆ"
	s2 := "hello 世界! Φ ǅ"
	strings.EqualFold(s1, s2)

	//通過字符串s創建strings.Reader對象
	//func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
	r := strings.NewReader(s)

	//Len返回r.i之後的所有數據的字節長度:16
	//func (r *Reader) Len() int
	r.Len()

	//Read將r.i之後的所有數據寫入到b中（如果b的容量足夠大）
	//func (r *Reader) Read(b []byte) (n int, err error)
	b := make([]byte, 5)
	for n, _ := r.Read(b); n > 0; n, _ = r.Read(b) {
		fmt.Println(string(b[:n])) //"Hello", ",Worl", "d!!!!", "!",
	}

	//ReadAt將off之後的所有數據寫入到b中（如果 b 的容量足夠大）
	//func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
	n, _ := r.ReadAt(b, 6)
	fmt.Println(string(b[:n])) //World

	//ReadByte將r.i之後的一個字節寫入到返回值b中
	//func (r *Reader) ReadByte() (b byte, err error)
	r.ReadByte()

	//Seek用來移動r中的索引位置，offset：要移動的偏移量，負數表示反向移動
	//whence：從那裏開始移動，0：起始位置，1：當前位置，2：結尾位置
	//func (r *Reader) Seek(offset int64, whence int) (int64, error)
	r1 := strings.NewReader(s)
	b1 := make([]byte, 5)
	r1.Seek(6, 0) // 移動索引位置到第7個字節
	r1.Read(b1)   // 開始讀取
	fmt.Printf("%q\n", b1)
	r1.Seek(-5, 1) // 將索引位置移回去
	r1.Read(b1)    // 繼續讀取
	fmt.Printf("%q\n", b1)

	//WriteTo 將 r.i 之後的數據寫入接口 w 中
	//func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
	r2 := strings.NewReader(s)
	buf := bytes.NewBuffer(nil)
	r2.WriteTo(buf)
	fmt.Printf("%q\n", buf) //"Hello,World!!!!!"

	//NewReplacer 通過“替換列表”創建一個 Replacer 對象。
	//按照“替換列表”中的順序進行替換，只替換非重疊部分。
	//如果參數的個數不是偶數，則拋出異常。
	//func NewReplacer(oldnew ...string) *Replacer

	//Replace 返回對 s 進行“查找和替換”後的結果
	//Replace 使用的是 Boyer-Moore 算法，速度很快
	//func (r *Replacer) Replace(s string) string
	srp := strings.NewReplacer("Hello", "你好", "World", "世界", "!", "！")
	s = "Hello World!Hello World!hello world!"
	rst := srp.Replace(s)
	fmt.Print(rst) //你好 世界！你好 世界！hello world！
}
