package main

import (
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// HTMLParser handles parsing HTML documents to extract Tailwind classes
type HTMLParser struct {
	cssGenerator *CSSGenerator
}

// NewHTMLParser creates a new HTML parser
func NewHTMLParser() *HTMLParser {
	return &HTMLParser{
		cssGenerator: NewCSSGenerator(),
	}
}

// ExtractedClasses represents the result of extracting classes from HTML
type ExtractedClasses struct {
	Classes       []string
	UniqueClasses map[string]bool
	TotalCount    int
}

// ParseHTMLString parses an HTML string and extracts all Tailwind classes
func (p *HTMLParser) ParseHTMLString(htmlContent string) (*ExtractedClasses, error) {
	reader := strings.NewReader(htmlContent)
	return p.ParseHTML(reader)
}

// ParseHTML parses HTML from an io.Reader and extracts all Tailwind classes
func (p *HTMLParser) ParseHTML(reader io.Reader) (*ExtractedClasses, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	extracted := &ExtractedClasses{
		Classes:       []string{},
		UniqueClasses: make(map[string]bool),
		TotalCount:    0,
	}

	p.extractClassesFromNode(doc, extracted)

	return extracted, nil
}

// extractClassesFromNode recursively extracts classes from HTML nodes
func (p *HTMLParser) extractClassesFromNode(node *html.Node, extracted *ExtractedClasses) {
	if node.Type == html.ElementNode {
		// Look for class attribute
		for _, attr := range node.Attr {
			if attr.Key == "class" {
				classes := p.parseClassAttribute(attr.Val)
				for _, class := range classes {
					if p.isTailwindClass(class) {
						extracted.Classes = append(extracted.Classes, class)
						extracted.UniqueClasses[class] = true
						extracted.TotalCount++
					}
				}
			}
		}
	}

	// Recursively process child nodes
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		p.extractClassesFromNode(child, extracted)
	}
}

// parseClassAttribute parses a class attribute value into individual classes
func (p *HTMLParser) parseClassAttribute(classAttr string) []string {
	// Split by whitespace and filter empty strings
	classAttr = strings.TrimSpace(classAttr)
	if classAttr == "" {
		return []string{}
	}

	classes := strings.Fields(classAttr)
	var result []string
	for _, class := range classes {
		class = strings.TrimSpace(class)
		if class != "" {
			result = append(result, class)
		}
	}
	return result
}

// isTailwindClass determines if a class name is a Tailwind utility class
func (p *HTMLParser) isTailwindClass(className string) bool {
	// Basic heuristic to identify Tailwind classes
	// This could be made more sophisticated based on specific requirements

	// Skip obviously non-Tailwind classes
	if strings.Contains(className, "_") ||
		strings.HasPrefix(className, "js-") ||
		strings.HasPrefix(className, "ng-") ||
		strings.HasPrefix(className, "react-") {
		return false
	}

	// Try to parse the class to see if it matches Tailwind patterns
	parsed := ParseClass(className)

	// Check if it's a known Tailwind pattern
	return p.isKnownTailwindPattern(parsed)
}

// isKnownTailwindPattern checks if a parsed class matches known Tailwind patterns
func (p *HTMLParser) isKnownTailwindPattern(parsed *ParsedClass) bool {
	// Check common Tailwind prefixes
	commonPrefixes := []string{
		"m", "mt", "mr", "mb", "ml", "mx", "my", // margin
		"p", "pt", "pr", "pb", "pl", "px", "py", // padding
		"w", "h", "min-w", "min-h", "max-w", "max-h", // sizing
		"text", "bg", "border", "rounded", // color/styling
		"flex", "grid", "block", "inline", "hidden", // display
		"justify", "items", "self", "content", // flexbox
		"absolute", "relative", "fixed", "static", "sticky", // position
		"top", "right", "bottom", "left", // positioning
		"z", "opacity", "shadow", "ring", // effects
		"space", "divide", "gap", // spacing
		"font", "leading", "tracking", // typography
		"cursor", "pointer", "select", // interactivity
		"transition", "duration", "ease", "delay", // transitions
		"transform", "scale", "rotate", "translate", "skew", // transforms
		"overflow", "scroll", "snap", // scrolling
		"sr", "not-sr", // screen readers
	}

	base := parsed.Base

	// Check if the base matches any known prefix
	for _, prefix := range commonPrefixes {
		if base == prefix || strings.HasPrefix(base, prefix+"-") {
			return true
		}
	}

	// Check for responsive/state variants
	if parsed.Variant != "" {
		knownVariants := []string{
			"sm", "md", "lg", "xl", "2xl", // responsive
			"hover", "focus", "active", "visited", "disabled", // states
			"first", "last", "odd", "even", "group-hover", // pseudo
			"dark", "light", // color scheme
		}

		for _, variant := range knownVariants {
			if parsed.Variant == variant || strings.HasPrefix(parsed.Variant, variant+":") {
				return true
			}
		}
	}

	// Check for arbitrary values [...]
	if strings.Contains(parsed.Original, "[") && strings.Contains(parsed.Original, "]") {
		return true
	}

	// Check for fraction values (1/2, 1/3, etc.)
	if strings.Contains(parsed.Value, "/") {
		return true
	}

	return false
}

