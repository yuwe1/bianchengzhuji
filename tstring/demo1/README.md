<!--
 * @Author: your name
 * @Date: 2019-12-03 19:49:44
 * @LastEditTime: 2019-12-03 19:53:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebf:\mayun\bianchengzhuji\tstring\demo1\README.md
 -->
 ##  蛮力移位
 1、定义指向该字符串的指针 str
 2、字符串的长度为n
 3、编写函数，功能为每次将我们的字符串的第一个字符移动到最后
 若要使字符串前m个字符移动到后面，则只需调用函数m

 ## 时间、空间复杂度
 假设需要移动m个字符到字符串的尾部，那么总共需要m * n 次的操作，因此时间复杂度为O(mn)，空间复杂读为O(1)

 ## 知识点
 ```go
 //string 转[]byte
b := []byte(str)

//[]byte转string
str = string(b)

//string 转 rune
r := []rune(str)

//rune 转 string
str = string(r)
```
