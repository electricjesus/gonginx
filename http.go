package gonginx

import (
	"errors"
)

//Http represents http block
type Http struct {
	Block IBlock
}

//NewHttp create an http block from a directive which has a block
func NewHttp(directive IDirective) (*Http, error) {
	if block := directive.GetBlock(); block != nil {
		return &Http{
			Block: block,
		}, nil
	}
	return nil, errors.New("http directive must have a block")
}

//GetName get directive name to construct the statment string
func (h *Http) GetName() string { //the directive name.
	return "http"
}

//GetParameters get directive parameters if any
func (h *Http) GetParameters() []string {
	return []string{}
}

//GetBlock get block if any
func (h *Http) GetBlock() IBlock {
	return h.Block
}
