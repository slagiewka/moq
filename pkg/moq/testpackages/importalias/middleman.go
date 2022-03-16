package importalias

import (
	srcclient "github.com/slagiewka/moq/pkg/moq/testpackages/importalias/source/client"
	tgtclient "github.com/slagiewka/moq/pkg/moq/testpackages/importalias/target/client"
)

type MiddleMan interface {
	Connect(src srcclient.Client, tgt tgtclient.Client)
}
