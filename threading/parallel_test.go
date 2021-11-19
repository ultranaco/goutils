package threading_test

import (
	"log"
	"testing"
	"time"

	"github.com/ultranaco/goutils/threading"
)

func TestParallelRun(t *testing.T) {

	parallel := threading.Parallel{}
	parallel.New(5)

	for iterator := 0; iterator < 15; iterator++ {
		parallel.Run(func(i interface{}) {

			time.Sleep(1 * time.Second)
			log.Println(time.Now(), i)

		}, iterator)
	}

	parallel.Wait()
	log.Println("finish")

}
