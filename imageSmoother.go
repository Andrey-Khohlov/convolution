// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var img [][]int = [][]int{{7, 7, 5}, {2, 4, 6}, {8, 2, 0}}
	t1 := time.Now()
	r := imageSmoother(img)
	t2 := time.Now()
	fmt.Println(r)
	fmt.Println("time", t2.Sub(t1))
}

func imageSmoother(img [][]int) [][]int {
	var m, n int = len(img), len(img[0])

	var av = make([][]int, m)
	for i := range av {
		av[i] = make([]int, n)
	}

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			wg.Add(1)
			ii, jj := i, j
			go func() {
				time.Sleep(100 * time.Millisecond)
				var sum, count int
				var ileft, iright, jtop, jbottom int = ii - 1, ii + 1, jj - 1, jj + 1

				if ileft < 0 {
					ileft = 0
				}
				if iright >= len(img) {
					iright = len(img) - 1
				}
				if jtop < 0 {
					jtop = 0
				}
				if jbottom >= len(img[0]) {
					jbottom = len(img[0]) - 1
				}

				for m := ileft; m <= iright; m++ {
					for n := jtop; n <= jbottom; n++ {
						sum += img[m][n]
						count++
					}
				}
				mu.Lock()
				av[ii][jj] = sum / count
				mu.Unlock()
				wg.Done()
			}()
		}

	}
	wg.Wait()
	return av
}
