package xomologisou

import (
	"regexp"
	"strings"
)

const trimCharacters = " \t\r\n"

//Newlines and messages
var newLine = regexp.MustCompile(`([\[{]n[]}])`)
var message = regexp.MustCompile(`([\[{]m[]}])`)
var space = regexp.MustCompile(`([\[{]s[]}])`)

var removeParenthesisStart = regexp.MustCompile(`(\[\s*\(\s*])(\[\s*[^)]\s*])*`)
var removeParenthesisEnd = regexp.MustCompile(`(\[\s*[^(]\s*])+?(\[\s*\)\s*])`)
var removeParenthesis = regexp.MustCompile(`(\[\s*\(\s*])(\[\s*\)\s*])`)
var removeExtraDashes = regexp.MustCompile(`(\[\s*-\s*]){2,}`)
var removeFirstOptional = regexp.MustCompile(`^(\[[^(]]*])+`)
var removeTrailingOptional = regexp.MustCompile(`(\[[^)]*])+$`)
var escapeOptionals = regexp.MustCompile(`\[(.+?)]`)

var hasSpace = regexp.MustCompile(`^\s`)

func (p *Publishable) BuildFormats(c *Carrier, similar int) {

	if p.NoFormat(c) || c.Posting == nil {
		p.Content = strings.Trim(p.Content, trimCharacters)
		p.FinalForm = strings.Trim(p.Content, trimCharacters)
		return
	}

	text := c.Posting.Format

	for k, v := range p.Options {
		if set, ok := c.Form.OptionSets[k]; ok {
			display, ok := set.Options[v]
			if !ok && !set.AllowCustom {
				continue
			}
			val := v
			if ok && display != "" {
				val = display
			}
			text = strings.Replace(text, "{"+k+"}", val, -1)
		}
	}
	for k := range c.Form.OptionSets {
		text = strings.Replace(text, "{"+k+"}", "", -1)
	}
	if !message.MatchString(text) {
		text += "{m}"
	}

	text = removeExtraDashes.ReplaceAllString(text, "$1")
	text = removeParenthesisStart.ReplaceAllString(text, "$1")
	text = removeParenthesisEnd.ReplaceAllString(text, "$2")
	text = removeParenthesis.ReplaceAllString(text, "")
	text = removeFirstOptional.ReplaceAllString(text, "")
	text = removeTrailingOptional.ReplaceAllString(text, "")
	text = newLine.ReplaceAllString(text, "\n")
	text = space.ReplaceAllString(text, " ")
	text = escapeOptionals.ReplaceAllString(text, "$1")

	p.Properties = strings.Trim(message.ReplaceAllString(text, ""), trimCharacters)
	text = message.ReplaceAllString(text, strings.Trim(p.Content, trimCharacters))

	if !hasSpace.MatchString(text) {
		text = " " + text
	}
	p.FinalForm = "#" + c.NextHashtag() + text
	p.FinalForm = strings.Trim(p.FinalForm, trimCharacters)
}
