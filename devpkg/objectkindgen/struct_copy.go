package objectkindgen

import (
	"context"
	"fmt"
	"go/types"
	"iter"
	"reflect"

	"github.com/octohelm/gengo/pkg/gengo"
	"github.com/octohelm/gengo/pkg/gengo/snippet"
	"github.com/octohelm/objectkind/pkg/runtime"
)

type structCopy struct {
	*objectKindGen

	gengo.Context
	canMissing bool

	srcVar string
	dstVar string

	dst *types.Named
	src *types.Named
}

func (c *structCopy) IsNil() bool {
	return false
}

func (c *structCopy) Frag(ctx context.Context) iter.Seq[string] {
	return snippet.Snippets(func(yield func(snippet.Snippet) bool) {
		if c.isObjectType(c.Context, c.src) {
			if !yield(snippet.T(`
@runtimeCopy(@dst, @src)
`, snippet.Args{
				"src": snippet.ID(c.srcVar),
				"dst": snippet.ID(c.dstVar),

				"runtimeCopy": func() snippet.Snippet {
					if c.isObject(c.src) {
						if c.isCodable(c.src) {
							return snippet.PkgExposeFor[runtime.R]("CopyCodableObject")
						}
						return snippet.PkgExposeFor[runtime.R]("CopyObject")
					}
					if c.isCodable(c.src) {
						return snippet.PkgExposeFor[runtime.R]("CopyCodable")
					}
					return snippet.PkgExposeFor[runtime.R]("Copy")
				}(),
			})) {
				return
			}
		}

		dstFields := c.dst.Underlying().(*types.Struct)
		originFields := map[string]*types.Var{}
		for i := 0; i < dstFields.NumFields(); i++ {
			f := dstFields.Field(i)
			originFields[f.Name()] = f
		}

		if s, ok := c.src.Underlying().(*types.Struct); ok {
			for i := 0; i < s.NumFields(); i++ {
				srcField := s.Field(i)
				tag := reflect.StructTag(s.Tag(i))

				if c.isMetaV1Exposed(srcField.Type()) {
					continue
				}

				if v, ok := tag.Lookup("json"); ok {
					if v == "-" {
						continue
					}

					dstField, ok := originFields[srcField.Name()]
					if !ok {
						if c.canMissing {
							continue
						}

						panic(fmt.Errorf("dst struct %s missing field %s", c.dst, srcField.Name()))
					}

					if !yield(&fieldCopy{
						structCopy: c,
						srcField:   srcField,
						dstField:   dstField,
					}) {
						return
					}
				}
			}
		} else {
			panic(fmt.Errorf("type unmatched: dst: %s, src: %s", c.dst, c.src))
		}
	}).Frag(ctx)
}

type fieldCopy struct {
	*structCopy

	srcField *types.Var
	dstField *types.Var
}

func (c *fieldCopy) IsNil() bool {
	return false
}

