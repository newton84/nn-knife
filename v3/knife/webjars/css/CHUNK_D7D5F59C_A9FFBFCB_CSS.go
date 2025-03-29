package css

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/newton84/nn-knife/constant"
	"github.com/newton84/nn-knife/utils"
)

const (
	CHUNK_D7D5F59C_A9FFBFCB_CSS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/css/chunk-d7d5f59c.a9ffbfcb.css"
	// 文件内容的16进制表示
	CHUNK_D7D5F59C_A9FFBFCB_CSS_HEX_CONTENT = `.api-tab[data-v-3e1ec994]{margin-top:15px}.api-tab .ant-tag[data-v-3e1ec994]{height:32px;line-height:32px}.api-basic[data-v-3e1ec994]{padding:11px}.api-basic-title[data-v-3e1ec994]{font-size:14px;font-weight:700}.api-basic-body[data-v-3e1ec994]{font-size:14px;font-family:-webkit-body}.api-description[data-v-3e1ec994]{border-left:4px solid #ddd;line-height:30px}.api-body-desc[data-v-3e1ec994]{padding:10px;min-height:35px;-webkit-box-sizing:border-box;box-sizing:border-box;border:1px solid #e8e8e8}.ant-card-body[data-v-3e1ec994]{padding:5px}.api-title[data-v-3e1ec994]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}.content-line[data-v-3e1ec994]{height:25px;line-height:25px}.content-line-count[data-v-3e1ec994]{height:35px;line-height:35px}.divider[data-v-3e1ec994]{margin:4px 0}`
)

func AddRouterOfChunkD7d5f59cA9ffbfcbCss(s *ghttp.Server) {

	utils.GetCss(s, CHUNK_D7D5F59C_A9FFBFCB_CSS_RELATIVE_PATH, CHUNK_D7D5F59C_A9FFBFCB_CSS_HEX_CONTENT)

}
