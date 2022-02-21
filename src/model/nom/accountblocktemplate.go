package nom

type BlockTypeEnum int

const (
	unknown         BlockTypeEnum = iota
	genesisReceive                = iota
	userSend                      = iota
	userReceive                   = iota
	contractSend                  = iota
	contractReceive               = iota
)
