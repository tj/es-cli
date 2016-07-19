package main

import (
	"fmt"
	"log"

	"github.com/dustin/go-humanize"
	"github.com/tj/es"
)

func size(n float64) string {
	return humanize.Bytes(uint64(n))
}

func stat(name, format string, vals ...interface{}) {
	fmt.Printf("\033[1m%15s\033[0m: %s\n", name, fmt.Sprintf(format, vals...))
}

func nodesMemory(client *es.Client) {
	stats, err := client.NodeStats()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("\n")
	for _, node := range stats.Nodes {
		fmt.Printf("  \033[36m%s\033[0m\n\n", node.Name)
		stat("Documents", humanize.Commaf(node.Indices.Docs.Count))
		stat("Memory", "%s free (%d%%) – %s used (%d%%)", size(node.Os.Mem.ActualFreeInBytes), int(node.Os.Mem.FreePercent), size(node.Os.Mem.ActualUsedInBytes), int(node.Os.Mem.UsedPercent))
		stat("JVM Heap", "%s committed – %s used", size(node.Jvm.Mem.HeapCommittedInBytes), size(node.Jvm.Mem.HeapUsedInBytes))
		stat("JVM GC", "%s young – %s survivor – %s old", size(node.Jvm.Mem.Pools.Young.UsedInBytes), size(node.Jvm.Mem.Pools.Survivor.UsedInBytes), size(node.Jvm.Mem.Pools.Old.UsedInBytes))
		stat("Field Data", "%s (%d evictions)", size(node.Indices.Fielddata.MemorySizeInBytes), int(node.Indices.Fielddata.Evictions))
		stat("Filter Cache", "%s (%d evictions)", size(node.Indices.FilterCache.MemorySizeInBytes), int(node.Indices.FilterCache.Evictions))
		stat("Query Cache", "%s (%d evictions)", size(node.Indices.QueryCache.MemorySizeInBytes), int(node.Indices.QueryCache.Evictions))
		stat("ID Cache", "%s", size(node.Indices.IDCache.MemorySizeInBytes))
		stat("Segments", "%s (%d)", size(node.Indices.Segments.MemoryInBytes), int(node.Indices.Segments.Count))
		fmt.Printf("\n")
	}
}
