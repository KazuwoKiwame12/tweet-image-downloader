package utility

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"tweet-image-downloader/entity"
)

type ParamBuilder struct {
	url.Values
}

func NewParamBuilder() *ParamBuilder {
	p := ParamBuilder{}
	p.Values = url.Values{}
	return &p
}

func (p *ParamBuilder) Expansions(fields entity.ExpansionFields) *ParamBuilder {
	p.Add("expansions", strings.Join(fields.ToStringSlice(), ","))
	return p
}

func (p *ParamBuilder) MaxResults(val int) *ParamBuilder {
	p.Add("max_results", strconv.Itoa(val))
	return p
}

func (p *ParamBuilder) MediaFields(fields entity.MediaFields) *ParamBuilder {
	p.Add("media.fields", strings.Join(fields.ToStringSlice(), ","))
	return p
}

func (p *ParamBuilder) Query(fields entity.QueryFields) *ParamBuilder {
	// 複数の演算子を一つのクエリにまとめるには、
	// コンマ(",")ではなくANDを意味する空白(" ")を用いる
	p.Add("query", strings.Join(fields.ToStringSlice(), " "))
	return p
}

func (p *ParamBuilder) TweetFields(fields entity.TweetFields) *ParamBuilder {
	p.Add("tweet.fields", strings.Join(fields.ToStringSlice(), ","))
	return p
}

func (p *ParamBuilder) Build() (string, error) {
	encoded, err := p.validateEmptyParamValue()
	return encoded, err
}

func (p *ParamBuilder) validateEmptyParamValue() (string, error) {
	e := p.Encode()
	r := regexp.MustCompile(`(?i)(expansions|media.fields|query|tweet.fields)=(&|$)`)
	if r.MatchString(e) {
		return e, entity.ErrorEmptyParameterValue
	}
	return e, nil
}
