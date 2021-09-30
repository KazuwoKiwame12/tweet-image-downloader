package utility

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"tweet-image-downloader/entity"
)

type paramBuilder struct {
	url.Values
}

func NewParamBuilder() *paramBuilder {
	p := paramBuilder{}
	p.Values = url.Values{}
	return &p
}

func (p *paramBuilder) Expansions(fields entity.ExpansionFields) *paramBuilder {
	p.Add("expansions", strings.Join(fields.ToStringSlice(), ","))
	return p
}

func (p *paramBuilder) MaxResults(val int) *paramBuilder {
	p.Add("max_results", strconv.Itoa(val))
	return p
}

func (p *paramBuilder) MediaFields(fields entity.MediaFields) *paramBuilder {
	p.Add("media.fields", strings.Join(fields.ToStringSlice(), ","))
	return p
}

func (p *paramBuilder) Query(fields entity.QueryFields) *paramBuilder {
	// 複数の演算子を一つのクエリにまとめるには、
	// コンマ(",")ではなくANDを意味する空白(" ")を用いる
	p.Add("query", strings.Join(fields.ToStringSlice(), " "))
	return p
}

func (p *paramBuilder) TweetFields(fields entity.TweetFields) *paramBuilder {
	p.Add("tweet.fields", strings.Join(fields.ToStringSlice(), ","))
	return p
}

func (p *paramBuilder) Build() (string, error) {
	encoded, err := p.validateEmptyParamValue()
	return encoded, err
}

func (p *paramBuilder) validateEmptyParamValue() (string, error) {
	e := p.Encode()
	r := regexp.MustCompile(`(?i)(expansions|max_results|media.fields|query|tweet.fields)=&`)
	if r.MatchString(e) {
		return e, entity.ErrorEmptyParameterValue
	}
	return e, nil
}
