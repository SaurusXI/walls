package src_port

import (
	"github.com/SaurusXI/protecc/src/gate"
)

type Config struct {
	Setting gate.Gate
	Blocked map[string]bool
}
