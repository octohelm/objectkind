package v1

import (
	"fmt"

	"github.com/octohelm/courier/pkg/statuserror"
)

type ErrProductForbidden struct {
	statuserror.Forbidden
}

func (e *ErrProductForbidden) Error() string {
	return fmt.Sprintf("没有商品管理权限")
}

type ErrProductNotFound struct {
	statuserror.NotFound
}

func (e *ErrProductNotFound) Error() string {
	return fmt.Sprintf("商品不存在")
}

type ErrProductStateConflict struct {
	statuserror.Conflict
}

func (e *ErrProductStateConflict) Error() string {
	return fmt.Sprintf("商品当前状态不允许执行该操作")
}

type ErrSkuNotFound struct {
	statuserror.NotFound
}

func (e *ErrSkuNotFound) Error() string {
	return fmt.Sprintf("商品规格不存在")
}
