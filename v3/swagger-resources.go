package v3

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/newton84/nn-knife/constant"
	"github.com/newton84/nn-knife/utils"
)

const (
	// TODO 路径要改
	SWAGGER_RESOURCES_CONTENT     = `{"configUrl": "` + constant.ROOT_PATH + `/v3/api-docs/swagger-config","oauth2RedirectUrl": "` + constant.ROOT_PATH + `/swagger-ui/oauth2-redirect.html","url": "` + constant.ROOT_PATH + `/v3/api-docs","validatorUrl": ""}`
	SWAGGER_RESOURCES_CONFIG_PATH = constant.ROOT_PATH + "/v3/api-docs/swagger-config"
)

func AddSwaggerResourcesRouter(s *ghttp.Server) {
	utils.GetJson(s, SWAGGER_RESOURCES_CONFIG_PATH, SWAGGER_RESOURCES_CONTENT)
}
