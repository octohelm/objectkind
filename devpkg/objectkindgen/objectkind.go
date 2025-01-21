package objectkindgen

import (
	"cmp"
	"go/types"
	"reflect"
	"sync"

	"github.com/octohelm/gengo/pkg/gengo"
	"github.com/octohelm/gengo/pkg/gengo/snippet"
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/schema"
)

func init() {
	gengo.Register(&objectKindGen{})
}

type objectKindGen struct {
	objectInterface *types.Interface
	once            sync.Once
}

func (*objectKindGen) Name() string {
	return "objectkind"
}

func (g *objectKindGen) init(c gengo.Context) {
	g.objectInterface = typeInterfaceFor[schema.Object](c)

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

	variants := tags["gengo:objectkind:variant"]
	if len(variants) > 0 {
		pkgPath, name := gengo.PkgImportPathAndExpose(variants[0])
		pkg := c.Package(cmp.Or(pkgPath, t.Obj().Pkg().Path()))
		g.generateObjectKind(c, t, name)
		g.generateObjectVariantCopies(c, t, pkg.Type(name).Type().(*types.Named))
	} else {
		g.generateObjectKind(c, t, "")
	}

	return nil
}

func (g *objectKindGen) generateObjectKind(c gengo.Context, t *types.Named, as string) {
	c.RenderT(`
func(@Type) GetObjectKind() @ObjectKind {
   apiVersion, kind := SchemeGroupVersion.WithKind("@TypeName").ToAPIVersionAndKind()

	return &@TypeMeta{
		APIVersion: apiVersion,
		Kind: kind,
	}
}
`, snippet.Args{
		"Type": snippet.ID(t.Obj()),
		"TypeName": func() snippet.Snippet {
			if as != "" {
				return snippet.ID(as)
			}
			return snippet.ID(t.Obj())
		}(),
		"TypeMeta":   snippet.PkgExposeFor[metav1.TypeMeta](),
		"ObjectKind": snippet.PkgExposeFor[metav1.ObjectKind](),
	})
}

func (g *objectKindGen) generateObjectVariantCopies(c gengo.Context, variantType *types.Named, baseType *types.Named) {
	for i, pair := range [][]*types.Named{
		{variantType, baseType},
		{baseType, variantType},
	} {
		c.RenderT(`
func (src *@SrcType) As@DstType() *@DstType {
	dst := new(@DstType)
	@fieldsCopy
	return dst
}

`, snippet.Args{
			"SrcType": snippet.ID(pair[0]),
			"DstType": snippet.ID(pair[1]),

			"fieldsCopy": &structCopy{
				Context:       c,
				objectKindGen: g,
				src:           pair[0],
				dst:           pair[1],
				srcVar:        "src",
				dstVar:        "dst",
				canMissing:    i > 0,
			},
		})
	}

}

func (g *objectKindGen) isMetaObjectHead(t types.Type) bool {
	return types.Implements(t, g.objectInterface) || types.Implements(types.NewPointer(t), g.objectInterface)
}

func (g *objectKindGen) variantOf(c gengo.Context, t types.Type) (*types.Named, bool) {
	switch x := t.(type) {
	case *types.Pointer:
		return g.variantOf(c, x.Elem())
	case *types.Named:
		tags, _ := c.Doc(x.Obj())
		variants := tags["gengo:objectkind:variant"]
		if len(variants) > 0 {
			pkgPath, name := gengo.PkgImportPathAndExpose(variants[0])
			pkg := c.Package(cmp.Or(pkgPath, x.Obj().Pkg().Path()))
			return pkg.Type(name).Type().(*types.Named), true
		}
	}
	return nil, false
}
