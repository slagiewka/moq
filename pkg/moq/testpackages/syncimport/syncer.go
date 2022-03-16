package syncimport

import (
	stdsync "sync"

	"github.com/slagiewka/moq/pkg/moq/testpackages/syncimport/sync"
)

type Syncer interface {
	Blah(s sync.Thing, wg *stdsync.WaitGroup)
}
