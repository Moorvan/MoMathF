package MathFClient

import (
	"encoding/json"
	"errors"
	mlog "github.com/Moorvan/MoMathF/mo-log"
	"github.com/go-resty/resty/v2"
)

var (
	log = mlog.Log
)

type MoMathFClient struct {
	ApiConfig *APIConfig
	client    *resty.Client
}

func NewMoMathFClient(conf *APIConfig) (*MoMathFClient, error) {
	client := resty.New()
	if conf == nil {
		return nil, errors.New("config is nil")
	}
	client.SetHeaders(map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.5 Safari/605.1.15",
		"app_id":     conf.ApiId,
		"app_key":    conf.ApiKey,
	})
	MoMath := &MoMathFClient{
		ApiConfig: conf,
		client:    client,
	}
	return MoMath, nil
}

func (client *MoMathFClient) GetLatexFromPic(path string) (string, error) {
	req := client.client.R()
	api := "https://api.mathpix.com/v3/text"
	req.SetFile("file", path)
	options := &struct {
		MathInlineDelimiters []string `json:"math_inline_delimiters"`
		RmSpaces             bool     `json:"rm_spaces,string"`
	}{
		MathInlineDelimiters: []string{"", ""},
		RmSpaces:             true,
	}
	optionsStr, _ := json.Marshal(options)
	req.SetFormData(map[string]string{
		"options_json": string(optionsStr),
	})
	response := &struct {
		Text string `json:"text"`
	}{}
	req.SetResult(response)
	resp, err := req.Post(api)
	if err != nil {
		log.Errorln("Request for latex from pic error:", err)
		return "", err
	}
	log.Println(resp.String())
	if response.Text == "" {
		log.Errorln("Response latex is empty")
		return "", errors.New("latex is empty")
	}
	return response.Text, nil
}

func (client *MoMathFClient) GetLatexFromPicUrl(url string) (string, error) {
	req := client.client.R()
	api := "https://api.mathpix.com/v3/text"
	request := &struct {
		Src                  string   `json:"src"`
		MathInlineDelimiters []string `json:"math_inline_delimiters"`
		RmSpaces             bool     `json:"rm_spaces,string"`
	}{
		Src:                  url,
		MathInlineDelimiters: []string{"", ""},
		RmSpaces:             true,
	}
	req.SetBody(request)
	response := &struct {
		Text string `json:"text"`
	}{}
	req.SetResult(response)
	resp, err := req.Post(api)
	if err != nil {
		log.Errorln("Request for latex from pic error:", err)
		return "", err
	}
	log.Println(resp.String())
	if response.Text == "" {
		log.Errorln("Response latex is empty")
		return "", errors.New("latex is empty")
	}
	return response.Text, nil
}
