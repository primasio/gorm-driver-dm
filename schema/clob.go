package dmSchema

import (
	"database/sql/driver"
	"fmt"
	"github.com/nfjBill/gorm-driver-dm/dmr"
)

type Clob string

func (clob Clob) Value() (driver.Value, error) {
	if len(clob) == 0 {
		return nil, nil
	}
	return string(clob), nil
}

func (clob *Clob) Scan(v interface{}) error {
	switch v.(type) {
	case *dmr.DmClob:
		tmp := v.(*dmr.DmClob)
		le, err := tmp.GetLength()
		if err != nil {
			return fmt.Errorf("err：%w", err)
		}

		str, err := tmp.ReadString(1, int(le))
		*clob = Clob(str)
		break
	case []uint8:
		*clob = Clob(v.([]uint8))
	default:
		*clob = Clob(v.(string))
	}
	return nil
}
