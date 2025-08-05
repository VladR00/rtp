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

func generateMultiplier(sum float64, count int, allcount int, rtp float64, check *bool, seq float64) float64 {
	sumFinal := rtp * float64(allcount)
	sumNeed := sumFinal - sum + float64(allcount) - float64(count)
	if sum <= sumNeed {
		fmt.Println("sum: ", sum)
		fmt.Println("sumneed: ", sumNeed)
		return random(1, 150)
	}
	fmt.Println("SumNeed: ", sumNeed)
	if sumNeed < 0 {
		sumNeed *= -1
	}
	if count <= 100 && *check == true && allcount-1 != count {
		val := random(1, 1000)
		if val > seq {
			*check = false
			fmt.Println("Check turn false, i: ", count)
		}
		return val
	}
	fmt.Println(sumNeed)
	if allcount-1 == count && *check == true && allcount <= 1000 {
		val := random(1, 10000)
		fmt.Println("last case, Value:", val)
		return val
	}
	return random(1, sumNeed)
}

func (storage *Storage) Calculation(count int) float64 {
	//fmt.Println("start calc")
	sequence := make([]float64, count)
	multiplier := make([]float64, count)
	var sum, allsum float64
	var firstcheck bool = true
	chAllSum := make(chan float64, count)
	wg := &sync.WaitGroup{}
	var mu sync.Mutex
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			multiplier[i] = generateMultiplier(sum, i, count, storage.RTP, &firstcheck, sequence[i])
			if multiplier[i] > sequence[i] {
				mu.Lock()
				defer mu.Unlock()
				sum += sequence[i]
			}
		}()
	}
	wg.Wait()

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
