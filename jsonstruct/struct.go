package jsonstruct

type LocalNode struct {
	Id string
	HostName string
	Master string
	NumberOfNodesSeen float64
}

type Os struct {
	Name string
	Version string
	Cores float64
	LodeAverage float64
}

type Jvm struct {
	Name string
	Vendor string
	Version string
	StartTime float64
	UpTime float64
}

type Memory struct {
		Heap map[string]float64
		NonHeap map[string]float64
	}

type Gc struct {
	CollectionTime float64
	CollectionCount float64
}

type Index struct {
	Status string
	ActiveShards float64
	ActivePrimaryShards float64
	ActiveReplicas float64
	UnassignedShards float64
	RelocatingShards float64
	InitializingShards float64
	Documents string
	PrimaryShardsStoreSize string
	TotalStoreSize string
}

type Product struct {
	Name string
	Version string
	Edition string
}

type Member struct {
	Address string
	Id string
	IsMaster bool
	NumberOfNodesSeen float64
	Version string
}

type Cluster struct {
	LocalNode *LocalNode
	Members *map[string]Member
}

type XpClusterStatus struct {
	LocalNode *LocalNode
	Members *map[string]Member
}

type CmsStatus struct {
	Cluster *Cluster
	Os *Os
	Jvm *Jvm
	Memory *Memory
	Gc *Gc
	Index *Index
	Product *Product
}
