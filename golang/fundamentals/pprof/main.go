package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// f, err := os.Create("cpu.prof")
	// if err != nil {
	// 	log.Fatal("could not create CPU profile: ", err)
	// }
	// defer f.Close()

	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	log.Fatal("could not start CPU profile: ", err)
	// }
	//defer pprof.StopCPUProfile()

	expensiveFunc()
	http.ListenAndServe(":8080", nil)

}

//go:noinline
func expensiveFunc() {
	var sum float64
	for i := 0; i < 10_000_000; i++ {
		sum += rand.Float64()
	}
	fmt.Printf("Inside expensiveFunc %f\n", sum)

	anotherExpensiveFunc()
}

//go:noinline
func anotherExpensiveFunc() {
	var sum int
	for i := 0; i < 1_000_000; i++ {
		sum += rand.Intn(10)
	}

	fmt.Printf("Inside anotherExpensiveFucn %d\n", sum)
}

// func main() {
// 	// f, err := os.Create("cpu.prof")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// pprof.StartCPUProfile(f)

// 	for i := 0; i < 1000000; i++ {
// 		_ = fmt.Sprintf("Hello, %d", i)
// 	}

// 	http.ListenAndServe(":12345", nil)
// 	//defer pprof.StopCPUProfile()
// }

// func main() {
// 	iterateLong()
// 	iterateShort()

// 	http.ListenAndServe(":8080", nil)
// }

// func iterateLong() {
// 	iterate(9_000_000_000)
// }

// func iterateShort() {
// 	iterate(1_000_000_000)
// }

// func iterate(iterations int) {
// 	for i := 0; i < iterations; i++ {
// 	}
// }

// func main() {
// 	log.Println("booting on localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func main() {
// 	go func() {
// 		log.Println(http.ListenAndServe("localhost:6060", nil))
// 	}()

// 	fmt.Println("hello world")
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go leakyFunction(&wg)
// 	wg.Wait()
// }

// func leakyFunction(wg *sync.WaitGroup) {

// 	s := make([]string, 3)
// 	for i := 0; i < 10000000; i++ {
// 		s = append(s, "magical pandas")
// 		if (i % 100000) == 0 {
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}

// 	defer wg.Done()
// }

// SimpleClient is a thin statsd client.
// type SimpleClient struct {
// 	c  net.PacketConn
// 	ra *net.UDPAddr
// }

// // NewSimpleClient instantiates a new SimpleClient instance which binds
// // to the provided UDP address.
// func NewSimpleClient(addr string) (*SimpleClient, error) {
// 	c, err := net.ListenPacket("udp", ":0")
// 	if err != nil {
// 		return nil, err
// 	}

// 	ra, err := net.ResolveUDPAddr("udp", addr)
// 	if err != nil {
// 		c.Close()
// 		return nil, err
// 	}

// 	return &SimpleClient{
// 		c:  c,
// 		ra: ra,
// 	}, nil
// }

// // Timing sends a statsd timing call.
// func (sc *SimpleClient) Timing(s string, d time.Duration, sampleRate float64,
// 	tags map[string]string) error {
// 	return sc.send(fmtStatStr(
// 		fmt.Sprintf("%s:%d|ms", s, d/time.Millisecond), tags),
// 	)
// }

// func (sc *SimpleClient) send(s string) error {
// 	_, err := sc.c.(*net.UDPConn).WriteToUDP([]byte(s), sc.ra)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func fmtStatStr(stat string, tags map[string]string) string {
// 	parts := []string{}
// 	for k, v := range tags {
// 		if v != "" {
// 			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
// 		}
// 	}

// 	return fmt.Sprintf("%s|%s", stat, strings.Join(parts, ","))
// }

// func main() {
// 	stats, err := NewSimpleClient("localhost:6060")
// 	if err != nil {
// 		log.Fatal("could not start stas client: ", err)
// 	}

// 	// add handlers to default mux
// 	http.HandleFunc("/ping", pingHandler(stats))

// 	s := &http.Server{
// 		Addr: ":8080",
// 	}

// 	log.Fatal(s.ListenAndServe())
// }

// func pingHandler(s *SimpleClient) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		st := time.Now()
// 		defer func() {
// 			_ = s.Timing("http.ping", time.Since(st), 1.0, nil)
// 		}()

// 		w.WriteHeader(200)
// 	}
// }
