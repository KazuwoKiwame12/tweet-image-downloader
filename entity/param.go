package entity

import (
	"net/url"
	"strconv"
)

type paramBuilder struct {
	url.Values
}

func NewParamBuilder() *paramBuilder {
	p := paramBuilder{}
	return &p
}

func (p *paramBuilder) Expansions(fields []ExpansionField) *paramBuilder {
	for _, field := range fields {
		p.Add("expansions", field.String())
	}
	return p
}

func (p *paramBuilder) MaxResults(val int) *paramBuilder {
	p.Add("max_results", strconv.Itoa(val))
	return p
}

func (p *paramBuilder) MediaFields(fields []MediaField) *paramBuilder {
	for _, field := range fields {
		p.Add("media.fields", field.String())
	}
	return p
}

func (p *paramBuilder) Query(fields []QueryField) *paramBuilder {
	var val string = ""
	// 複数の演算子を一つのクエリにまとめるには、
	// コンマ(",")ではなくANDを意味する空白(" ")を用いる
	for i, field := range fields {
		val += field.String()
		if i != len(fields)-1 { // 最後に空白はいらない
			// urlのパラメータでは空白を含む場合は"+"か"%20"を利用
			// 今回は"+"を利用
			val += "+"
		}
	}
	p.Add("query", val)
	return p
}

func (p *paramBuilder) TweetFields(fields []TweetField) *paramBuilder {
	for _, field := range fields {
		p.Add("tweet.fields", field.String())
	}
	return p
}

func (p *paramBuilder) String() string {
	return p.Encode()
}
