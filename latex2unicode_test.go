package latex2unicode

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`\alpha + \beta = \gamma`, `α + β = γ`},
		{`\sqrt{4} = 2`, `√4 = 2`},
		{`\frac{1}{2} + \frac{3}{4}`, `1⁄2 + 3⁄4`},
		{`\overline{x}`, `x̅`},
		{`\text{hello}`, `hello`},
		{`\left( a \right)`, `( a )`},
	}

	for _, tt := range tests {
		result := ConvertLatexToUnicode(tt.input)
		if result != tt.expected {
			t.Errorf("Convert(%q) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}
