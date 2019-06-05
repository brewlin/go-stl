package test
import(
	"testing"
)

func BenchmarkByteToStringDefault(b *testing.B){
	b.ReportAllocs()
	bs := []byte("hello world")
	for i := 0 ; i < b.N ; i ++ {
		res := string(bs)
		if res == ""{

		}

	}
}

func BenchmarkByteToStringPointer(b *testing.B){
	b.ReportAllocs()
	bs := []byte("hello world")
	for i := 0 ; i < b.N ; i ++ {
		res := Bytes2string(bs)
		if res == ""{
			
		}

	}
}