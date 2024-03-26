package chaintree

type ChainTree struct {
	Root *Root
	Host *Host
}

func NewChainTree(
	rootURL string,
	rootWSURL string,
	hostURL string,
	hostWSURL string,
) *ChainTree {
	root := NewRoot(rootURL, rootWSURL)

	host := NewHost(hostURL, hostWSURL)

	return &ChainTree{Root: &root, Host: &host}

}
