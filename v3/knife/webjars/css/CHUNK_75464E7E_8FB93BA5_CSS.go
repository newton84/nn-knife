package css

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/newton84/nn-knife/constant"
	"github.com/newton84/nn-knife/utils"
)

const (
	CHUNK_75464E7E_8FB93BA5_CSS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/css/chunk-75464e7e.8fb93ba5.css"
	// 文件内容的16进制表示
	CHUNK_75464E7E_8FB93BA5_CSS_HEX_CONTENT = `.api-title[data-v-b092cbdc]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}`
)

func AddRouterOfChunk75464e7e8fb93ba5Css(s *ghttp.Server) {

	utils.GetCss(s, CHUNK_75464E7E_8FB93BA5_CSS_RELATIVE_PATH, CHUNK_75464E7E_8FB93BA5_CSS_HEX_CONTENT)

}
