package custom

import "github.com/khulnasoft-labs/defsec/pkg/terraform"

type customCheckVariables map[string]string

type customContext struct {
	module    *terraform.Module
	variables customCheckVariables
}

func NewEmptyCustomContext() *customContext {
	return &customContext{
		module:    nil,
		variables: make(customCheckVariables),
	}
}

func NewCustomContext(module *terraform.Module) *customContext {
	return &customContext{
		module:    module,
		variables: make(customCheckVariables),
	}
}

func NewCustomContextWithVariables(module *terraform.Module, variables customCheckVariables) *customContext {
	return &customContext{
		module:    module,
		variables: variables,
	}
}
