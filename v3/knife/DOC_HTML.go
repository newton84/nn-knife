package knife

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/newton84/nn-knife/constant"
	"github.com/newton84/nn-knife/utils"
)

const (
	DOC_HTML_RELATIVE_PATH = constant.ROOT_PATH + "/doc.html"
	// 文件内容的16进制表示
	DOC_HTML_HEX_CONTENT = `<!DOCTYPE html><html lang=en><head><meta charset=utf-8><meta http-equiv=X-UA-Compatible content="IE=edge"><meta name=viewport content="width=device-width,initial-scale=1"><link rel=icon href=favicon.ico><title></title><link href=webjars/css/chunk-75464e7e.8fb93ba5.css rel=prefetch><link href=webjars/css/chunk-d7d5f59c.a9ffbfcb.css rel=prefetch><link href=webjars/js/chunk-069eb437.2cfebf27.js rel=prefetch><link href=webjars/js/chunk-0d102d5a.b2bddffc.js rel=prefetch><link href=webjars/js/chunk-0fd67716.d57e2c41.js rel=prefetch><link href=webjars/js/chunk-260d712a.390177fe.js rel=prefetch><link href=webjars/js/chunk-2d0af44e.392afcd6.js rel=prefetch><link href=webjars/js/chunk-2d0bd799.eb48b7f1.js rel=prefetch><link href=webjars/js/chunk-2d0da532.591ad7fc.js rel=prefetch><link href=webjars/js/chunk-3b888a65.8737ce4f.js rel=prefetch><link href=webjars/js/chunk-3ec4aaa8.a79d19f8.js rel=prefetch><link href=webjars/js/chunk-589faee0.5bfd1708.js rel=prefetch><link href=webjars/js/chunk-735c675c.5b409314.js rel=prefetch><link href=webjars/js/chunk-75464e7e.b130271b.js rel=prefetch><link href=webjars/js/chunk-adb9e944.2c7f24fe.js rel=prefetch><link href=webjars/js/chunk-d7d5f59c.e61130f3.js rel=prefetch><link href=webjars/css/app.ac23e017.css rel=preload as=style><link href=webjars/css/chunk-vendors.f24a310a.css rel=preload as=style><link href=webjars/js/app.2fab4ac5.js rel=preload as=script><link href=webjars/js/chunk-vendors.d51cf6f8.js rel=preload as=script><link href=webjars/css/chunk-vendors.f24a310a.css rel=stylesheet><link href=webjars/css/app.ac23e017.css rel=stylesheet></head><body><noscript><strong>We're sorry but knife4j-vue doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id=app></div><script src=webjars/js/chunk-vendors.d51cf6f8.js></script><script src=webjars/js/app.2fab4ac5.js></script></body></html>`
)

func AddRouterOfDocHtml(s *ghttp.Server) {

	utils.GetHtml(s, DOC_HTML_RELATIVE_PATH, DOC_HTML_HEX_CONTENT)

}
