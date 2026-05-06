package v1

import (
	"fmt"

	"github.com/octohelm/courier/pkg/statuserror"
)

type ErrOrderForbidden struct {
	statuserror.Forbidden
}

func (e *ErrOrderForbidden) Error() string {
	return fmt.Sprintf("没有订单访问权限")
}

type ErrOrderNotFound struct {
	statuserror.NotFound
}

func (e *ErrOrderNotFound) Error() string {
	return fmt.Sprintf("订单不存在")
}

type ErrOrderStateConflict struct {
	statuserror.Conflict
}

func (e *ErrOrderStateConflict) Error() string {
	return fmt.Sprintf("订单当前状态不允许执行该操作")
}
