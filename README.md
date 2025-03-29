# nn-knife

#### 介绍

goframe 不想使用那个又丑又不好用的waggerUI 就可能使用这个，是基于https://gitee.com/xiaoym/knife4j 项目，适配Goframe框架

使用方法
 
func enhanceOpenAPIDoc(s *ghttp.Server) {
    openapi := s.GetOpenApi()
    openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
    openapi.Config.CommonResponseDataField = `Data`

    openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "GoFrame",
			URL:  "https://goframe.org",
		},
	}
	nnknife.InitSwaggerKnife(s)

}

http://ip:port/doc.html

<img src="imgs\img.png"/>

