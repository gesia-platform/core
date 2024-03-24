package chaintree

type ChainTree struct {
	Root *Root
	Host *Host
}

func NewChainTree(
	rootURL string,
	hostURL string,
) *ChainTree {
	root := NewRoot(rootURL)

	host := NewHost(hostURL)

	return &ChainTree{Root: &root, Host: &host}

}
