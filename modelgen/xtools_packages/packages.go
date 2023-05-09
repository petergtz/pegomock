package xtools_packages

import (
	"errors"
	"fmt"
	"go/types"
	"path"

	"github.com/petergtz/pegomock/v3/model"
	"golang.org/x/tools/go/packages"
)

type Bla[K comparable, V Number] interface {
	SumNumbers(m map[K]V, i int, s string, a []float32, sss ...string) V
}

type Number interface {
	int64 | float64
}

type Blub[K comparable, V Number] struct{}

func (b *Blub[K, V]) SumNumbers(m map[K]V, i int, q string, a []float32) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func NewBlub[K comparable, V Number]() *Blub[K, V] {
	return &Blub[K, V]{}
}

func GenerateModel(importPath string, interfaceName string) (*model.Package, error) {

	pkgs, e := packages.Load(&packages.Config{Mode: packages.NeedTypes}, importPath)
	if e != nil {
		return nil, e
	}
	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		obj := scope.Lookup(interfaceName)
		if obj == nil {
			continue
		}
		// from here, things follow the spec in https://tip.golang.org/ref/spec
		if typeName, isTypeName := obj.(*types.TypeName); isTypeName {
			if iface, isIface := typeName.Type().Underlying().(*types.Interface); isIface {

				g := &modelGenerator2{typeParams: make(map[string]*model.Parameter)}
				methods := g.modelMethodsFrom(iface)

				return &model.Package{
					Name: path.Base(pkg.Types.Name()),
					Interfaces: []*model.Interface{{
						Name:       interfaceName,
						Methods:    methods,
						TypeParams: sliceFrom(g.typeParams),
					}},
				}, nil
			}
		}
	}

	return nil, errors.New("Did not find interface name \"" + interfaceName + "\"")
}

type modelGenerator2 struct {
	typeParams map[string]*model.Parameter
}

func (g *modelGenerator2) modelMethodsFrom(iface *types.Interface) (modelMethods []*model.Method) {
	for i := 0; i < iface.NumMethods(); i++ {
		modelMethods = append(modelMethods, g.modelMethodFrom(iface.Method(i)))
	}
	return
}

func (g *modelGenerator2) modelMethodFrom(method *types.Func) *model.Method {
	signature := method.Type().(*types.Signature)
	in, variadic := g.inParamsFrom(signature)
	return &model.Method{
		Name:     method.Name(),
		In:       in,
		Variadic: variadic,
		Out:      g.outParamsFrom(signature),
	}
}

func (g *modelGenerator2) inParamsFrom(signature *types.Signature) (in []*model.Parameter, variadic *model.Parameter) {
	for u := 0; u < signature.Params().Len(); u++ {
		if signature.Variadic() && u == signature.Params().Len()-1 {
			variadic = &model.Parameter{
				Name: signature.Params().At(u).Name(),
				Type: g.modelTypeFrom(signature.Params().At(u).Type().(*types.Slice).Elem()),
			}
			break
		}
		in = append(in, &model.Parameter{
			Name: signature.Params().At(u).Name(),
			Type: g.modelTypeFrom(signature.Params().At(u).Type()),
		})
	}
	return
}

func (g *modelGenerator2) outParamsFrom(signature *types.Signature) (out []*model.Parameter) {
	if signature.Results() != nil {
		for u := 0; u < signature.Results().Len(); u++ {
			out = append(out, &model.Parameter{
				Name: signature.Results().At(u).Name(),
				Type: g.modelTypeFrom(signature.Results().At(u).Type()),
			})
		}
	}
	return
}

func (g *modelGenerator2) modelTypeFrom(typesType types.Type) model.Type {
	switch typedTyp := typesType.(type) {
	case *types.Basic:
		if !predeclared(typedTyp.Kind()) {
			panic(fmt.Sprintf("Unexpected Basic Type %v", typedTyp.Name()))
		}
		return model.PredeclaredType(typedTyp.Name())
	case *types.Pointer:
		return &model.PointerType{
			Type: g.modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Array:
		return &model.ArrayType{
			Len:  int(typedTyp.Len()),
			Type: g.modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Slice:
		return &model.ArrayType{
			Len:  -1,
			Type: g.modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Map:
		return &model.MapType{
			Key:   g.modelTypeFrom(typedTyp.Key()),
			Value: g.modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Chan:
		var dir model.ChanDir
		switch typedTyp.Dir() {
		case types.SendOnly:
			dir = model.SendDir
		case types.RecvOnly:
			dir = model.RecvDir
		default:
			dir = 0
		}
		return &model.ChanType{
			Dir:  dir,
			Type: g.modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Named:
		if typedTyp.Obj().Pkg() == nil {
			return model.PredeclaredType(typedTyp.Obj().Name())
		}
		return &model.NamedType{
			Package: typedTyp.Obj().Pkg().Path(),
			Type:    typedTyp.Obj().Name(),
		}
	case *types.Interface, *types.Struct:
		return model.PredeclaredType(typedTyp.String())
	case *types.Signature:
		in, variadic := g.inParamsFrom(typedTyp)
		return &model.FuncType{In: in, Out: g.outParamsFrom(typedTyp), Variadic: variadic}
	case *types.TypeParam:
		g.typeParams[typedTyp.Obj().Name()] = &model.Parameter{
			Name: typedTyp.Obj().Name(),
			Type: g.modelTypeFrom(typedTyp.Constraint()),
		}
		return model.PredeclaredType(typedTyp.Obj().Name())
	default:
		panic(fmt.Sprintf("Unknown types.Type: %v (%T)", typesType, typesType))
	}
}

func sliceFrom(params map[string]*model.Parameter) (result []*model.Parameter) {
	for _, v := range params {
		result = append(result, v)
	}
	return
}

func predeclared(basicKind types.BasicKind) bool {
	return basicKind >= types.Bool && basicKind <= types.String
}
