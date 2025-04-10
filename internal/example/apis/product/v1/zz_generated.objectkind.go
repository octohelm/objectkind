/*
Package v1 GENERATED BY gengo:objectkind
DON'T EDIT THIS FILE
*/
package v1

import (
	iter "iter"

	object "github.com/octohelm/objectkind/pkg/object"
	pkgruntime "github.com/octohelm/objectkind/pkg/runtime"
)

func (Product) GetKind() string {
	return "Product"
}

func (Product) GetPluralizedKind() string {
	return "Products"
}

func (Product) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (ProductReference) GetKind() string {
	return "Product"
}

func (ProductReference) GetPluralizedKind() string {
	return "Products"
}

func (ProductReference) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (src *Product) AsProductReference() *ProductReference {
	dst := pkgruntime.New[ProductReference]()
	pkgruntime.CopyObject(dst, src)

	return dst
}

func (src *ProductReference) AsProduct() *Product {
	dst := pkgruntime.New[Product]()
	pkgruntime.CopyObject(dst, src)

	return dst
}

func (ProductRequestForCreate) GetKind() string {
	return "Product"
}

func (ProductRequestForCreate) GetPluralizedKind() string {
	return "Products"
}

func (ProductRequestForCreate) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (src *Product) AsProductRequestForCreate() *ProductRequestForCreate {
	dst := pkgruntime.New[ProductRequestForCreate]()
	pkgruntime.CopyObject(dst, src)

	return dst
}

func (src *ProductRequestForCreate) AsProduct() *Product {
	dst := pkgruntime.New[Product]()
	pkgruntime.Copy(dst, src)

	return dst
}

func (ProductRequestForUpdate) GetKind() string {
	return "Product"
}

func (ProductRequestForUpdate) GetPluralizedKind() string {
	return "Products"
}

func (ProductRequestForUpdate) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (src *Product) AsProductRequestForUpdate() *ProductRequestForUpdate {
	dst := pkgruntime.New[ProductRequestForUpdate]()
	pkgruntime.CopyObject(dst, src)

	return dst
}

func (src *ProductRequestForUpdate) AsProduct() *Product {
	dst := pkgruntime.New[Product]()
	pkgruntime.Copy(dst, src)

	return dst
}

func (Sku) GetKind() string {
	return "Sku"
}

func (Sku) GetPluralizedKind() string {
	return "Skus"
}

func (Sku) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (v *Sku) Parents() iter.Seq[object.Type] {
	return func(yield func(object.Type) bool) {
		if v.Product != nil {
			if !yield(v.Product) {
				return
			}

			if x, ok := any(v.Product).(object.ParentIter); ok {
				for t := range x.Parents() {
					if !yield(t) {
						return
					}
				}
			}
		}
	}
}

func (SkuReference) GetKind() string {
	return "Sku"
}

func (SkuReference) GetPluralizedKind() string {
	return "Skus"
}

func (SkuReference) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (src *Sku) AsSkuReference() *SkuReference {
	dst := pkgruntime.New[SkuReference]()
	pkgruntime.CopyCodableObject(dst, src)

	return dst
}

func (src *SkuReference) AsSku() *Sku {
	dst := pkgruntime.New[Sku]()
	pkgruntime.CopyObject(dst, src)

	return dst
}

func (SkuRequestForCreate) GetKind() string {
	return "Sku"
}

func (SkuRequestForCreate) GetPluralizedKind() string {
	return "Skus"
}

func (SkuRequestForCreate) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (src *Sku) AsSkuRequestForCreate() *SkuRequestForCreate {
	dst := pkgruntime.New[SkuRequestForCreate]()
	pkgruntime.CopyCodableObject(dst, src)
	dst.Spec = src.Spec

	return dst
}

func (src *SkuRequestForCreate) AsSku() *Sku {
	dst := pkgruntime.New[Sku]()
	pkgruntime.CopyCodable(dst, src)
	dst.Spec = src.Spec

	return dst
}

func (SkuRequestForUpdate) GetKind() string {
	return "Sku"
}

func (SkuRequestForUpdate) GetPluralizedKind() string {
	return "Skus"
}

func (SkuRequestForUpdate) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

func (src *Sku) AsSkuRequestForUpdate() *SkuRequestForUpdate {
	dst := pkgruntime.New[SkuRequestForUpdate]()
	pkgruntime.CopyCodableObject(dst, src)
	dst.Spec = src.Spec

	return dst
}

func (src *SkuRequestForUpdate) AsSku() *Sku {
	dst := pkgruntime.New[Sku]()
	pkgruntime.Copy(dst, src)
	dst.Spec = src.Spec

	return dst
}
