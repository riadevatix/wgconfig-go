package wgconfig

type Interface struct {
	PrivateKey string   `ini:"PrivateKey,omitempty"`
	Address    string   `ini:"Address,omitempty"`
	ListenPort uint     `ini:"ListenPort,omitempty"`
	DNS        []string `ini:"DNS,omitempty"`
	Table      int      `ini:"Table,omitempty"` // Routing table.
	MTU        int      `ini:"MTU,omitempty"`
	PreUp      string   `ini:"PreUp,omitempty"`
	PostUp     string   `ini:"PostUp,omitempty"`
	PreDown    string   `ini:"PreDown,omitempty"`
	PostDown   string   `ini:"PostDown,omitempty"`
}
