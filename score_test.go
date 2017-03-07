package ssdeep

import "testing"

var h1 = FuzzyHash{
	blockSize:   196608,
	hashString1: "7DSC8olnoL1v/uawvbQD7XlZUFYzYyMb615NktYHF7dREN/JNnQrmhnUPI+/n2Y7",
	hashString2: "3DHoJXv7XOq7Mb2TwYHXREN/3QrmktPt",
}
var h2 = FuzzyHash{
	blockSize:   196608,
	hashString1: "pDSC8olnoL1v/uawvbQD7XlZUFYzYyMb615NktYHF7dREN/JNnQrmhnUPI+/n2Yr",
	hashString2: "5DHoJXv7XOq7Mb2TwYHXREN/3QrmktPd",
}
var h3 = FuzzyHash{
	blockSize:   196607,
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

func TestHashDistanceSame(t *testing.T) {
	d := HashDistance(h1, h1)
	if d != 100 {
		t.Errorf("Invalid edit distance: %d", d)
	}
}

func TestHashDistance(t *testing.T) {
	d := HashDistance(h1, h2)
	if d != 97 {
		t.Errorf("Invalid edit distance: %d", d)
	}
}

func TestHashDistanceNoName(t *testing.T) {
	d := HashDistance(h2, h4)
	if d != 97 {
		t.Errorf("Invalid edit distance: %d", d)
	}
}

func BenchmarkDistance(b *testing.B) {
	var h1 = `7DSC8olnoL1v/uawvbQD7XlZUFYzYyMb615NktYHF7dREN/JNnQrmhnUPI+/n2Y7`
	var h2 = `7DSC8olnoL1v/uawvbQD7XlZUFYzYyMb615NktYHF7dREN/JNnQrmhnUPI+/ngrr`
	for i := 0; i < b.N; i++ {
		distance(h1, h2)
	}
}
