package objectkindgen

import (
	"go/types"
	"reflect"
	"strings"
	"sync"

	"github.com/octohelm/gengo/pkg/inflector"

	"github.com/octohelm/gengo/pkg/gengo"
	"github.com/octohelm/gengo/pkg/gengo/snippet"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/runtime"
)

func init() {
	gengo.Register(&objectKindGen{})
}

type objectKindGen struct {
	once sync.Once

	objectTypeInterface               *types.Interface
	objectRefCodeConvertableInterface *types.Interface

	objectRefIDConvertableInterface       *types.Interface
	objectRefStringIDConvertableInterface *types.Interface
}

func (*objectKindGen) Name() string {
	return "objectkind"
}

func (g *objectKindGen) init(c gengo.Context) {
	g.objectTypeInterface = typeInterfaceFor[object.Type](c)

	g.objectRefIDConvertableInterface = typeInterfaceFor[object.RefIDConvertable](c)
	g.objectRefStringIDConvertableInterface = typeInterfaceFor[object.RefStringIDConvertable](c)

	g.objectRefCodeConvertableInterface = typeInterfaceFor[object.RefCodeConvertable](c)
}

func typeInterfaceFor[T any](c gengo.Context) *types.Interface {
	t := reflect.TypeFor[T]()

	return c.Package(t.PkgPath()).Type(t.Name()).Type().(*types.Named).Underlying().(*types.Interface)
}

func (g *objectKindGen) GenerateType(c gengo.Context, t *types.Named) error {
	tags, _ := c.Doc(t.Obj())

	if !gengo.IsGeneratorEnabled(g, tags) {
		return nil
	}

	g.once.Do(func() {
		g.init(c)
	})

	if _, ok := tags["gengo:objectkind:variant"]; ok {
		g.generateObjectVariant(c, t)
	} else {
		g.generateObjectKind(c, t, nil)
	}

	return nil
}

func (g *objectKindGen) generateObjectVariant(c gengo.Context, t *types.Named) {
	structType, ok := t.Underlying().(*types.Struct)
	if ok {
		if structType.NumFields() > 0 {
			if head, ok := structType.Field(0).Type().(*types.Named); ok {
				if head.TypeArgs().Len() > 0 {
					baseType := head.TypeArgs().At(0).(*types.Named)

					g.generateObjectKind(c, t, baseType)
					g.generateObjectVariantCopies(c, t, baseType)
				}
			}
		}
	}
}

func (g *objectKindGen) generateObjectKind(c gengo.Context, t *types.Named, as *types.Named) {
	singularTypeName := func() string {
		if as != nil {
			return as.Obj().Name()
		}
		return t.Obj().Name()
	}()

	c.RenderT(`
func(@Type) GetKind() string {
	return @TypeName
}

func(@Type) GetPluralizedKind() string {
	return @PluralizedTypeName
}

func(@Type) GetAPIVersion() string {
	return SchemeGroupVersion.String()
}

`, snippet.Args{
		"Type": snippet.ID(t.Obj()),

		"TypeName":           snippet.Value(singularTypeName),
		"PluralizedTypeName": snippet.Value(inflector.Pluralize(singularTypeName)),
	})

	if s, ok := t.Underlying().(*types.Struct); ok {
		for i := 0; i < s.NumFields(); i++ {
			st := reflect.StructTag(s.Tag(i))

			if rel, ok := st.Lookup("as"); ok {
				switch rel {
				case "owner":
					f := s.Field(i)

					c.RenderT(`
func(t @Type) GetOwner() @objectType {
	return t.@FieldName
}

`, snippet.Args{
						"Type":       snippet.ID(t.Obj()),
						"FieldName":  snippet.ID(f.Name()),
						"objectType": snippet.PkgExposeFor[object.Type](),
					})
				}
				break
			}

		}
	}
}

func (g *objectKindGen) generateObjectVariantCopies(c gengo.Context, variantType *types.Named, baseType *types.Named) {
	for i, pair := range [][]*types.Named{
		{baseType, variantType},
		{variantType, baseType},
	} {
		c.RenderT(`
func (src *@SrcType) As@DstType() *@DstType {
	dst := @runtimeNew[@DstType]()
	@fieldsCopy
	return dst
}

`, snippet.Args{
			"SrcType": snippet.ID(pair[0]),
			"DstType": snippet.ID(pair[1]),

			"runtimeNew": snippet.PkgExposeFor[runtime.R]("New"),

			"fieldsCopy": &structCopy{
				Context:       c,
				objectKindGen: g,
				src:           pair[0],
				dst:           pair[1],
				srcVar:        "src",
				dstVar:        "dst",
				canMissing:    i == 0,
			},
		})
	}
}

func (g *objectKindGen) isMetaV1Exposed(t types.Type) bool {
	if n, ok := t.(*types.Named); ok {
		return strings.HasSuffix(n.Obj().Pkg().Path(), "/meta/v1")
	}
	return false
}

func (g *objectKindGen) isCodableObject(t types.Type) bool {
	return g.isObject(t) && implements(t, g.objectRefCodeConvertableInterface)
}

func (g *objectKindGen) isObject(t types.Type) bool {
	return implements(t, g.objectRefIDConvertableInterface) || implements(t, g.objectRefStringIDConvertableInterface)
}

func (g *objectKindGen) isObjectType(c gengo.Context, t types.Type) bool {
	return implements(t, g.objectTypeInterface)
}

func implements(t types.Type, itype *types.Interface) bool {
	return types.Implements(t, itype) || types.Implements(types.NewPointer(t), itype)
}