// GenerateCSSFromHTML parses HTML and generates complete CSS for all found Tailwind classes
func (p *HTMLParser) GenerateCSSFromHTML(htmlContent string) (string, error) {
	extracted, err := p.ParseHTMLString(htmlContent)
	if err != nil {
		return "", err
	}

	return p.GenerateCSSFromClasses(extracted.UniqueClasses)
}

// GenerateCSSFromClasses generates CSS for a map of unique classes
func (p *HTMLParser) GenerateCSSFromClasses(classes map[string]bool) (string, error) {
	var cssRules []string
	var mediaQueryRules map[string][]string = make(map[string][]string)

	for className := range classes {
		rules, err := p.cssGenerator.GenerateCSS(className)
		if err != nil {
			// Log error but continue with other classes
			continue
		}

		for _, rule := range rules {
			formattedRule := p.cssGenerator.FormatCSSRule(rule)

			if rule.MediaQuery != "" {
				// Group media query rules
				if mediaQueryRules[rule.MediaQuery] == nil {
					mediaQueryRules[rule.MediaQuery] = []string{}
				}
				// Extract the inner rule from media query
				innerRule := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(formattedRule, rule.MediaQuery+" { "), " }"))
				mediaQueryRules[rule.MediaQuery] = append(mediaQueryRules[rule.MediaQuery], innerRule)
			} else {
				cssRules = append(cssRules, formattedRule)
			}
		}
	}

	// Build final CSS
	var finalCSS strings.Builder

	// Add base rules
	for _, rule := range cssRules {
		finalCSS.WriteString(rule)
		finalCSS.WriteString("\n")
	}

	// Add media query rules
	for mediaQuery, rules := range mediaQueryRules {
		finalCSS.WriteString("\n")
		finalCSS.WriteString(mediaQuery)
		finalCSS.WriteString(" {\n")
		for _, rule := range rules {
			finalCSS.WriteString("  ")
			finalCSS.WriteString(rule)
			finalCSS.WriteString("\n")
		}
		finalCSS.WriteString("}\n")
	}

	return finalCSS.String(), nil
}

// ExtractClassesWithRegex extracts classes using regex as a fallback method
func (p *HTMLParser) ExtractClassesWithRegex(htmlContent string) []string {
	// Regex to match class attributes
	classRegex := regexp.MustCompile(`class\s*=\s*["']([^"']+)["']`)
	matches := classRegex.FindAllStringSubmatch(htmlContent, -1)

	var allClasses []string
	uniqueClasses := make(map[string]bool)

	for _, match := range matches {
		if len(match) > 1 {
			classes := p.parseClassAttribute(match[1])
			for _, class := range classes {
				if p.isTailwindClass(class) && !uniqueClasses[class] {
					allClasses = append(allClasses, class)
					uniqueClasses[class] = true
				}
			}
		}
	}

	return allClasses
}

// ProcessHTMLFile processes an HTML file and generates corresponding CSS
type ProcessResult struct {
	OriginalHTML     string
	ExtractedClasses *ExtractedClasses
	GeneratedCSS     string
	ClassCount       int
	UniqueClassCount int
}

// ProcessHTMLContent processes HTML content and returns complete result
func (p *HTMLParser) ProcessHTMLContent(htmlContent string) (*ProcessResult, error) {
	extracted, err := p.ParseHTMLString(htmlContent)
	if err != nil {
		return nil, err
	}

	css, err := p.GenerateCSSFromClasses(extracted.UniqueClasses)
	if err != nil {
		return nil, err
	}

	return &ProcessResult{
		OriginalHTML:     htmlContent,
		ExtractedClasses: extracted,
		GeneratedCSS:     css,
		ClassCount:       extracted.TotalCount,
		UniqueClassCount: len(extracted.UniqueClasses),
	}, nil
}

// GetClassStatistics returns statistics about the extracted classes
func (extracted *ExtractedClasses) GetClassStatistics() map[string]int {
	stats := make(map[string]int)

	for _, class := range extracted.Classes {
		parsed := ParseClass(class)
		base := parsed.Base

		// Count by utility type
		if strings.HasPrefix(base, "m") || strings.HasPrefix(base, "p") {
			stats["spacing"]++
		} else if strings.HasPrefix(base, "text") || strings.HasPrefix(base, "bg") {
			stats["colors"]++
		} else if strings.HasPrefix(base, "w") || strings.HasPrefix(base, "h") {
			stats["sizing"]++
		} else if strings.HasPrefix(base, "flex") || strings.HasPrefix(base, "grid") {
			stats["layout"]++
		} else {
			stats["other"]++
		}

		// Count by variant
		if parsed.Variant != "" {
			stats["with_variants"]++
		} else {
			stats["no_variants"]++
		}
	}

	return stats
}
