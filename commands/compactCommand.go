package commands

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb/util"
)

func Compact() string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	err := dbh.CompactRange(util.Range{Start: nil, Limit: nil})
	if err != nil {
		return fmt.Sprintf(
			AppError(ErrUnableToDelete),
			err.Error(),
		)
	}

	return "Success"
}
