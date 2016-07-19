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

func nodesMemory(client *es.Client) {
	stats, err := client.NodeStats()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("\n")
	for _, node := range stats.Nodes {
		fmt.Printf("  \033[36m%s\033[0m\n\n", node.Name)
		fmt.Printf("       Documents: %s\n", humanize.Commaf(node.Indices.Docs.Count))
		fmt.Printf("          Memory: %s free (%d%%) – %s used (%d%%)\n", size(node.Os.Mem.ActualFreeInBytes), int(node.Os.Mem.FreePercent), size(node.Os.Mem.ActualUsedInBytes), int(node.Os.Mem.UsedPercent))
		fmt.Printf("        JVM Heap: %s committed – %s used\n", size(node.Jvm.Mem.HeapCommittedInBytes), size(node.Jvm.Mem.HeapUsedInBytes))
		fmt.Printf("          JVM GC: %s young – %s survivor – %s old\n", size(node.Jvm.Mem.Pools.Young.UsedInBytes), size(node.Jvm.Mem.Pools.Survivor.UsedInBytes), size(node.Jvm.Mem.Pools.Old.UsedInBytes))
		fmt.Printf("      Field Data: %s (%d evictions)\n", size(node.Indices.Fielddata.MemorySizeInBytes), int(node.Indices.Fielddata.Evictions))
		fmt.Printf("    Filter Cache: %s (%d evictions)\n", size(node.Indices.FilterCache.MemorySizeInBytes), int(node.Indices.FilterCache.Evictions))
		fmt.Printf("     Query Cache: %s (%d evictions)\n", size(node.Indices.QueryCache.MemorySizeInBytes), int(node.Indices.QueryCache.Evictions))
		fmt.Printf("        ID Cache: %s\n", size(node.Indices.IDCache.MemorySizeInBytes))
		fmt.Printf("        Segments: %s (%d)\n", size(node.Indices.Segments.MemoryInBytes), int(node.Indices.Segments.Count))
		fmt.Printf("\n")
	}
}
