package calculation

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

type Storage struct {
	RTP float64
}

func NewStorage(rtp float64) *Storage {
	return &Storage{RTP: rtp}
}

func random(min float64, max float64) float64 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng := min + (random.Float64() * (max - min))
	return math.Round(rng*10) / 10
}

func generateMultiplier(sum float64, count int, rtp float64) float64 {
	var newMultiplier float64
	rtpNow := sum / float64(count)
	rtpNow = math.Round(rtpNow*10) / 10
	// log.Println("sum: ", sum)
	// log.Println("rtpNow: ", rtpNow)
	// log.Println("rtpNeed: ", rtp)
	if rtpNow > rtp {
		newMultiplier = 1
	} else {
		newMultiplier = random(1, 10000)
	}
	return newMultiplier
}

func (storage *Storage) Calculation(count int) float64 {
	var sequence []float64
	var multiplier []float64
	var transformed []float64
	var sum float64
	var allsum float64

	start := time.Now()
	log.Println("Start goroutines")
	for i := 0; i < count; i++ {
		val := random(1, 10000)
		sequence = append(sequence, val)
		allsum += val
	}
	log.Println("End goroutines, len seq: ", len(sequence))
	log.Println(time.Since(start))
	for i := 0; i < count; i++ {
		multiplier = append(multiplier, generateMultiplier(sum, i, storage.RTP))
		if multiplier[i] > sequence[i] {
			transformed = append(transformed, sequence[i])
			sum += sequence[i]
		} else {
			transformed = append(transformed, 0)
		}
	}
	// log.Println("sequence", sequence)
	// log.Println("multiplier", multiplier)
	// log.Println("transformed", transformed)
	log.Println("SUM: ", sum)
	allsum = math.Round(allsum*10) / 10
	sum = math.Round(sum*10) / 10
	fmt.Printf("AllSum: %.2f\n", allsum)
	rtp := sum / float64(count)
	rtp = math.Round(rtp*10) / 10
	log.Println(rtp)
	return rtp
}
