package nodestats

// Stats for all nodes.
type Stats struct {
	ClusterName string           `json:"cluster_name"`
	Nodes       map[string]*Node `json:"nodes"`
}

// Node stats.
type Node struct {
	Breakers struct {
		Fielddata struct {
			EstimatedSize        string  `json:"estimated_size"`
			EstimatedSizeInBytes float64 `json:"estimated_size_in_bytes"`
			LimitSize            string  `json:"limit_size"`
			LimitSizeInBytes     float64 `json:"limit_size_in_bytes"`
			Overhead             float64 `json:"overhead"`
			Tripped              float64 `json:"tripped"`
		} `json:"fielddata"`
		Parent struct {
			EstimatedSize        string  `json:"estimated_size"`
			EstimatedSizeInBytes float64 `json:"estimated_size_in_bytes"`
			LimitSize            string  `json:"limit_size"`
			LimitSizeInBytes     float64 `json:"limit_size_in_bytes"`
			Overhead             float64 `json:"overhead"`
			Tripped              float64 `json:"tripped"`
		} `json:"parent"`
		Request struct {
			EstimatedSize        string  `json:"estimated_size"`
			EstimatedSizeInBytes float64 `json:"estimated_size_in_bytes"`
			LimitSize            string  `json:"limit_size"`
			LimitSizeInBytes     float64 `json:"limit_size_in_bytes"`
			Overhead             float64 `json:"overhead"`
			Tripped              float64 `json:"tripped"`
		} `json:"request"`
	} `json:"breakers"`
	Fs struct {
		Data []struct {
			AvailableInBytes     float64 `json:"available_in_bytes"`
			DiskIoOp             float64 `json:"disk_io_op"`
			DiskIoSizeInBytes    float64 `json:"disk_io_size_in_bytes"`
			DiskQueue            string  `json:"disk_queue"`
			DiskReadSizeInBytes  float64 `json:"disk_read_size_in_bytes"`
			DiskReads            float64 `json:"disk_reads"`
			DiskServiceTime      string  `json:"disk_service_time"`
			DiskWriteSizeInBytes float64 `json:"disk_write_size_in_bytes"`
			DiskWrites           float64 `json:"disk_writes"`
			FreeInBytes          float64 `json:"free_in_bytes"`
			TotalInBytes         float64 `json:"total_in_bytes"`
		} `json:"data"`
		Timestamp float64 `json:"timestamp"`
		Total     struct {
			AvailableInBytes     float64 `json:"available_in_bytes"`
			DiskIoOp             float64 `json:"disk_io_op"`
			DiskIoSizeInBytes    float64 `json:"disk_io_size_in_bytes"`
			DiskQueue            string  `json:"disk_queue"`
			DiskReadSizeInBytes  float64 `json:"disk_read_size_in_bytes"`
			DiskReads            float64 `json:"disk_reads"`
			DiskServiceTime      string  `json:"disk_service_time"`
			DiskWriteSizeInBytes float64 `json:"disk_write_size_in_bytes"`
			DiskWrites           float64 `json:"disk_writes"`
			FreeInBytes          float64 `json:"free_in_bytes"`
			TotalInBytes         float64 `json:"total_in_bytes"`
		} `json:"total"`
	} `json:"fs"`
	Indices struct {
		Completion struct {
			SizeInBytes float64 `json:"size_in_bytes"`
		} `json:"completion"`
		Docs struct {
			Count   float64 `json:"count"`
			Deleted float64 `json:"deleted"`
		} `json:"docs"`
		Fielddata struct {
			Evictions         float64 `json:"evictions"`
			MemorySizeInBytes float64 `json:"memory_size_in_bytes"`
		} `json:"fielddata"`
		FilterCache struct {
			Evictions         float64 `json:"evictions"`
			MemorySizeInBytes float64 `json:"memory_size_in_bytes"`
		} `json:"filter_cache"`
		Flush struct {
			Total             float64 `json:"total"`
			TotalTimeInMillis float64 `json:"total_time_in_millis"`
		} `json:"flush"`
		Get struct {
			Current             float64 `json:"current"`
			ExistsTimeInMillis  float64 `json:"exists_time_in_millis"`
			ExistsTotal         float64 `json:"exists_total"`
			MissingTimeInMillis float64 `json:"missing_time_in_millis"`
			MissingTotal        float64 `json:"missing_total"`
			TimeInMillis        float64 `json:"time_in_millis"`
			Total               float64 `json:"total"`
		} `json:"get"`
		IDCache struct {
			MemorySizeInBytes float64 `json:"memory_size_in_bytes"`
		} `json:"id_cache"`
		Indexing struct {
			DeleteCurrent        float64 `json:"delete_current"`
			DeleteTimeInMillis   float64 `json:"delete_time_in_millis"`
			DeleteTotal          float64 `json:"delete_total"`
			IndexCurrent         float64 `json:"index_current"`
			IndexTimeInMillis    float64 `json:"index_time_in_millis"`
			IndexTotal           float64 `json:"index_total"`
			IsThrottled          bool    `json:"is_throttled"`
			NoopUpdateTotal      float64 `json:"noop_update_total"`
			ThrottleTimeInMillis float64 `json:"throttle_time_in_millis"`
		} `json:"indexing"`
		Merges struct {
			Current            float64 `json:"current"`
			CurrentDocs        float64 `json:"current_docs"`
			CurrentSizeInBytes float64 `json:"current_size_in_bytes"`
			Total              float64 `json:"total"`
			TotalDocs          float64 `json:"total_docs"`
			TotalSizeInBytes   float64 `json:"total_size_in_bytes"`
			TotalTimeInMillis  float64 `json:"total_time_in_millis"`
		} `json:"merges"`
		Percolate struct {
			Current           float64 `json:"current"`
			MemorySize        string  `json:"memory_size"`
			MemorySizeInBytes float64 `json:"memory_size_in_bytes"`
			Queries           float64 `json:"queries"`
			TimeInMillis      float64 `json:"time_in_millis"`
			Total             float64 `json:"total"`
		} `json:"percolate"`
		QueryCache struct {
			Evictions         float64 `json:"evictions"`
			HitCount          float64 `json:"hit_count"`
			MemorySizeInBytes float64 `json:"memory_size_in_bytes"`
			MissCount         float64 `json:"miss_count"`
		} `json:"query_cache"`
		Recovery struct {
			CurrentAsSource      float64 `json:"current_as_source"`
			CurrentAsTarget      float64 `json:"current_as_target"`
			ThrottleTimeInMillis float64 `json:"throttle_time_in_millis"`
		} `json:"recovery"`
		Refresh struct {
			Total             float64 `json:"total"`
			TotalTimeInMillis float64 `json:"total_time_in_millis"`
		} `json:"refresh"`
		Search struct {
			FetchCurrent      float64 `json:"fetch_current"`
			FetchTimeInMillis float64 `json:"fetch_time_in_millis"`
			FetchTotal        float64 `json:"fetch_total"`
			OpenContexts      float64 `json:"open_contexts"`
			QueryCurrent      float64 `json:"query_current"`
			QueryTimeInMillis float64 `json:"query_time_in_millis"`
			QueryTotal        float64 `json:"query_total"`
		} `json:"search"`
		Segments struct {
			Count                       float64 `json:"count"`
			FixedBitSetMemoryInBytes    float64 `json:"fixed_bit_set_memory_in_bytes"`
			IndexWriterMaxMemoryInBytes float64 `json:"index_writer_max_memory_in_bytes"`
			IndexWriterMemoryInBytes    float64 `json:"index_writer_memory_in_bytes"`
			MemoryInBytes               float64 `json:"memory_in_bytes"`
			VersionMapMemoryInBytes     float64 `json:"version_map_memory_in_bytes"`
		} `json:"segments"`
		Store struct {
			SizeInBytes          float64 `json:"size_in_bytes"`
			ThrottleTimeInMillis float64 `json:"throttle_time_in_millis"`
		} `json:"store"`
		Suggest struct {
			Current      float64 `json:"current"`
			TimeInMillis float64 `json:"time_in_millis"`
			Total        float64 `json:"total"`
		} `json:"suggest"`
		Translog struct {
			Operations  float64 `json:"operations"`
			SizeInBytes float64 `json:"size_in_bytes"`
		} `json:"translog"`
		Warmer struct {
			Current           float64 `json:"current"`
			Total             float64 `json:"total"`
			TotalTimeInMillis float64 `json:"total_time_in_millis"`
		} `json:"warmer"`
	} `json:"indices"`
	Jvm struct {
		BufferPools struct {
			Direct struct {
				Count                float64 `json:"count"`
				TotalCapacityInBytes float64 `json:"total_capacity_in_bytes"`
				UsedInBytes          float64 `json:"used_in_bytes"`
			} `json:"direct"`
			Mapped struct {
				Count                float64 `json:"count"`
				TotalCapacityInBytes float64 `json:"total_capacity_in_bytes"`
				UsedInBytes          float64 `json:"used_in_bytes"`
			} `json:"mapped"`
		} `json:"buffer_pools"`
		Gc struct {
			Collectors struct {
				Old struct {
					CollectionCount        float64 `json:"collection_count"`
					CollectionTimeInMillis float64 `json:"collection_time_in_millis"`
				} `json:"old"`
				Young struct {
					CollectionCount        float64 `json:"collection_count"`
					CollectionTimeInMillis float64 `json:"collection_time_in_millis"`
				} `json:"young"`
			} `json:"collectors"`
		} `json:"gc"`
		Mem struct {
			HeapCommittedInBytes    float64 `json:"heap_committed_in_bytes"`
			HeapMaxInBytes          float64 `json:"heap_max_in_bytes"`
			HeapUsedInBytes         float64 `json:"heap_used_in_bytes"`
			HeapUsedPercent         float64 `json:"heap_used_percent"`
			NonHeapCommittedInBytes float64 `json:"non_heap_committed_in_bytes"`
			NonHeapUsedInBytes      float64 `json:"non_heap_used_in_bytes"`
			Pools                   struct {
				Old struct {
					MaxInBytes      float64 `json:"max_in_bytes"`
					PeakMaxInBytes  float64 `json:"peak_max_in_bytes"`
					PeakUsedInBytes float64 `json:"peak_used_in_bytes"`
					UsedInBytes     float64 `json:"used_in_bytes"`
				} `json:"old"`
				Survivor struct {
					MaxInBytes      float64 `json:"max_in_bytes"`
					PeakMaxInBytes  float64 `json:"peak_max_in_bytes"`
					PeakUsedInBytes float64 `json:"peak_used_in_bytes"`
					UsedInBytes     float64 `json:"used_in_bytes"`
				} `json:"survivor"`
				Young struct {
					MaxInBytes      float64 `json:"max_in_bytes"`
					PeakMaxInBytes  float64 `json:"peak_max_in_bytes"`
					PeakUsedInBytes float64 `json:"peak_used_in_bytes"`
					UsedInBytes     float64 `json:"used_in_bytes"`
				} `json:"young"`
			} `json:"pools"`
		} `json:"mem"`
		Threads struct {
			Count     float64 `json:"count"`
			PeakCount float64 `json:"peak_count"`
		} `json:"threads"`
		Timestamp      float64 `json:"timestamp"`
		UptimeInMillis float64 `json:"uptime_in_millis"`
	} `json:"jvm"`
	Name string `json:"name"`
	Os   struct {
		CPU struct {
			Idle   float64 `json:"idle"`
			Stolen float64 `json:"stolen"`
			Sys    float64 `json:"sys"`
			Usage  float64 `json:"usage"`
			User   float64 `json:"user"`
		} `json:"cpu"`
		LoadAverage []float64 `json:"load_average"`
		Mem         struct {
			ActualFreeInBytes float64 `json:"actual_free_in_bytes"`
			ActualUsedInBytes float64 `json:"actual_used_in_bytes"`
			FreeInBytes       float64 `json:"free_in_bytes"`
			FreePercent       float64 `json:"free_percent"`
			UsedInBytes       float64 `json:"used_in_bytes"`
			UsedPercent       float64 `json:"used_percent"`
		} `json:"mem"`
		Swap struct {
			FreeInBytes float64 `json:"free_in_bytes"`
			UsedInBytes float64 `json:"used_in_bytes"`
		} `json:"swap"`
		Timestamp      float64 `json:"timestamp"`
		UptimeInMillis float64 `json:"uptime_in_millis"`
	} `json:"os"`
	Process struct {
		CPU struct {
			Percent       float64 `json:"percent"`
			SysInMillis   float64 `json:"sys_in_millis"`
			TotalInMillis float64 `json:"total_in_millis"`
			UserInMillis  float64 `json:"user_in_millis"`
		} `json:"cpu"`
		Mem struct {
			ResidentInBytes     float64 `json:"resident_in_bytes"`
			ShareInBytes        float64 `json:"share_in_bytes"`
			TotalVirtualInBytes float64 `json:"total_virtual_in_bytes"`
		} `json:"mem"`
		OpenFileDescriptors float64 `json:"open_file_descriptors"`
		Timestamp           float64 `json:"timestamp"`
	} `json:"process"`
	ThreadPool struct {
		Bulk struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"bulk"`
		Flush struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"flush"`
		Generic struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"generic"`
		Get struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"get"`
		Index struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"index"`
		Listener struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"listener"`
		Management struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"management"`
		Merge struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"merge"`
		Optimize struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"optimize"`
		Percolate struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"percolate"`
		Refresh struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"refresh"`
		Search struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"search"`
		Snapshot struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"snapshot"`
		Suggest struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"suggest"`
		Warmer struct {
			Active    float64 `json:"active"`
			Completed float64 `json:"completed"`
			Largest   float64 `json:"largest"`
			Queue     float64 `json:"queue"`
			Rejected  float64 `json:"rejected"`
			Threads   float64 `json:"threads"`
		} `json:"warmer"`
	} `json:"thread_pool"`
	Timestamp float64 `json:"timestamp"`
}
