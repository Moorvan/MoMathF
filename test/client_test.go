package test

import (
	"MoMathF/MathFClient"
	"MoMathF/core"
	"MoMathF/global"
	mlog "MoMathF/mo-log"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"testing"
)

var (
	log = mlog.Log
)
var (
	apiId      string
	apiKey     string
	client     *resty.Client
	mathClient *MathFClient.MoMathFClient
)

func init() {
	core.Viper("../config.yaml")
	apiId = global.GB_CONFIG.Api.ApiId
	apiKey = global.GB_CONFIG.Api.ApiKey
	client = resty.New()
	client.SetHeaders(map[string]string{
		"app_id":  apiId,
		"app_key": apiKey,
	})
	mathClient = core.NewClient()
}

func TestConnect(t *testing.T) {
	api := "https://api.mathpix.com/v3/text"
	req := &struct {
		Src         string   `json:"src"`
		Formats     []string `json:"formats"`
		DataOptions struct {
			IncludeAsciimath string `json:"include_asciimath"`
		} `json:"data_options"`
	}{
		Src:     "https://mathpix-ocr-examples.s3.amazonaws.com/cases_hw.jpg",
		Formats: []string{"text", "data"},
		DataOptions: struct {
			IncludeAsciimath string `json:"include_asciimath"`
		}{IncludeAsciimath: "true"},
	}

	response := &struct {
		Data []struct {
			Type  string
			Value string
		}
	}{}

	res, err := client.R().SetBody(req).SetResult(response).Post(api)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	log.PrintStruct(response)
}

func TestApiFile(t *testing.T) {
	api := "https://api.mathpix.com/v3/text"
	req := &struct {
		MathInlineDelimiters []string `json:"math_inline_delimiters"`
		RmSpaces             string   `json:"rm_spaces"`
		Formats              []string `json:"formats"`
	}{
		MathInlineDelimiters: []string{"$$", "$$"},
		RmSpaces:             "True",
		Formats:              []string{"text", "html"},
	}
	response := &struct {
		LatexStyled string `json:"latex_styled"`
		Text        string `json:"text"`
	}{}

	jn, _ := json.Marshal(req)
	log.Println(string(jn))

	res, err := client.R().SetFormData(map[string]string{
		"options_json": string(jn),
	}).SetFile("file", "../testPics/2.png").SetResult(response).Post(api)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	log.PrintStruct(response)
}

func TestClient(t *testing.T) {
	res, err := mathClient.GetLatexFromPic("../testPics/2.png")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
