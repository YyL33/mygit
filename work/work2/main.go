package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	//读取文件（方法1）ioutil读取
	byteStr, err := ioutil.ReadFile("F:\\testdata3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))//输出文本文件内容
	//解析编译正则表达式
	reg1 := regexp.MustCompile(`\d\.\d\sus`)
	reg2 := regexp.MustCompile(`\d{5}\sused`)
	//提取需要信息
	cpu := reg1.FindAllStringSubmatch(string(byteStr), -1)
	fmt.Println("cpu:", cpu)
	mem := reg2.FindAllStringSubmatch(string(byteStr), -1)
	fmt.Println("mem:", mem)
	//遍历输出
	fmt.Println("cpu 使用率")
	for _, v := range cpu {
		fmt.Println(v)
	}
	fmt.Println("\n内存使用量")
	for _, v := range mem {
		fmt.Println(v)
	}

	//********************************************************读取文件的其他方法
	/*打开文件
		file,err:=os.Open("F:\\testdata.txt")
		if err!=nil{
			fmt.Println(err)
			return
		}
		defer file.Close()

		//方法2

		//读取文件
		var strSlice []byte
		var n int
		var tempSlice=make([]byte,300)
		for {
			n, err = file.Read(tempSlice)
			if err == io.EOF {
				fmt.Println("读取完毕")
				break
			}
			if err != nil {
				fmt.Println("读取失败")
				return
			}
			strSlice = append(strSlice, tempSlice[:n]...)
		}

	    fmt.Println(string(strSlice))

		//（方法3)bufio读取

		var fileStr string
		reader:=bufio.NewReader(file)
		for{
			str,err:=reader.ReadString('\n')//表示每次读取一行
		if err==io.EOF{
			fileStr+=str
			break
		}
		if err!=nil{
			fmt.Println(err)
			return
		}
		fileStr+=str
		}
		fmt.Println(fileStr)*/
	//************************************************************

}
