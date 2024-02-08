package runtimer

import (
	"fmt"
	"runtime"
)

type MemoryStats struct {
	Allocations      uint64
	TotalAllocations uint64
	System           uint64
	NumGC            uint32
}

func (m *MemoryStats) String() string {
	return fmt.Sprintf("Allocations: %d, TotalAllocations: %d, System: %d, NumGC: %d", m.Allocations, m.TotalAllocations, m.System, m.NumGC)
}

func NewMemoryStats() *MemoryStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &MemoryStats{
		Allocations:      m.Alloc,
		TotalAllocations: m.TotalAlloc,
		System:           m.Sys,
		NumGC:            m.NumGC,
	}
}

func PrintMemoryStats() {
	fmt.Printf("Memory Stats: %s\n", NewMemoryStats())
}