func (c *fieldCopy) Frag(ctx context.Context) iter.Seq[string] {
	if types.AssignableTo(c.srcField.Type(), c.dstField.Type()) {
		return snippet.Snippets(func(yield func(snippet.Snippet) bool) {
			if !yield(snippet.T(`
@dst.@FieldName = @src.@FieldName
`, snippet.Args{
				"FieldName": snippet.ID(c.name()),
				"src":       snippet.ID(c.srcVar),
				"dst":       snippet.ID(c.dstVar),
			})) {
				return
			}
		}).Frag(ctx)
	}

	srcType, isSrcPtr := underlyingType(c.srcField.Type())
	dstType, isDstPtr := underlyingType(c.dstField.Type())

	if srcType == dstType {
		return snippet.Snippets(func(yield func(snippet.Snippet) bool) {
			if isDstPtr && !isSrcPtr {
				if !yield(snippet.T(`
if @dst.@FieldName == nil {
	@dst.@FieldName = new(@DstType)
}
@dst.@FieldName = &@src.@FieldName
`, snippet.Args{
					"FieldName": snippet.ID(c.name()),
					"DstType":   snippet.ID(dstType),
					"src":       snippet.ID(c.srcVar),
					"dst":       snippet.ID(c.dstVar),
				})) {
					return
				}
				return
			}

			if !isDstPtr && isSrcPtr {
				if !yield(snippet.T(`
if @src.@FieldName != nil {
	@dst.@FieldName = *@src.@FieldName
}
`, snippet.Args{
					"FieldName": snippet.ID(c.name()),
					"src":       snippet.ID(c.srcVar),
					"dst":       snippet.ID(c.dstVar),
				})) {
					return
				}
				return
			}

			if !yield(snippet.T(`
@dst.@FieldName = @src.@FieldName
`, snippet.Args{
				"FieldName": snippet.ID(c.name()),
				"src":       snippet.ID(c.srcVar),
				"dst":       snippet.ID(c.dstVar),
			})) {
				return
			}
		}).Frag(ctx)
	}

	_, isSrcStruct := srcType.Underlying().(*types.Struct)
	_, isDstStruct := dstType.Underlying().(*types.Struct)

	if isSrcStruct && isDstStruct {
		doCopy := snippet.Snippets(func(yield func(snippet.Snippet) bool) {
			if !yield(snippet.T(`
copy@FieldName := func(d *@DstType, s *@SrcType) {
	@fieldsCopy
}
`, snippet.Args{
				"FieldName": snippet.ID(c.name()),
				"SrcType":   snippet.ID(srcType),
				"DstType":   snippet.ID(dstType),

				"fieldsCopy": &structCopy{
					Context:       c.Context,
					objectKindGen: c.objectKindGen,
					canMissing:    c.canMissing,
					dstVar:        "d",
					srcVar:        "s",
					dst:           dstType.(*types.Named),
					src:           srcType.(*types.Named),
				},
			})) {
				return
			}

			if isDstPtr {
				if !yield(snippet.T(`
if @dst.@FieldName == nil {
	@dst.@FieldName = new(@DstType)
}
`, snippet.Args{
					"FieldName": snippet.ID(c.name()),
					"DstType":   snippet.ID(dstType),
					"src":       snippet.ID(c.srcVar),
					"dst":       snippet.ID(c.dstVar),
				})) {
					return
				}

				if isSrcPtr {
					yield(snippet.T(`
copy@FieldName(@dst.@FieldName, @src.@FieldName)
`, snippet.Args{
						"FieldName": snippet.ID(c.name()),
						"src":       snippet.ID(c.srcVar),
						"dst":       snippet.ID(c.dstVar),
					}))

					return
				}

				yield(snippet.T(`
copy@FieldName(@dst.@FieldName, &@src.@FieldName)
`, snippet.Args{
					"FieldName": snippet.ID(c.name()),
					"src":       snippet.ID(c.srcVar),
					"dst":       snippet.ID(c.dstVar),
				}))

				return
			}

			if isSrcPtr {
				yield(snippet.T(`
copy@FieldName(&@dst.@FieldName, @src.@FieldName)
`, snippet.Args{
					"FieldName": snippet.ID(c.name()),
					"src":       snippet.ID(c.srcVar),
					"dst":       snippet.ID(c.dstVar),
				}))

				return
			}

			yield(snippet.T(`
copy@FieldName(&@dst.@FieldName, &@src.@FieldName)
`, snippet.Args{
				"FieldName": snippet.ID(c.name()),
				"src":       snippet.ID(c.srcVar),
				"dst":       snippet.ID(c.dstVar),
			}))

			return
		})

		return snippet.Snippets(func(yield func(snippet.Snippet) bool) {
			if isSrcPtr {
				yield(snippet.T(`

if (@src.@FieldName != nil) {
	@doCopy
}
`, snippet.Args{
					"FieldName": snippet.ID(c.name()),
					"src":       snippet.ID(c.srcVar),
					"dst":       snippet.ID(c.dstVar),
					"doCopy":    doCopy,
				}))

				return
			}

			yield(doCopy)
			return
		}).Frag(ctx)
	}

	_, isSrcSlice := srcType.Underlying().(*types.Slice)
	_, isDstSlice := dstType.Underlying().(*types.Slice)

	if isSrcSlice && isDstSlice {
		srcElemType, _ := underlyingType(srcType.Underlying().(*types.Slice).Elem())
		dstElemType, _ := underlyingType(dstType.Underlying().(*types.Slice).Elem())

		_, isSrcSliceElemStruct := srcElemType.Underlying().(*types.Struct)
		_, isDstSliceElemStruct := dstElemType.Underlying().(*types.Struct)

		if isSrcSliceElemStruct && isDstSliceElemStruct {
			return snippet.Snippets(func(yield func(snippet.Snippet) bool) {
				if !yield(snippet.T(`
if n := len(@src.@FieldName); n > 0 {
	@dst.@FieldName = make(@DstFieldType, n)
	for i, x := range @src.@FieldName {
		@dst.@FieldName[i] = x.As@DstElemTypeName()
	}
}
`, snippet.Args{
					"FieldName":       snippet.ID(c.name()),
					"DstFieldType":    snippet.ID(dstType),
					"DstElemTypeName": snippet.ID(dstElemType.(*types.Named).Obj().Name()),
					"src":             snippet.ID(c.srcVar),
					"dst":             snippet.ID(c.dstVar),
				})) {
					return
				}
			}).Frag(ctx)
		}
	}

	panic(fmt.Errorf("field %s is not copy from %s to %s", c.name(), c.srcField.Type(), c.dstField.Type()))

	return func(yield func(string) bool) {
	}
}

func (c *fieldCopy) name() string {
	return c.srcField.Name()
}

func underlyingType(t types.Type) (types.Type, bool) {
	v, ok := t.(*types.Pointer)
	if ok {
		return v.Elem(), true
	}
	return t, false
}
