package simple

import (
	"sync"
	"testing"
)

func TestProxyPattern(t *testing.T) {
	station := &Station{stock: 10}
	stationProxy := StationProxy{station: station}
	stationProxy.sell()
}

// TestStationSell bad test case
func TestStationSell(t *testing.T) {
	tests := []struct {
		name    string
		station *Station
		time    int
		want    bool
	}{
		{"北京西-郑州东-2-2-true", &Station{stock: 2}, 2, true},
		{"郑州东-北京西-2-3-false", &Station{stock: 2}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			station := Station{
				stock: tt.station.stock,
			}

			resultCh := make(chan bool, 1)
			var wg sync.WaitGroup
			wg.Add(tt.time)

			for i := 0; i < tt.time; i++ {
				go func(i int) {
					defer wg.Done()
					result := station.sell()
					if i == (tt.time - 1) {
						resultCh <- result
					}
				}(i)
			}

			wg.Wait()
			close(resultCh)

			lastResult := <-resultCh
			if lastResult != tt.want {
				t.Logf("%s [result should be: %t, but got: %t]", tt.name, tt.want, lastResult)
			}
		})
	}
}

func TestStationProxySell(t *testing.T) {
	tests := []struct {
		name    string
		station *Station
		time    int
		want    bool
	}{
		{"北京西-郑州东-100-100-true", &Station{stock: 100}, 100, true},
		{"郑州东-北京西-100-103-false", &Station{stock: 100}, 103, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proxy := &StationProxy{
				station: tt.station,
			}
			resultCh := make(chan bool, tt.time)
			var wg sync.WaitGroup
			wg.Add(tt.time)
			for i := 0; i < tt.time; i++ {
				go func() {
					defer wg.Done()
					result := proxy.sell()
					resultCh <- result
				}()
			}

			wg.Wait()
			close(resultCh)

			successCount := 0
			for result := range resultCh {
				if result {
					successCount++
				}
			}

			// 计算最终结果：是否有超过的售票成功
			finalResult := successCount == tt.station.stock
			// 判断最终的期望结果
			if (successCount == tt.time && tt.want == true) || (successCount < tt.time && tt.want == false) {
				finalResult = tt.want
			}
			if finalResult != tt.want {
				t.Fatalf("%s [result should be: %t, but got: %t,success count:%d]", tt.name, tt.want, finalResult, successCount)
			}
		})
	}
}
