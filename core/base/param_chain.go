package base

type ParamChain interface {
	DirectParams() []Node
}

type paramChain struct {
	chain [][]Node
}

func NewParamChain(chain [][]Node) ParamChain {
	return &paramChain{chain: chain}
}

func (c *paramChain) DirectParams() []Node {
	if len(c.chain) == 0 {
		return nil
	}
	return c.chain[0]
}
