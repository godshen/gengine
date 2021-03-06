package base

import (
	"gengine/context"
)

type Args struct {
	ArgList          []*Arg
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (as *Args) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	as.knowledgeContext = kc
	as.dataCtx = dc

	if as.ArgList != nil {
		for _, val := range as.ArgList {
			val.Initialize(kc, dc)
		}
	}
}

func (as *Args) AcceptFunctionCall(funcCall *FunctionCall) error {
	holder := &Arg{
		FunctionCall: funcCall,
	}
	as.ArgList = append(as.ArgList, holder)
	return nil
}

func (as *Args) AcceptMethodCall(methodCall *MethodCall) error {
	holder := &Arg{
		MethodCall: methodCall,
	}
	as.ArgList = append(as.ArgList, holder)
	return nil
}

func (as *Args) AcceptVariable(name string) error {
	holder := &Arg{
		Variable: name,
	}
	as.ArgList = append(as.ArgList, holder)
	return nil
}

func (as *Args) AcceptConstant(cons *Constant) error {
	holder := &Arg{
		Constant: cons,
	}
	as.ArgList = append(as.ArgList, holder)
	return nil
}

func (as *Args) AcceptMapVar(mapVar *MapVar) error {
	holder := &Arg{
		MapVar:           mapVar,
	}
	as.ArgList = append(as.ArgList, holder)
	return nil
}

func (as *Args) Evaluate(Vars map[string]interface{}) ([]interface{}, error) {
	if as.ArgList == nil || len(as.ArgList) == 0 {
		return make([]interface{}, 0), nil
	}
	retVal := make([]interface{}, len(as.ArgList))
	for i, v := range as.ArgList {
		rv, err := v.Evaluate(Vars)
		if err != nil {
			return retVal, err
		}
		retVal[i] = rv
	}
	return retVal, nil
}

