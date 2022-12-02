package rpccall

type PeerInfo struct {
	BestHash        string
	BestNumber      int
	PeerID          string
	Index           int
	ProtocolVersion int
	roles           string
}
