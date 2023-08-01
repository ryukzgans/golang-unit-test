package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// harus dimulai dengan "Test" kemudian nama functionnya
// t *testing.T

/*
 Menggagalkan Test
	t.Fail() > Fail output, masih menjalankan eksekusi berikutnya
	t.FailNow() > Fail output, langsung membreak eksekusi berikutnya
	t.Error(message) > Fail output dengan print message, masih menjalankan eksekusi berikutnya
	t.Fatal(message) > Fail output dengan print message, langsung membreak eksekusi berikutnya
*/

// #Testing Manual
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

// #Menggunakan Library testify

// go test -v -run=TestHelloWorldAssert
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ilham")
	assert.Equal(t, "Hello Ilham", result, "Result must be 'Hello Ilham'") // t, expected result, actual result, msg
	fmt.Println("Print ini akan tetap dijalankan")                         // assert akan memanggil Fail(), makanya akan tetap dijalankan
}

// go test -v -run=TestHelloWorldRequire
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ilham")
	require.Equal(t, "Hello Ilham", result, "Result must be 'Hello Ilham'") // t, expected result, actual result, msg
	fmt.Println("Print ini tidak akan dijalankan")                          // require akan memanggil FailNow(), makanya tidak akan dijalankan
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Only can run on MacOs")
	}

	result := HelloWorld("Ilham")
	assert.Equal(t, "Hello Ilham", result, "Result must be 'Hello Ilham'")
}

// Before and After

func TestMain(m *testing.M) {
	// BEFORE
	fmt.Println("--> BEFORE UNIT TEST <--")

	m.Run()

	// AFTER
	fmt.Println("--> AFTER UNIT TEST <--")
}
