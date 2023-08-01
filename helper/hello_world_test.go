package helper

import (
	"fmt"
	"testing"
)

// harus dimulai dengan "Test" kemudian nama functionnya
// t *testing.T

/*
 Mengagalkan Test
	t.Fail() > Fail output, masih menjalankan eksekusi berikutnya
	t.FailNow() > Fail output, langsung membreak eksekusi berikutnya
	t.Error(message) > Fail output dengan print message, masih menjalankan eksekusi berikutnya
	t.Fatal(message) > Fail output dengan print message, langsung membreak eksekusi berikutnya
*/

func TestHelloWorldIlham(t *testing.T) {
	result := HelloWorld("Ilham")
	if result != "Hello Ilham" {
		// test error
		t.Error("Result must be 'Hello Ilham'")
	}
	fmt.Println("Test")
}

func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("Kurniawan")
	if result != "Hello Kurniawan" {
		// test error
		t.Fatal("Result must be 'Hello Kurniawan'")
	}
	fmt.Println("Test")
}
