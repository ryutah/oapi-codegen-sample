package interactor

import (
	"github.com/ryutah/oapi-codegen-sample/infrastructure/fake"
	"github.com/ryutah/oapi-codegen-sample/presentation/rest"
	"github.com/ryutah/oapi-codegen-sample/presentation/rest/oapi"
	"github.com/ryutah/oapi-codegen-sample/usecase"
)

func InjectServer() oapi.ServerInterface {
	return rest.NewServer(usecase.NewHello(fake.NewHelloRepository()))
}
