package rpccall

type Health struct {
	Peers           int
	IsSyncing       bool
	ShouldHavePeers bool
}
