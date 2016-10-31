package loader

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/petergtz/pegomock/model"
	"golang.org/x/tools/go/loader"
)

func GenerateModel(importPath string, interfaceName string) (*model.Package, error) {
	var conf loader.Config
	conf.Import(importPath)
	program, e := conf.Load()
	if e != nil {
		panic(e)
	}
	info := program.Imported[importPath]

	for def := range info.Defs {
		if def.Name == interfaceName && def.Obj.Kind == ast.Typ {
			interfacetype, ok := def.Obj.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType)
			if ok {
				ast.Print(program.Fset, interfacetype)
				g := &modelGenerator{info: info}
				methods := g.generateMethods(interfacetype.Methods)
				iface := &model.Interface{Name: interfaceName, Methods: methods}
				return &model.Package{Name: info.Pkg.Name(), Interfaces: []*model.Interface{iface}}, nil
			}
		}
	}

	return nil, errors.New("Did not find interface name TODO")
}

type modelGenerator struct {
	info *loader.PackageInfo
}

func (g *modelGenerator) generateMethods(astMethods *ast.FieldList) (modelMethods []*model.Method) {
	for _, astMethod := range astMethods.List {
		modelMethods = append(modelMethods, g.generateMethod(astMethod))
	}
	return
}

func (g *modelGenerator) generateMethod(astMethod *ast.Field) *model.Method {
	in, out, variadic := g.generateSignature(astMethod.Type.(*ast.FuncType))
	return &model.Method{Name: astMethod.Names[0].Name, In: in, Variadic: variadic, Out: out}
}

func (g *modelGenerator) generateSignature(astFuncType *ast.FuncType) (in, out []*model.Parameter, variadic *model.Parameter) {
	in, variadic = g.generateInParams(astFuncType.Params)
	out = g.generateOutParams(astFuncType.Results)
	return
}

func (g *modelGenerator) generateInParams(params *ast.FieldList) (in []*model.Parameter, variadic *model.Parameter) {
	for _, param := range params.List {
		for _, name := range param.Names {
			in = append(in, g.generateParam(name.Name, param.Type))
		}
	}
	return
}

func (g *modelGenerator) generateOutParams(results *ast.FieldList) (out []*model.Parameter) {
	return
}

func (g *modelGenerator) generateParam(name string, typ ast.Expr) *model.Parameter {
	fmt.Println("Type:", g.info.TypeOf(typ))
	switch typedTyp := g.info.TypeOf(typ).(type) {
	case *types.Basic:
		if predeclared(typedTyp.Kind()) {
			return &model.Parameter{
				Name: name,
				Type: model.PredeclaredType(typedTyp.Name()),
			}
		} else {
			parts := strings.Split(typedTyp.Name(), ".")
			return &model.Parameter{
				Name: name,
				Type: &model.NamedType{
					Package: parts[0],
					Type:    parts[1],
				},
			}
		}
	// case *types.Pointer:
	// 	panic("implement")
	// case *types.Array:
	// 	panic("implement")
	// case *types.Slice:
	// 	panic("implement")
	// case *types.Map:
	// 	panic("implement")
	// case *types.Chan:
	// 	panic("implement")
	// case *types.Struct:
	// 	panic("implement")
	// case *types.Tuple:
	// 	panic("implement")
	// case *types.Signature:
	// 	panic("implement")
	// case *types.Named:
	// 	panic("implement")
	// case *types.Interface:
	// 	panic("implement")
	default:
		// parts := strings.Split(typedTyp.String(), ".")
		return &model.Parameter{
			Name: name,
			Type: &model.NamedType{
				// Package: parts[0],
				Type: "TODO",
			},
		}
	}
	// typeIdent, ok := typ.(*ast.Ident)
	// if ok {

	// 	return &model.Parameter{
	// 		Name: name,
	// 		Type: model.PredeclaredType(typeIdent.Name),

	// 		// &model.NamedType{
	// 		// 	/*Package: typeIdent.Obj,*/
	// 		// 	Type: typeIdent.Name,
	// 		// },
	// 	}
	// }
	// return &model.Parameter{Name: name}
}

func predeclared(basicKind types.BasicKind) bool {
	return basicKind >= types.Bool && basicKind <= types.String
}
