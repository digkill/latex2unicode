package latex2unicode

import (
	"regexp"
	"strings"
)

var latexToUnicode = map[string]string{
	`alpha`: "α", `beta`: "β", `gamma`: "γ", `delta`: "δ",
	`epsilon`: "ε", `theta`: "θ", `lambda`: "λ", `pi`: "π",
	`phi`: "φ", `sigma`: "σ", `omega`: "ω", `times`: "×",
	`cdot`: "⋅", `leq`: "≤", `geq`: "≥", `neq`: "≠",
	`pm`: "±", `sqrt`: "√", `int`: "∫", `sum`: "∑",
	`prod`: "∏", `infty`: "∞", `rightarrow`: "→",
	`leftarrow`: "←", `uparrow`: "↑", `downarrow`: "↓",
}

var superscripts = map[rune]rune{
	'0': '⁰', '1': '¹', '2': '²', '3': '³', '4': '⁴',
	'5': '⁵', '6': '⁶', '7': '⁷', '8': '⁸', '9': '⁹',
	'+': '⁺', '-': '⁻', '=': '⁼', '(': '⁽', ')': '⁾',
}

var subscripts = map[rune]rune{
	'0': '₀', '1': '₁', '2': '₂', '3': '₃', '4': '₄',
	'5': '₅', '6': '₆', '7': '₇', '8': '₈', '9': '₉',
}

var fractions = map[string]string{
	"1/2": "½", "1/3": "⅓", "2/3": "⅔", "1/4": "¼", "3/4": "¾",
	"1/5": "⅕", "2/5": "⅖", "3/5": "⅗", "4/5": "⅘", "1/6": "⅙",
	"5/6": "⅚", "1/8": "⅛", "3/8": "⅜", "5/8": "⅝", "7/8": "⅞",
}

func convertSuperscripts(match string) string {
	var result strings.Builder
	for _, char := range match {
		if val, exists := superscripts[char]; exists {
			result.WriteRune(val)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func convertSubscripts(match string) string {
	var result strings.Builder
	for _, char := range match {
		if val, exists := subscripts[char]; exists {
			result.WriteRune(val)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func ConvertLatexToUnicode(latex string) string {

	for latexSymbol, unicodeSymbol := range latexToUnicode {
		latex = strings.ReplaceAll(latex, `\`+latexSymbol, unicodeSymbol)
	}

	reSuperscript := regexp.MustCompile(`(\w+)\^(\d+)`)
	latex = reSuperscript.ReplaceAllStringFunc(latex, func(match string) string {
		parts := reSuperscript.FindStringSubmatch(match)
		return parts[1] + convertSuperscripts(parts[2])
	})

	reSubscript := regexp.MustCompile(`(\w+)_([0-9]+)`)
	latex = reSubscript.ReplaceAllStringFunc(latex, func(match string) string {
		parts := reSubscript.FindStringSubmatch(match)
		return parts[1] + convertSubscripts(parts[2])
	})

	reSqrt := regexp.MustCompile(`\\sqrt{([^}]+)}`)
	latex = reSqrt.ReplaceAllString(latex, "√$1")

	reFrac := regexp.MustCompile(`\\frac{(\d+)}{(\d+)}`)
	latex = reFrac.ReplaceAllStringFunc(latex, func(match string) string {
		parts := reFrac.FindStringSubmatch(match)
		if fraction, exists := fractions[parts[1]+"/"+parts[2]]; exists {
			return fraction
		}
		return parts[1] + "⁄" + parts[2]
	})

	reNestedFrac := regexp.MustCompile(`\\frac{(\\frac{[^}]+})}{(\d+)}`)
	latex = reNestedFrac.ReplaceAllString(latex, "($1)/$2")

	reOverline := regexp.MustCompile(`\\overline{([^}]+)}`)
	latex = reOverline.ReplaceAllString(latex, "$1̅")

	reText := regexp.MustCompile(`\\text{([^}]+)}`)
	latex = reText.ReplaceAllString(latex, "$1")

	latex = strings.ReplaceAll(latex, `\left(`, "(")
	latex = strings.ReplaceAll(latex, `\right)`, ")")

	return latex
}
