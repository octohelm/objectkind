package product

import (
	"fmt"

	"github.com/octohelm/courier/pkg/statuserror"
)

type ErrSkuNotFound struct {
	statuserror.NotFound
}

func (e *ErrSkuNotFound) Error() string {
	return fmt.Sprintf("SKU 不存在")
}
