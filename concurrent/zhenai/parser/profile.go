package parser

import (
	"crawler.com/concurrent/engine"
	"crawler.com/concurrent/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://localhost:8080/mock/album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

// 已采用闭包方式从上级获取名字
var nameRe = regexp.MustCompile(
	`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)

var idReg = regexp.MustCompile(
	`http://localhost:8080/mock/album.zhenai.com/u/([\d]+)`)

func ParseProfile(content []byte, name string, url string) engine.ParseResult {

	profile := model.Profile{}

	profile.Age = extractInt(content, ageRe)
	profile.Height = extractInt(content, heightRe)
	profile.Weight = extractInt(content, weightRe)
	profile.Marriage = extractString(content, marriageRe)
	profile.Car = extractString(content, carRe)
	profile.Education = extractString(content, educationRe)
	profile.Gender = extractString(content, genderRe)
	profile.Hokou = extractString(content, hukouRe)
	profile.House = extractString(content, houseRe)
	profile.Income = extractString(content, incomeRe)
	profile.Occupation = extractString(content, occupationRe)
	profile.Name = name
	profile.Xinzuo = extractString(content, xinzuoRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      extractString([]byte(url), idReg),
				Type:    "zhenai",
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(content, -1)
	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        url,
				ParserFunc: ProfileParser(name),
			},
		)
	}

	return result
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, name, url)
	}
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func extractInt(content []byte, re *regexp.Regexp) int {
	n, err := strconv.Atoi(extractString(content, re))
	if err == nil {
		return n
	} else {
		return 0
	}
}
