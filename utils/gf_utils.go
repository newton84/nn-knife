package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
	"net/http"
	"strconv"
)

func GetHtml(s *ghttp.Server, relativePath string, content string) {
	rs := []byte(content)
	s.BindHandler(relativePath, func(r *ghttp.Request) {
		r.Response.Status = http.StatusOK
		r.Response.Header().Add("content-type", "text/html;charset=UTF-8")
		r.Response.Header().Add("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Add("connection", "keep-alive")
		_, err := r.Response.BufferWriter.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}

	})

}

func GetJson(s *ghttp.Server, relativePath string, json string) {
	s.BindHandler(relativePath, func(r *ghttp.Request) {
		rs := []byte(json)
		r.Response.Status = http.StatusOK
		r.Response.Header().Add("content-type", "text/json;charset=UTF-8")
		r.Response.Header().Add("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Add("connection", "keep-alive")
		_, err := r.Response.BufferWriter.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}
	})
}

func GetJs(s *ghttp.Server, relativePath string, hexContent string) {
	rs, err := hex.DecodeString(hexContent)
	if nil != err {
		fmt.Println("err:", err)
		return
	}
	s.BindHandler(relativePath, func(r *ghttp.Request) {

		r.Response.Status = http.StatusOK
		r.Response.Header().Add("content-type", "text/javascript;charset=UTF-8")
		r.Response.Header().Add("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Add("connection", "keep-alive")
		_, err := r.Response.BufferWriter.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}
	})

}

func GetCss(s *ghttp.Server, relativePath string, content string) {

	s.BindHandler(relativePath, func(r *ghttp.Request) {
		rs := []byte(content)
		r.Response.Status = http.StatusOK
		r.Response.Header().Add("content-type", "text/css;charset=UTF-8")
		r.Response.Header().Add("content-length", strconv.Itoa(len(rs)))
		r.Header.Add("connection", "keep-alive")
		_, err := r.Response.BufferWriter.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}
		r.Response.BufferWriter.Flush()
	})

}

func GetOther(s *ghttp.Server, relativePath string, hexContent string) {

	rs, err := hex.DecodeString(hexContent)
	if nil != err {
		fmt.Println("err:", err)
		return
	}
	s.BindHandler(relativePath, func(r *ghttp.Request) {

		r.Response.Status = http.StatusOK
		r.Response.Header().Add("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Add("connection", "keep-alive")
		_, err := r.Response.BufferWriter.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}
	})

}
