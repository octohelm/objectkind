package product

import (
	"fmt"

	"github.com/octohelm/courier/pkg/statuserror"
)

type ErrProductNotFound struct {
	statuserror.NotFound
}

func (e *ErrProductNotFound) Error() string {
	return fmt.Sprintf("产品不存在")
}
