package v3

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/newton84/nn-knife/constant"
	"log"
	"net/http"
	"strconv"
)

const (
	// TODO 路径要改
	API_DOCS_RELATIVE_PATH = constant.ROOT_PATH + "/v3/api-docs"
)

func AddApiDocRouter(s *ghttp.Server) {

	s.BindHandler(API_DOCS_RELATIVE_PATH, func(r *ghttp.Request) {

		if rs, err := json.Marshal(s.GetOpenApi()); err != nil {
			panic(gerror.Wrap(err, `WriteJson failed`))
		} else {

			r.Response.Status = http.StatusOK
			r.Header.Add("content-type", "text/html;charset=UTF-8")
			r.Header.Add("content-length", strconv.Itoa(len(rs)))
			r.Header.Add("connection", "keep-alive")
			_, err := r.Response.BufferWriter.Write(rs)
			if nil != err {
				log.Fatal(err)
				return
			}
		}

	})

}
