package schema

import (
	"fmt"
	"strings"
)

func ParseGroupVersion(gv string) (GroupVersion, error) {
	if (len(gv) == 0) || (gv == "/") {
		return GroupVersion{}, nil
	}

	switch strings.Count(gv, "/") {
	case 0:
		return GroupVersion{"", gv}, nil
	case 1:
		i := strings.Index(gv, "/")
		return GroupVersion{gv[:i], gv[i+1:]}, nil
	default:
		return GroupVersion{}, fmt.Errorf("unexpected GroupVersion string: %v", gv)
	}
}

type GroupVersion struct {
	Group   string
	Version string
}

func (gv GroupVersion) IsZero() bool {
	return len(gv.Group) == 0 && len(gv.Version) == 0
}

func (gv GroupVersion) String() string {
	if gv.Group != "" {
		return gv.Group + "/" + gv.Version
	}
	return gv.Version
}

func (gk GroupVersion) WithKind(kind string) GroupVersionKind {
	return GroupVersionKind{Group: gk.Group, Version: gk.Version, Kind: kind}
}

type GroupKind struct {
	Group string
	Kind  string
}

func (gk GroupKind) IsZero() bool {
	return len(gk.Group) == 0 && len(gk.Kind) == 0
}

func (gk GroupKind) WithVersion(version string) GroupVersionKind {
	return GroupVersionKind{Group: gk.Group, Version: version, Kind: gk.Kind}
}

func FromAPIVersionAndKind(apiVersion, kind string) GroupVersionKind {
	if gv, err := ParseGroupVersion(apiVersion); err == nil {
		return GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
	}
	return GroupVersionKind{Kind: kind}
}

type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

func (gvk GroupVersionKind) IsZero() bool {
	return len(gvk.Group) == 0 && len(gvk.Version) == 0 && len(gvk.Kind) == 0
}

func (gvk GroupVersionKind) GroupKind() GroupKind {
	return GroupKind{Group: gvk.Group, Kind: gvk.Kind}
}

func (gvk GroupVersionKind) GroupVersion() GroupVersion {
	return GroupVersion{Group: gvk.Group, Version: gvk.Version}
}

func (gvk GroupVersionKind) ToAPIVersionAndKind() (string, string) {
	if gvk.IsZero() {
		return "", ""
	}
	return gvk.GroupVersion().String(), gvk.Kind
}
