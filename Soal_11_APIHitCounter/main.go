package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type HitCounter struct {
	mu    sync.Mutex
	count int
}

func (hc *HitCounter) Increment() int {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	hc.count++
	return hc.count
}

func (hc *HitCounter) Get() int {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	return hc.count
}

func (hc *HitCounter) Reset() {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	hc.count = 0
}

func main() {
	counter := &HitCounter{}

	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		hit := counter.Increment()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{
			"hit": hit,
		})
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		counter.Reset()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "counter reset to 0",
		})
	})

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
=== Penjelasan Strategi Concurrency ===

Aplikasi ini menggunakan sync.Mutex untuk menjaga keamanan data counter saat
banyak goroutine (request) datang bersamaan. Setiap kali counter dibaca,
ditambah, atau direset, operasi tersebut dibungkus dengan Lock() dan Unlock()
melalui pattern defer. Hal ini memastikan bahwa hanya satu goroutine yang dapat
mengakses atau memodifikasi data counter pada satu waktu, sehingga mencegah
race condition.

Penggunaan defer pada Unlock() menjamin bahwa mutex selalu dilepas meskipun
terjadi panic di dalam fungsi, sehingga tidak akan terjadi deadlock. Pendekatan
ini ringan dan efisien untuk kasus sederhana seperti hit counter, di mana
operasi baca dan tulis sangat cepat dan tidak memerlukan mekanisme concurrency
yang lebih kompleks seperti channel atau sync.RWMutex.
*/
