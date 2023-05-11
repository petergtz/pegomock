package xtools_packages

import (
	"errors"
	"fmt"
	"go/types"
	"path"

	"github.com/petergtz/pegomock/v4/model"
	"golang.org/x/tools/go/packages"
)

type Bla[K comparable, V Number] interface {
	SumNumbers(m map[K]V, i int, s string, a []float32, sss ...string) V
}

type Number interface {
	int64 | float64
}

type Blub[V Number, K comparable] struct{}

func (b *Blub[V, K]) SumNumbers(m map[K]V, i int, _ string, a []float32, sss ...string) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func NewBlub[K1 comparable, V1 Number]() Bla[K1, V1] {
	return &Blub[V1, K1]{}
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
		if iface, isIface := obj.Type().Underlying().(*types.Interface); isIface {
			return &model.Package{
				Name: path.Base(pkg.Types.Name()),
				Interfaces: []*model.Interface{{
					Name:       interfaceName,
					Methods:    modelMethodsFrom(iface),
					TypeParams: typeParamsFrom(obj.Type().(*types.Named).TypeParams()),
				}},
			}, nil

		}
	}

	return nil, errors.New("Did not find interface name \"" + interfaceName + "\"")
}

func modelMethodsFrom(iface *types.Interface) (modelMethods []*model.Method) {
	for i := 0; i < iface.NumMethods(); i++ {
		modelMethods = append(modelMethods, modelMethodFrom(iface.Method(i)))
	}
	return
}

func modelMethodFrom(method *types.Func) *model.Method {
	signature := method.Type().(*types.Signature)
	in, variadic := inParamsFrom(signature)
	return &model.Method{
		Name:     method.Name(),
		In:       in,
		Variadic: variadic,
		Out:      outParamsFrom(signature),
	}
}

func inParamsFrom(signature *types.Signature) (in []*model.Parameter, variadic *model.Parameter) {
	for u := 0; u < signature.Params().Len(); u++ {
		if signature.Variadic() && u == signature.Params().Len()-1 {
			variadic = &model.Parameter{
				Name: signature.Params().At(u).Name(),
				Type: modelTypeFrom(signature.Params().At(u).Type().(*types.Slice).Elem()),
			}
			break
		}
		in = append(in, &model.Parameter{
			Name: signature.Params().At(u).Name(),
			Type: modelTypeFrom(signature.Params().At(u).Type()),
		})
	}
	return
}

func outParamsFrom(signature *types.Signature) (out []*model.Parameter) {
	if signature.Results() != nil {
		for u := 0; u < signature.Results().Len(); u++ {
			out = append(out, &model.Parameter{
				Name: signature.Results().At(u).Name(),
				Type: modelTypeFrom(signature.Results().At(u).Type()),
			})
		}
	}
	return
}

func modelTypeFrom(typesType types.Type) model.Type {
	switch typedTyp := typesType.(type) {
	case *types.Basic:
		if !predeclared(typedTyp.Kind()) {
			panic(fmt.Sprintf("Unexpected Basic Type %v", typedTyp.Name()))
		}
		return model.PredeclaredType(typedTyp.Name())
	case *types.Pointer:
		return &model.PointerType{
			Type: modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Array:
		return &model.ArrayType{
			Len:  int(typedTyp.Len()),
			Type: modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Slice:
		return &model.ArrayType{
			Len:  -1,
			Type: modelTypeFrom(typedTyp.Elem()),
		}
	case *types.Map:
		return &model.MapType{
			Key:   modelTypeFrom(typedTyp.Key()),
			Value: modelTypeFrom(typedTyp.Elem()),
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
			Type: modelTypeFrom(typedTyp.Elem()),
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
		in, variadic := inParamsFrom(typedTyp)
		return &model.FuncType{In: in, Out: outParamsFrom(typedTyp), Variadic: variadic}
	case *types.TypeParam:
		return model.PredeclaredType(typedTyp.Obj().Name())
	default:
		panic(fmt.Sprintf("Unknown types.Type: %v (%T)", typesType, typesType))
	}
}

func typeParamsFrom(typeParams *types.TypeParamList) (result []*model.Parameter) {
	for i := 0; i < typeParams.Len(); i++ {
		result = append(result, &model.Parameter{
			Name: typeParams.At(i).Obj().Name(),
			Type: modelTypeFrom(typeParams.At(i).Constraint()),
		})
	}
	return
}

func predeclared(basicKind types.BasicKind) bool {
	return basicKind >= types.Bool && basicKind <= types.String
}
