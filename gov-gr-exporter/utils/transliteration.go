package utils

import (
	"regexp"
)

type TransliterationRules struct {
	Regex *regexp.Regexp
	Replacement string
}

var GreekToLatinTransliterationRules = []TransliterationRules{
	{regexp.MustCompile("ΕΥ"), "EF"},
	{regexp.MustCompile("ΑΥ"), "AV"},
	{regexp.MustCompile("Ϊ"), "I"},
	{regexp.MustCompile("ΟΥ"), "OU"},
	{regexp.MustCompile("Α"), "A"},
	{regexp.MustCompile("Β"), "B"},
	{regexp.MustCompile("Γ"), "G"},
	{regexp.MustCompile("Δ"), "D"},
	{regexp.MustCompile("Ε"), "E"},
	{regexp.MustCompile("Ζ"), "Z"},
	{regexp.MustCompile("Η"), "I"},
	{regexp.MustCompile("Θ"), "TH"},
	{regexp.MustCompile("Ι"), "I"},
	{regexp.MustCompile("Κ"), "K"},
	{regexp.MustCompile("Λ"), "L"},
	{regexp.MustCompile("Μ"), "M"},
	{regexp.MustCompile("Ν"), "N"},
	{regexp.MustCompile("Ξ"), "KS"},
	{regexp.MustCompile("Ο"), "O"},
	{regexp.MustCompile("Π"), "P"},
	{regexp.MustCompile("Ρ"), "R"},
	{regexp.MustCompile("Σ"), "S"},
	{regexp.MustCompile("Τ"), "T"},
	{regexp.MustCompile("Υ"), "Y"},
	{regexp.MustCompile("Φ"), "F"},
	{regexp.MustCompile("Χ"), "X"},
	{regexp.MustCompile("Ψ"), "PS"},
	{regexp.MustCompile("Ω"), "O"},
}

func GreekToLatin(text string) string {
	for _, rule := range GreekToLatinTransliterationRules {
		text = rule.Regex.ReplaceAllString(text, rule.Replacement)
	}
	return text
}