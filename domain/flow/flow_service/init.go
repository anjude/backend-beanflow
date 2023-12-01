package flow_service

import (
	"github.com/anjude/backend-beanflow/infrastructure/repository/flow_repo"
)

func NewFlowService() *FlowService {
	return &FlowService{
		flowRepo: flow_repo.NewFlowRepo(),
	}
}
