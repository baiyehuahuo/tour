# tour
Go 语言编程之旅第一章，命令行程序

## 打开工具之旅
1. flag 解析的参数根据在命令行中的顺序进行解析
2. 只要实现了 flag.Value 接口， 就可以定制化实现 flag 可解析的参数

## 单词格式转换
1. 第三方开源库 cobra 可以用于构建子命令
2. internal 包不会被外部包使用
3. 多写 test 文件
4. go test ./... 可以运行当前目录下的所有 test 文件

## 便捷的时间工具
1. 子命令下仍然可以有子命令
2. 用标准库解析时间类。 比如 time.Parse, time.ParseDuration

## 表到结构体的转换
1. template 包 用于基于模板安全输出内容格式
2. 通过访问 information_schema 数据库中的 COLUMNS 表 获取表中列信息
3. 不太像是给学习者看的 没什么营养