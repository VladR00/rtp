package calculation

import (
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

func random() float64 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng := 1.0 + (random.Float64() * (10000.0 - 1.0))
	return math.Round(rng*10) / 10
}

func (storage *Storage) Calculation(count int) float64 {
	var sequence []float64
	var multiplier []float64
	var transformed []float64
	var sum float64

	for i := 0; i < count; i++ {
		sequence = append(sequence, random())
	}
	for i := 0; i < count; i++ {
		multiplier = append(multiplier, random())
	}
	for seq, multi := range multiplier {
		if multi > sequence[seq] {
			transformed = append(transformed, sequence[seq])
			sum += sequence[seq]
		} else {
			transformed = append(transformed, 0)
		}
	}
	log.Println("sequence", sequence)
	log.Println("multiplier", multiplier)
	log.Println("transformed", transformed)
	log.Println("SUM: ", sum)
	rtp := sum / float64(count)
	rtp = math.Round(rtp*10) / 10
	log.Println(rtp)
	return rtp
}
