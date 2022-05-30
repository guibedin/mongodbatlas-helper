package model

import "fmt"

type SimpleConnectionStrings struct {
	Standard    string
	StandardSrv string
}

func (scs SimpleConnectionStrings) String() string {
	return fmt.Sprintf("[ Standard: %s - StandardSrv: %s", scs.Standard, scs.StandardSrv)
}
