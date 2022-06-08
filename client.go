package genshin

import (
	"log"
	"strconv"

	"github.com/rushkii/genshin-go/utils"
	"github.com/valyala/fasthttp"
)

type Client struct {
	Ltuid  string
	Ltoken string
	IsCn   bool
	config Config
}

func NewClient(ltuid int, ltoken string, iscn bool) *Client {
	return &Client{
		Ltuid:  strconv.Itoa(ltuid),
		Ltoken: ltoken,
		IsCn:   iscn,
	}
}

func (c *Client) GET(endpoint string) []byte {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetCookie("ltuid", c.Ltuid)
	req.Header.SetCookie("ltoken", c.Ltoken)
	req.Header.Set("ds", utils.GenerateDs())
	req.Header.Set("x-rpc-app_version", "1.5.0")
	req.Header.Set("x-rpc-client_type", "5")
	req.Header.Set("x-rpc-language", "en-us")
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.SetRequestURI("https://bbs-api-os.hoyoverse.com/game_record" + endpoint)

	if err := fasthttp.Do(req, res); err != nil {
		log.Panicln("Request Err:", err)
	}

	return res.Body()
}

func (c *Client) POST(endpoint string, data []byte) []byte {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetBody(data)
	req.Header.SetMethod("POST")
	req.Header.SetCookie("ltuid", c.Ltuid)
	req.Header.SetCookie("ltoken", c.Ltoken)
	req.Header.SetContentType("application/json")
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.SetRequestURI("https://bbs-api-os.hoyoverse.com/game_record" + endpoint)
	req.Header.Set("ds", utils.GenerateDs())
	req.Header.Set("x-rpc-app_version", "1.5.0")
	req.Header.Set("x-rpc-client_type", "5")
	req.Header.Set("x-rpc-language", "en-us")

	if err := fasthttp.Do(req, res); err != nil {
		log.Panicln("Request Err:", err)
	}

	return res.Body()
}

func EnkaReq(uid string) []byte {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetContentType("application/json")
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.SetRequestURI("https://enka.shinshin.moe/u/" + uid + "/__data.json")

	if err := fasthttp.Do(req, res); err != nil {
		log.Panicln("Request Err:", err)
	}

	return res.Body()
}
