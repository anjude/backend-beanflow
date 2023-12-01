package interfaces

import "github.com/anjude/backend-beanflow/application/api_service"

func NewApiService() *ApiService {
	return &ApiService{
		user: api_service.NewUserController(),
		flow: api_service.NewFlowController(),
	}
}
