package chars

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"hash"
	"io"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

//GbkToUtf8 Gbk转Utf8
func GbkToUtf8(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

//Utf8ToGbk Utf8转Gbk
func Utf8ToGbk(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

//Md5ToStr 字符串转Md5
func Md5ToStr(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

//Md5Str 字符串转Md5
func Md5Str(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//RandSmsCode 随机短信验证码
func RandSmsCode(num int) string {
	txtChars := "0123456789"
	textNum := len(txtChars)
	text := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < num; i++ {
		text = text + string(txtChars[r.Intn(textNum)])
	}
	return text
}

//RandString 随机字符串
func RandString(n int, alphabets ...byte) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}

//EncodePassword Encode密码
func EncodePassword(salt, password string) string  {
	return hex.EncodeToString(PBKDF2([]byte(password), []byte(salt), 10000, 50, sha256.New))
}

func PBKDF2(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// N.B.: || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}

//SHA1QwJsSignature SHA1QwJs签名
func SHA1QwJsSignature(str string) string {
	t := sha1.New();
	io.WriteString(t,str);
	return fmt.Sprintf("%x",t.Sum(nil));
}

//ArrIntToString 数组转字符串，逗号隔开
func ArrIntToString(arr []int, sep string) (str string){
	for k, d := range arr{
		if k == len(arr)-1 {
			str += strconv.Itoa(d)
		}else{
			str += strconv.Itoa(d) + sep
		}
	}
	return
}

func ArrInt32ToString(arr []int32, sep string) (str string){
	for k, d := range arr{
		if k == len(arr)-1 {
			str += fmt.Sprint(d)
		}else{
			str += fmt.Sprint(d) + sep
		}
	}
	return
}

func ArrInt64ToString(arr []int64, sep string) (str string){
	for k, d := range arr{
		if k == len(arr)-1 {
			str += strconv.FormatInt(d,10)
		}else{
			str += strconv.FormatInt(d,10) + sep
		}
	}
	return
}

func ArrStrToString(arr []string, sep string) (str string) {
	for k, v := range arr {
		if k == len(arr) -1 {
			str += v
		}else{
			str += v + sep
		}
	}
	return
}

//InArray 判断是否在数据
func InArray(s string, arr []string) bool {
	for _,v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

//DistinctArr 切片去重
func DistinctArr(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}





