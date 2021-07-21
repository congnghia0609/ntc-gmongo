//Author: nghiatc
//Since: Jul 22, 2021

package mid

import (
	"fmt"
	"github.com/congnghia0609/ntc-gmongo/gmongo"
	"sync"
	"testing"
)

// Command run test:
// cd ~/go-projects/src/ntc-gmongo
// go test ./mid

func TestBenchmarkMId(t *testing.T) {
	// Init mongo
	gmongo.InitMongo()
	defer gmongo.MClose()

	N := 100 // n thread
	M := 100 // m number id
	name := "benchmark" // name id

	// Reset MIdGen
	rst1, _ := ResetID(name, 0)
	fmt.Println("rst1:", rst1)

	// Benchmark MIdGen
	wg := new(sync.WaitGroup)
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			for j := 0; j < M; j++ {
				n, err := GetNext(name)
				if err != nil {
					panic(err)
				}
				if n % 1000 == 0 {
					fmt.Printf(".")
				}
			}
			wg.Done()
			fmt.Println("Done!!!")
		}()
	}
	wg.Wait()

	// Check consistence MIdGen
	rs, _ := GetNext(name)
	fmt.Println("rs:", rs)
	if int(rs) != (N*M + 1) {
		t.Errorf("TestBenchmarkMId.GetNext: FAIL")
	}

	// Check reset MIdGen
	rst2, _ := ResetID(name, 0)
	fmt.Println("rst2:", rst2)
	if rst2 != 1 {
		t.Errorf("TestBenchmarkMId.ResetID: FAIL")
	}
}
