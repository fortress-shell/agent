package parser

import (
	"github.com/fortress-shell/agent/steps"
)

type GeneratedSteps struct {
	Checkout      []steps.Step
	OverrideBuild []steps.Step
}

func GenerateSteps(p *Payload) *GeneratedSteps {
	tasks := &GeneratedSteps{}
	checkout := &steps.OverrideCheckoutStep{
		Environment: p.Machine.Environment,
	}
	tasks.Checkout = append(tasks.Checkout, checkout)
	for _, v := range p.Test.Override {
		build := &steps.OverrideBuildStep{
			Command:     v,
			Environment: p.Machine.Environment,
		}
		tasks.OverrideBuild = append(tasks.OverrideBuild, build)
	}
	return tasks
}
