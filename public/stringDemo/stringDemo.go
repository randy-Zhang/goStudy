/*
 * @Author: zcw
 * @Date: 2024-03-15 10:12:38
 * @LastEditTime: 2024-03-15 14:37:50
 * @Description: In User Settings Edit
 * @FilePath: \studyDemo\stringDemo\stringDemo.go
 */
package stringdemo

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func StringTest() {

	str := "这是一个测试字符串，专门用来测试滴，请注意"
	fmt.Println(str)

	qz := strings.HasPrefix(str, "这是")
	fmt.Println("前缀：", qz)

	hz := strings.HasSuffix(str, "这是")
	fmt.Println("后缀：", hz)

	bh := strings.Contains(str, "门")
	fmt.Println("包含：", bh)

	th := strings.Replace(str, "测试", "验证", -1)
	fmt.Println("替换：", th)

	tj := strings.Count(str, "测试")
	fmt.Println("统计：", tj)

	cf := strings.Repeat(str, 3)
	fmt.Println("重复：", cf)

	// strings.ToLower()
	// strings.ToUpper()
	fg := strings.Split(str, "，")
	fmt.Println("分割：", fg)

	pjArr := []string{"测试", "字符串", "拼接", "呀"}
	pj := strings.Join(pjArr, "")
	fmt.Println("测试拼接：", pj)

	// 从字符串中读取内容
	r := strings.NewReader(str)

	buf := make([]byte, len(str))
	r.Read(buf)
	fmt.Println("读取byte：", string(buf[:]))

	r1 := strings.NewReader(str)

	fmt.Println("ReadByte读取：")
	for i := 0; i < len(str); i++ {
		if val, err := r1.ReadByte(); err == nil {
			fmt.Printf("%v", val)
		}
	}
	fmt.Println("")
	fmt.Println("ReadRune读取：")
	r2 := strings.NewReader(str)
	for i := 0; i < utf8.RuneCountInString(str); i++ {
		if val, _, err := r2.ReadRune(); err == nil {
			// fmt.Println(string(val))
			fmt.Printf("%v", string(val))
		}
	}
	fmt.Println("")

	builder := strings.Builder{}
	builder.Write([]byte(str))
	builder.Write([]byte("，测试Builder写入，"))
	builder.WriteRune('哈')
	builder.WriteString("，WriteString写入")
	fmt.Println("测试Builder写入：", builder.String())

	atoi, _ := strconv.Atoi("10")
	fmt.Println("Atoi：", atoi)
}
