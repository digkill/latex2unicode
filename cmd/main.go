package main

import (
	"fmt"
	"github.com/digkill/latex2unicode/internal/latex2unicode"
)

func main() {
	latexFormula := `\alpha_1 + \beta^2 = \gamma_3, \sqrt{4} = 2, \frac{1}{2} + \frac{3}{4} = x_1^2, \frac{\frac{1}{2}}{3}, \overline{x}, \text{hello}, \left( a \right) \rightarrow \infty`
	unicodeFormula := latex2unicode.ConvertLatexToUnicode(latexFormula)
	fmt.Println("Original (LaTeX):", latexFormula)
	fmt.Println("Result (Unicode):", unicodeFormula)
}
