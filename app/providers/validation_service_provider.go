package providers

import (
	"github.com/chenyuIT/framework/contracts/validation"
	"github.com/chenyuIT/framework/facades"
)

type ValidationServiceProvider struct {
}

func (receiver *ValidationServiceProvider) Register() {

}

func (receiver *ValidationServiceProvider) Boot() {
	if err := facades.Validation.AddRules(receiver.rules()); err != nil {
		facades.Log.Errorf("add rules error: %+v", err)
	}
}

func (receiver *ValidationServiceProvider) rules() []validation.Rule {
	return []validation.Rule{}
}
