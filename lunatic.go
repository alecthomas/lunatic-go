package lunatic

import "github.com/alecthomas/lunatic-go/internal/bindings/lunatic/version"

func Version() (uint32, uint32, uint32) {
	return version.Major(), version.Minor(), version.Patch()
}

type Mailbox struct {
}

func Main(mailbox *Mailbox) {
}
