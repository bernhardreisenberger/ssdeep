package ssdeep

import (
	"math/rand"
	"os"
	"testing"
)

var h1 = FuzzyHash{
	blockSize:   192,
	hashString1: "MUPMinqP6+wNQ7Q40L/iB3n2rIBrP0GZKF4jsef+0FVQLSwbLbj41iH8nFVYv980",
	hashString2: "x0CllivQiFmt",
}
var h2 = FuzzyHash{
	blockSize:   192,
	hashString1: "JkjRcePWsNVQza3ntZStn5VfsoXMhRD9+xJMinqF6+wNQ7Q40L/i737rPVt",
	hashString2: "JkjlQyIrx+kll2",
}
var h3 = FuzzyHash{
	blockSize:   196608,
	hashString1: "pDSC8olnoL1v/uawvbQD7XlZUFYzYyMb615NktYHF7dREN/JNnQrmhnUPI+/n2Yr",
	hashString2: "5DHoJXv7XOq7Mb2TwYHXREN/3QrmktPd",
}
var h4 = FuzzyHash{
	blockSize:   196608,
	hashString1: "7DSC8olnoL1v/uawvbQD7XlZUFYzYyMb615NktYHF7dREN/JNnQrmhnUPI+/n2Y7",
	hashString2: "3DHoJXv7XOq7Mb2TwYHXREN/3QrmktPt",
}
var h5 = FuzzyHash{
	blockSize:   196608,
	hashString1: "7DSC8olnoL1v/uawvbQD7XlZUFYzYyMb615NktYHF7dREN/JNnQrmhnUPI+/n2Y7",
	hashString2: "",
}

func TestFuzzy1(t *testing.T) {
	f1, err := os.Open("testfile1.txt")
	if err != nil {
		return
	}
	h, err := Fuzzy(f1)
	if h.hashString1 != h1.hashString1 {
		t.Error("testfile1.txt: hashString1 wrong")
	}
	if h.hashString2 != h1.hashString2 {
		t.Error("testfile1.txt: hashString2 wrong")
	}
}

func TestFuzzy2(t *testing.T) {
	f2, err := os.Open("testfile2.txt")
	if err != nil {
		return
	}
	h, err := Fuzzy(f2)
	if h.hashString1 != h2.hashString1 {
		t.Error("testfile2.txt: hashString1 wrong")
	}
	if h.hashString2 != h2.hashString2 {
		t.Error("testfile2.txt: hashString2 wrong")
	}
}

func TestToString(t *testing.T) {
	testHash := h1
	if testHash.String() != "192:MUPMinqP6+wNQ7Q40L/iB3n2rIBrP0GZKF4jsef+0FVQLSwbLbj41iH8nFVYv980:x0CllivQiFmt" {
		t.Error("String method does not work correctly")
	}
}

func TestRollingHash(t *testing.T) {
	sdeep := SSDEEP{
		rollingState: rollingState{
			window: make([]byte, rollingWindow),
		},
	}
	if sdeep.rollHash(byte('A')) != 585 {
		t.Error("Rolling hash not matching")
	}
}

func BenchmarkRollingHash(b *testing.B) {
	sdeep := SSDEEP{
		rollingState: rollingState{
			window: make([]byte, rollingWindow),
		},
	}
	for i := 0; i < b.N; i++ {
		sdeep.rollHash(byte(i))
	}
}

func BenchmarkSumHash(b *testing.B) {
	testHash := hashIinit
	data := []byte("Hereyougojustsomedatatomakeyouhappy")
	for i := 0; i < b.N; i++ {
		testHash = sumHash(data[rand.Intn(len(data))], testHash)
	}
}
