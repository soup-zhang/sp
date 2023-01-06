package chars

import "testing"

func TestEncodePassword(t *testing.T) {
	s := EncodePassword("zb#@)8", "123456")
	t.Log(s,len(s))
}

func TestSHA1QwJsSignature(t *testing.T)  {
	s := SHA1QwJsSignature("abcdef")
	t.Log(s,len(s))
}

func TestArrInt32ToString(t *testing.T)  {
	arr := []int32{1,2,3}
	s := ArrInt32ToString(arr, "-")
	t.Log(s)
}

func TestArrStrToString(t *testing.T) {
	arr := []string{"a","b","c"}
	s := ArrStrToString(arr, ":")
	t.Log(s)
}

func TestDistinctArr(t *testing.T)  {
	arr := []string{"a","b","c","a","f","c"}
	arr2 := DistinctArr(arr)
	t.Log(arr2)
}
