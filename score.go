package ssdeep

import "math"

// HashDistance between two strings
func HashDistance(hash1, hash2 FuzzyHash) int {
	blockSize1, hash11, hash12 := hash1.blockSize, hash1.hashString1, hash1.hashString2
	blockSize2, hash21, hash22 := hash2.blockSize, hash2.hashString1, hash2.hashString2

	// We can only compare equal or *2 block sizes
	if blockSize1 != blockSize2 && blockSize1 != blockSize2*2 && blockSize2 != blockSize1*2 {
		return 0
	}
	// TODO: remove char repetitions in hashes here as they skew the results
	// Could use some regex to do this: /(.)\1{9,}/
	// TODO: compare char by char to exit fast
	if blockSize1 == blockSize2 && hash11 == hash21 {
		return 100
	}
	var score int
	if blockSize1 == blockSize2 {
		d1 := scoreDistance(hash11, hash21, blockSize1)
		d2 := scoreDistance(hash12, hash22, blockSize1*2)
		score = int(math.Max(float64(d1), float64(d2)))
	} else if blockSize1 == blockSize2*2 {
		score = scoreDistance(hash11, hash22, blockSize1)
	} else {
		score = scoreDistance(hash12, hash21, blockSize2)
	}
	return score
}

func scoreDistance(h1, h2 string, blockSize int) int {
	d := distance(h1, h2)
	d = (d * spamSumLength) / (len(h1) + len(h2))
	d = (100 * d) / spamSumLength
	d = 100 - d
	/* TODO: Figure out this black magic...
	matchSize := float64(blockSize) / float64(blockMin) * math.Min(float64(len(h1)), float64(len(h2)))
	if d > int(matchSize) {
		d = int(matchSize)
	}
	*/
	return d
}
