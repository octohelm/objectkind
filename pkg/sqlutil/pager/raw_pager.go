package pager

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

// Deprecated use metav1.Pager instead
//
//go:fix inline
type RawPager = metav1.Pager
