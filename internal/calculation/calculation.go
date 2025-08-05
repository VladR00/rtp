package calculation

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Storage struct {
	RTP float64
}

func NewStorage(rtp float64) *Storage {
	return &Storage{RTP: rtp}
}

func random(min float64, max float64) float64 {
	if max > 10000 {
		max = 10000
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng := min + (random.Float64() * (max - min))
	return math.Round(rng*10) / 10
}

func generateMultiplier(sum float64, count int, allcount int, rtp float64, seq float64) float64 {
	sumFinal := rtp * float64(allcount)
	iterationforend := float64(allcount) - float64(count)
	sumNeed := sumFinal - sum //+ iterationforend
	fmt.Println("Iteration for end: ", iterationforend)
	fmt.Println("Sum need: ", sumNeed)
	fmt.Println("Current sum: ", sum)
	fmt.Println("Sequence: ", seq)

	if sum <= sumFinal {
		if seq <= sumNeed {
			return random(seq-1, 10000)
		} else {
			return random(1, seq-1)
		}
	} else {
		return random(1, seq-1)
	}
}

func (storage *Storage) Calculation(count int) float64 {
	//fmt.Println("start calc")
	sequence := make([]float64, count)
	multiplier := make([]float64, count)
	var sum, allsum float64
	chAllSum := make(chan float64, count)
	wg := &sync.WaitGroup{}
	fmt.Println("start for")

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := random(1, 10000)
			chAllSum <- val
			sequence[i] = val
			allsum += val
		}()
	}
	wg.Wait()
	close(chAllSum)
	fmt.Println("FOR chan")
	for val := range chAllSum {
		allsum += val
	}

	fmt.Println("Sequense len: ", len(sequence))

	for i := 0; i < count; i++ {
		multiplier[i] = generateMultiplier(sum, i, count, storage.RTP, sequence[i])
		if multiplier[i] > sequence[i] {
			sum += sequence[i]
		}
	}
	fmt.Println("multiplier len: ", len(multiplier))

	rtp := sum / float64(count)
	rtp = math.Round(rtp*10) / 10
	allsum = math.Round(allsum*10) / 10
	sum = math.Round(sum*10) / 10
	// fmt.Println(multiplier)
	fmt.Printf("AllSum: %.1f\n", allsum)
	fmt.Printf("SUM:    %.1f\n", sum)
	fmt.Printf("count: %d\n", count)
	fmt.Printf("countfloat: %f\n", float64(count))
	fmt.Printf("storagertp: %f\n", storage.RTP)
	fmt.Printf("NeedSum: %f\n", storage.RTP*float64(count))
	fmt.Println("RTP:   ", rtp)
	return rtp
}
