package knife

import (
	"github.com/gogf/gf/v2/net/ghttp"
	v3 "github.com/newton84/nn-knife/v3"
	css "github.com/newton84/nn-knife/v3/knife/webjars/css"

	fonts "github.com/newton84/nn-knife/v3/knife/webjars/fonts"

	icons "github.com/newton84/nn-knife/v3/knife/img/icons"

	img "github.com/newton84/nn-knife/v3/knife/webjars/img"

	js "github.com/newton84/nn-knife/v3/knife/webjars/js"

	knife "github.com/newton84/nn-knife/v3/knife"

	oauth "github.com/newton84/nn-knife/v3/knife/webjars/oauth"
)

func InitSwaggerKnife(s *ghttp.Server) {

	v3.AddApiDocRouter(s)

	v3.AddSwaggerResourcesRouter(s)

	knife.AddRouterOfDocHtml(s)

	icons.AddRouterOfAndroidChrome192x192Png(s)

	icons.AddRouterOfAndroidChrome512x512Png(s)

	icons.AddRouterOfAppleTouchIcon120x120Png(s)

	icons.AddRouterOfAppleTouchIcon152x152Png(s)

	icons.AddRouterOfAppleTouchIcon180x180Png(s)

	icons.AddRouterOfAppleTouchIcon60x60Png(s)

	icons.AddRouterOfAppleTouchIcon76x76Png(s)

	icons.AddRouterOfAppleTouchIconPng(s)

	icons.AddRouterOfFavicon16x16Png(s)

	icons.AddRouterOfFavicon32x32Png(s)

	icons.AddRouterOfMsapplicationIcon144x144Png(s)

	icons.AddRouterOfMstile150x150Png(s)

	icons.AddRouterOfSafariPinnedTabSvg(s)

	css.AddRouterOfAppAc23e017Css(s)

	css.AddRouterOfAppAc23e017CssGz(s)

	css.AddRouterOfChunk75464e7e8fb93ba5Css(s)

	css.AddRouterOfChunkD7d5f59cA9ffbfcbCss(s)

	css.AddRouterOfChunkVendorsF24a310aCss(s)

	css.AddRouterOfChunkVendorsF24a310aCssGz(s)

	fonts.AddRouterOfFontawesomeWebfont706450d7Ttf(s)

	fonts.AddRouterOfFontawesomeWebfont97493d3fWoff2(s)

	fonts.AddRouterOfFontawesomeWebfontD9ee23d5Woff(s)

	fonts.AddRouterOfFontawesomeWebfontF7c2b4b7Eot(s)

	fonts.AddRouterOfIconfont4ca3d0c0Ttf(s)

	fonts.AddRouterOfIconfontE2d2b98eEot(s)

	img.AddRouterOfEditormdLogo53ea80e2Svg(s)

	img.AddRouterOfFontawesomeWebfont29800836Svg(s)

	img.AddRouterOfIconfont1d48c203Svg(s)

	img.AddRouterOfLoadingC929501eGif(s)

	img.AddRouterOfLoading2x695405a9Gif(s)

	img.AddRouterOfLoading3x65eacf61Gif(s)

	js.AddRouterOfApp2fab4ac5Js(s)

	js.AddRouterOfApp2fab4ac5JsGz(s)

	js.AddRouterOfChunk069eb4372cfebf27Js(s)

	js.AddRouterOfChunk069eb4372cfebf27JsLICENSETxt(s)

	js.AddRouterOfChunk069eb4372cfebf27JsGz(s)

	js.AddRouterOfChunk0d102d5aB2bddffcJs(s)

	js.AddRouterOfChunk0d102d5aB2bddffcJsGz(s)

	js.AddRouterOfChunk0fd67716D57e2c41Js(s)

	js.AddRouterOfChunk0fd67716D57e2c41JsGz(s)

	js.AddRouterOfChunk260d712a390177feJs(s)

	js.AddRouterOfChunk260d712a390177feJsLICENSETxt(s)

	js.AddRouterOfChunk260d712a390177feJsGz(s)

	js.AddRouterOfChunk2d0af44e392afcd6Js(s)

	js.AddRouterOfChunk2d0bd799Eb48b7f1Js(s)

	js.AddRouterOfChunk2d0da532591ad7fcJs(s)

	js.AddRouterOfChunk3b888a658737ce4fJs(s)

	js.AddRouterOfChunk3b888a658737ce4fJsGz(s)

	js.AddRouterOfChunk3ec4aaa8A79d19f8Js(s)

	js.AddRouterOfChunk3ec4aaa8A79d19f8JsGz(s)

	js.AddRouterOfChunk589faee05bfd1708Js(s)

	js.AddRouterOfChunk589faee05bfd1708JsLICENSETxt(s)

	js.AddRouterOfChunk589faee05bfd1708JsGz(s)

	js.AddRouterOfChunk735c675c5b409314Js(s)

	js.AddRouterOfChunk735c675c5b409314JsGz(s)

	js.AddRouterOfChunk75464e7eB130271bJs(s)

	js.AddRouterOfChunk75464e7eB130271bJsGz(s)

	js.AddRouterOfChunkAdb9e9442c7f24feJs(s)

	js.AddRouterOfChunkAdb9e9442c7f24feJsLICENSETxt(s)

	js.AddRouterOfChunkAdb9e9442c7f24feJsGz(s)

	js.AddRouterOfChunkD7d5f59cE61130f3Js(s)

	js.AddRouterOfChunkD7d5f59cE61130f3JsLICENSETxt(s)

	js.AddRouterOfChunkD7d5f59cE61130f3JsGz(s)

	js.AddRouterOfChunkVendorsD51cf6f8Js(s)

	js.AddRouterOfChunkVendorsD51cf6f8JsLICENSETxt(s)

	js.AddRouterOfChunkVendorsD51cf6f8JsGz(s)

	oauth.AddRouterOfAxiosMinJs(s)

	oauth.AddRouterOfOauth2Html(s)

}
