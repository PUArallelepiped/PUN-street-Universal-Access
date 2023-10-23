// package strings

// import "math/rand"

// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// // 生成随机字符串，n 代表生成长度
// func Rand(n int) string {
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = letterBytes[rand.Intn(len(letterBytes))]
// 	}
// 	return string(b)
// }
