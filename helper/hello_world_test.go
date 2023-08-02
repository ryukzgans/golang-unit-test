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

// #Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Only can run on MacOs")
	}

	result := HelloWorld("Ilham")
	assert.Equal(t, "Hello Ilham", result, "Result must be 'Hello Ilham'")
}

// #Before and After
func TestMain(m *testing.M) {
	// BEFORE
	fmt.Println("--> BEFORE UNIT TEST <--")

	m.Run()

	// AFTER
	fmt.Println("--> AFTER UNIT TEST <--")
}

// #Sub Test
func TestSubTest(t *testing.T) {
	// go test -v -run=TestSubTest/IlhamTest
	t.Run("IlhamTest", func(t *testing.T) {
		result := HelloWorld("Ilham")
		require.Equal(t, "Hello Ilham", result, "Result must be 'Hello Ilham'")
	})

	// go test -v -run=TestSubTest/KurniawanTest
	t.Run("KurniawanTest", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'")
	})
}

// #Table Test
// go test -v -run=TestTableHelloWorld
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "IlhamTest",
			request:  "Ilham",
			expected: "Hello Ilham",
		},
		{
			name:     "KurniawanTest",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "TarmijiTest",
			request:  "Tarmiji",
			expected: "Hello Tarmiji",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// #Benchmark

// go test -v -run=TestTidakAda -bench=.
// go test -v -run=TestTidakAda -bench=BenchmarkHelloWorldIlham
func BenchmarkHelloWorldIlham(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ilham")
	}
}
func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}
