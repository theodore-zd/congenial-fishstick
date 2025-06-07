package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CSSGenerator handles the generation of CSS from Tailwind classes
type CSSGenerator struct {
	utilities *TailwindUtilities
	spacing   map[string]string
}

// NewCSSGenerator creates a new CSS generator
func NewCSSGenerator() *CSSGenerator {
	return &CSSGenerator{
		utilities: NewTailwindUtilities(),
		spacing:   GenerateSpacingValues(),
	}
}

// GenerateCSS generates CSS for a given class
func (g *CSSGenerator) GenerateCSS(className string) ([]CSSRule, error) {
	parsed := ParseClass(className)
	var rules []CSSRule

	// Handle different utility types
	switch {
	case g.isSpacingUtility(parsed.Base):
		rule, err := g.generateSpacingCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isColorUtility(parsed.Base):
		rule, err := g.generateColorCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isSizingUtility(parsed.Base):
		rule, err := g.generateSizingCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isDisplayUtility(parsed.Base):
		rule, err := g.generateDisplayCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isFlexUtility(parsed.Base):
		rule, err := g.generateFlexCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isTextUtility(parsed.Base):
		rule, err := g.generateTextCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isPositionUtility(parsed.Base):
		rule, err := g.generatePositionCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isBorderUtility(parsed.Base):
		rule, err := g.generateBorderCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isBackgroundUtility(parsed.Base):
		rule, err := g.generateBackgroundCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	case g.isEffectUtility(parsed.Base):
		rule, err := g.generateEffectCSS(parsed)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	default:
		// Try to find in static utilities
		if cssValue, exists := g.findStaticUtility(className); exists {
			rule := CSSRule{
				Selector:    "." + className,
				Property:    "",
				Value:       cssValue,
				Important:   parsed.Important,
				PseudoClass: g.getPseudoClass(parsed.Variant),
				MediaQuery:  g.getMediaQuery(parsed.Variant),
			}
			rules = append(rules, rule)
		} else {
			return nil, fmt.Errorf("unknown utility class: %s", className)
		}
	}

	return rules, nil
}

// isSpacingUtility checks if the class is a spacing utility
func (g *CSSGenerator) isSpacingUtility(base string) bool {
	spacingUtilities := []string{"m", "mt", "mr", "mb", "ml", "mx", "my",
		"p", "pt", "pr", "pb", "pl", "px", "py", "space"}
	for _, util := range spacingUtilities {
		if base == util {
			return true
		}
	}
	return false
}

// generateSpacingCSS generates CSS for spacing utilities
func (g *CSSGenerator) generateSpacingCSS(parsed *ParsedClass) (CSSRule, error) {
	var property string
	var value string

	// Determine the CSS property
	switch parsed.Base {
	case "m":
		property = "margin"
	case "mt":
		property = "margin-top"
	case "mr":
		property = "margin-right"
	case "mb":
		property = "margin-bottom"
	case "ml":
		property = "margin-left"
	case "mx":
		property = "margin-left; margin-right"
	case "my":
		property = "margin-top; margin-bottom"
	case "p":
		property = "padding"
	case "pt":
		property = "padding-top"
	case "pr":
		property = "padding-right"
	case "pb":
		property = "padding-bottom"
	case "pl":
		property = "padding-left"
	case "px":
		property = "padding-left; padding-right"
	case "py":
		property = "padding-top; padding-bottom"
	default:
		return CSSRule{}, fmt.Errorf("unknown spacing utility: %s", parsed.Base)
	}

	// Get the value
	if spacingValue, exists := g.spacing[parsed.Value]; exists {
		value = spacingValue
	} else if parsed.Value != "" {
		// Handle arbitrary values
		value = g.parseArbitraryValue(parsed.Value)
	} else {
		return CSSRule{}, fmt.Errorf("no value provided for spacing utility: %s", parsed.Base)
	}

	// Handle negative values
	if parsed.Negative && value != "0px" {
		value = "-" + value
	}

	cssValue := ""
	if strings.Contains(property, ";") {
		// Handle multiple properties (mx, my, px, py)
		props := strings.Split(property, "; ")
		var declarations []string
		for _, prop := range props {
			declarations = append(declarations, prop+": "+value)
		}
		cssValue = strings.Join(declarations, "; ") + ";"
	} else {
		cssValue = property + ": " + value + ";"
	}

	return CSSRule{
		Selector:    "." + parsed.Original,
		Property:    property,
		Value:       cssValue,
		Important:   parsed.Important,
		PseudoClass: g.getPseudoClass(parsed.Variant),
		MediaQuery:  g.getMediaQuery(parsed.Variant),
	}, nil
}

// isColorUtility checks if the class is a color utility
func (g *CSSGenerator) isColorUtility(base string) bool {
	colorUtilities := []string{"text", "bg", "border", "ring", "divide", "placeholder", "accent", "caret", "fill", "stroke"}
	for _, util := range colorUtilities {
		if base == util {
			return true
		}
	}
	return false
}

// generateColorCSS generates CSS for color utilities
func (g *CSSGenerator) generateColorCSS(parsed *ParsedClass) (CSSRule, error) {
	var property string

	switch parsed.Base {
	case "text":
		property = "color"
	case "bg":
		property = "background-color"
	case "border":
		property = "border-color"
	case "ring":
		property = "--tw-ring-color"
	case "divide":
		property = "--tw-divide-color"
	case "placeholder":
		property = "color"
	case "accent":
		property = "accent-color"
	case "caret":
		property = "caret-color"
	case "fill":
		property = "fill"
	case "stroke":
		property = "stroke"
	default:
		return CSSRule{}, fmt.Errorf("unknown color utility: %s", parsed.Base)
	}

	// Parse color value
	colorValue, err := g.parseColorValue(parsed.Value)
	if err != nil {
		return CSSRule{}, err
	}

	// Handle opacity modifier
	if parsed.Modifier != "" {
		opacity, err := strconv.ParseFloat(parsed.Modifier, 64)
		if err == nil {
			opacity = opacity / 100
			colorValue = g.addOpacityToColor(colorValue, opacity)
		}
	}

	cssValue := property + ": " + colorValue + ";"

	return CSSRule{
		Selector:    "." + parsed.Original,
		Property:    property,
		Value:       cssValue,
		Important:   parsed.Important,
		PseudoClass: g.getPseudoClass(parsed.Variant),
		MediaQuery:  g.getMediaQuery(parsed.Variant),
	}, nil
}

// parseColorValue parses a color value from the Tailwind color system
func (g *CSSGenerator) parseColorValue(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("no color value provided")
	}

	// Handle special colors
	switch value {
	case "inherit":
		return "inherit", nil
	case "current":
		return "currentColor", nil
	case "transparent":
		return "transparent", nil
	case "black":
		return "#000000", nil
	case "white":
		return "#ffffff", nil
	}

	// Parse color-shade format (e.g., "red-500")
	parts := strings.Split(value, "-")
	if len(parts) == 2 {
		colorName := parts[0]
		shade := parts[1]

		if colorMap, exists := g.utilities.Colors[colorName]; exists {
			if colorValue, exists := colorMap[shade]; exists {
				return colorValue, nil
			}
		}
	}

	// Handle arbitrary values
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		return strings.Trim(value, "[]"), nil
	}

	return "", fmt.Errorf("unknown color value: %s", value)
}

// addOpacityToColor adds opacity to a color value
func (g *CSSGenerator) addOpacityToColor(color string, opacity float64) string {
	// Simple implementation - convert hex to rgba
	if strings.HasPrefix(color, "#") && len(color) == 7 {
		r, _ := strconv.ParseInt(color[1:3], 16, 64)
		g, _ := strconv.ParseInt(color[3:5], 16, 64)
		b, _ := strconv.ParseInt(color[5:7], 16, 64)
		return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", r, g, b, opacity)
	}
	return color
}

// isSizingUtility checks if the class is a sizing utility
func (g *CSSGenerator) isSizingUtility(base string) bool {
	sizingUtilities := []string{"w", "h", "min-w", "min-h", "max-w", "max-h", "size"}
	for _, util := range sizingUtilities {
		if base == util {
			return true
		}
	}
	return false
}

// generateSizingCSS generates CSS for sizing utilities
func (g *CSSGenerator) generateSizingCSS(parsed *ParsedClass) (CSSRule, error) {
	var property string

	switch parsed.Base {
	case "w":
		property = "width"
	case "h":
		property = "height"
	case "min-w":
		property = "min-width"
	case "min-h":
		property = "min-height"
	case "max-w":
		property = "max-width"
	case "max-h":
		property = "max-height"
	case "size":
		property = "width; height"
	default:
		return CSSRule{}, fmt.Errorf("unknown sizing utility: %s", parsed.Base)
	}

	value := g.parseSizingValue(parsed.Value)

	cssValue := ""
	if strings.Contains(property, ";") {
		// Handle multiple properties (size)
		props := strings.Split(property, "; ")
		var declarations []string
		for _, prop := range props {
			declarations = append(declarations, prop+": "+value)
		}
		cssValue = strings.Join(declarations, "; ") + ";"
	} else {
		cssValue = property + ": " + value + ";"
	}

	return CSSRule{
		Selector:    "." + parsed.Original,
		Property:    property,
		Value:       cssValue,
		Important:   parsed.Important,
		PseudoClass: g.getPseudoClass(parsed.Variant),
		MediaQuery:  g.getMediaQuery(parsed.Variant),
	}, nil
}

// parseSizingValue parses sizing values
func (g *CSSGenerator) parseSizingValue(value string) string {
	sizingMap := map[string]string{
		"auto":   "auto",
		"full":   "100%",
		"screen": "100vh",
		"svw":    "100svw",
		"lvw":    "100lvw",
		"dvw":    "100dvw",
		"min":    "min-content",
		"max":    "max-content",
		"fit":    "fit-content",
	}

	if val, exists := sizingMap[value]; exists {
		return val
	}

	// Check if it's a spacing value
	if val, exists := g.spacing[value]; exists {
		return val
	}

	// Handle fractions
	if strings.Contains(value, "/") {
		parts := strings.Split(value, "/")
		if len(parts) == 2 {
			numerator, err1 := strconv.ParseFloat(parts[0], 64)
			denominator, err2 := strconv.ParseFloat(parts[1], 64)
			if err1 == nil && err2 == nil && denominator != 0 {
				percentage := (numerator / denominator) * 100
				return fmt.Sprintf("%.6f%%", percentage)
			}
		}
	}

	// Handle arbitrary values
	return g.parseArbitraryValue(value)
}

// parseArbitraryValue parses arbitrary values in square brackets
func (g *CSSGenerator) parseArbitraryValue(value string) string {
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		return strings.Trim(value, "[]")
	}
	return value
}

// isDisplayUtility checks if the class is a display utility
func (g *CSSGenerator) isDisplayUtility(base string) bool {
	_, exists := g.utilities.Display[base]
	return exists
}

// generateDisplayCSS generates CSS for display utilities
func (g *CSSGenerator) generateDisplayCSS(parsed *ParsedClass) (CSSRule, error) {
	if cssValue, exists := g.utilities.Display[parsed.Base]; exists {
		return CSSRule{
			Selector:    "." + parsed.Original,
			Property:    "display",
			Value:       cssValue,
			Important:   parsed.Important,
			PseudoClass: g.getPseudoClass(parsed.Variant),
			MediaQuery:  g.getMediaQuery(parsed.Variant),
		}, nil
	}
	return CSSRule{}, fmt.Errorf("unknown display utility: %s", parsed.Base)
}

// isFlexUtility checks if the class is a flex utility
func (g *CSSGenerator) isFlexUtility(base string) bool {
	flexUtilities := []string{"flex", "justify", "items", "self", "order", "grow", "shrink", "basis"}
	for _, util := range flexUtilities {
		if strings.HasPrefix(base, util) {
			return true
		}
	}
	return false
}

// generateFlexCSS generates CSS for flex utilities
func (g *CSSGenerator) generateFlexCSS(parsed *ParsedClass) (CSSRule, error) {
	fullClass := parsed.Base
	if parsed.Value != "" {
		fullClass = parsed.Base + "-" + parsed.Value
	}

	// Check flex direction
	if cssValue, exists := g.utilities.FlexDirection[fullClass]; exists {
		return CSSRule{
			Selector:    "." + parsed.Original,
			Property:    "flex-direction",
			Value:       cssValue,
			Important:   parsed.Important,
			PseudoClass: g.getPseudoClass(parsed.Variant),
			MediaQuery:  g.getMediaQuery(parsed.Variant),
		}, nil
	}

	// Check justify content
	if cssValue, exists := g.utilities.JustifyContent[fullClass]; exists {
		return CSSRule{
			Selector:    "." + parsed.Original,
			Property:    "justify-content",
			Value:       cssValue,
			Important:   parsed.Important,
			PseudoClass: g.getPseudoClass(parsed.Variant),
			MediaQuery:  g.getMediaQuery(parsed.Variant),
		}, nil
	}

	// Check align items
	if cssValue, exists := g.utilities.AlignItems[fullClass]; exists {
		return CSSRule{
			Selector:    "." + parsed.Original,
			Property:    "align-items",
			Value:       cssValue,
			Important:   parsed.Important,
			PseudoClass: g.getPseudoClass(parsed.Variant),
			MediaQuery:  g.getMediaQuery(parsed.Variant),
		}, nil
	}

	return CSSRule{}, fmt.Errorf("unknown flex utility: %s", fullClass)
}

// isTextUtility checks if the class is a text utility
func (g *CSSGenerator) isTextUtility(base string) bool {
	textUtilities := []string{"text", "font", "leading", "tracking", "uppercase", "lowercase", "capitalize"}
	for _, util := range textUtilities {
		if strings.HasPrefix(base, util) {
			return true
		}
	}
	return false
}

// generateTextCSS generates CSS for text utilities
func (g *CSSGenerator) generateTextCSS(parsed *ParsedClass) (CSSRule, error) {
	fullClass := parsed.Base
	if parsed.Value != "" {
		fullClass = parsed.Base + "-" + parsed.Value
	}

	// Check text align
	if cssValue, exists := g.utilities.TextAlign[fullClass]; exists {
		return CSSRule{
			Selector:    "." + parsed.Original,
			Property:    "text-align",
			Value:       cssValue,
			Important:   parsed.Important,
			PseudoClass: g.getPseudoClass(parsed.Variant),
			MediaQuery:  g.getMediaQuery(parsed.Variant),
		}, nil
	}

	// Check font weight
	if cssValue, exists := g.utilities.FontWeight[fullClass]; exists {
		return CSSRule{
			Selector:    "." + parsed.Original,
			Property:    "font-weight",
			Value:       cssValue,
			Important:   parsed.Important,
			PseudoClass: g.getPseudoClass(parsed.Variant),
			MediaQuery:  g.getMediaQuery(parsed.Variant),
		}, nil
	}

	return CSSRule{}, fmt.Errorf("unknown text utility: %s", fullClass)
}

// isPositionUtility checks if the class is a position utility
func (g *CSSGenerator) isPositionUtility(base string) bool {
	_, exists := g.utilities.Position[base]
	return exists
}

// generatePositionCSS generates CSS for position utilities
func (g *CSSGenerator) generatePositionCSS(parsed *ParsedClass) (CSSRule, error) {
	if cssValue, exists := g.utilities.Position[parsed.Base]; exists {
		return CSSRule{
			Selector:    "." + parsed.Original,
			Property:    "position",
			Value:       cssValue,
			Important:   parsed.Important,
			PseudoClass: g.getPseudoClass(parsed.Variant),
			MediaQuery:  g.getMediaQuery(parsed.Variant),
		}, nil
	}
	return CSSRule{}, fmt.Errorf("unknown position utility: %s", parsed.Base)
}

// isBorderUtility checks if the class is a border utility
func (g *CSSGenerator) isBorderUtility(base string) bool {
	borderUtilities := []string{"border", "rounded", "ring", "outline"}
	for _, util := range borderUtilities {
		if strings.HasPrefix(base, util) {
			return true
		}
	}
	return false
}

// generateBorderCSS generates CSS for border utilities
func (g *CSSGenerator) generateBorderCSS(parsed *ParsedClass) (CSSRule, error) {
	// Simplified border handling
	return CSSRule{}, fmt.Errorf("border utilities not fully implemented yet")
}

// isBackgroundUtility checks if the class is a background utility
func (g *CSSGenerator) isBackgroundUtility(base string) bool {
	return strings.HasPrefix(base, "bg")
}

// generateBackgroundCSS generates CSS for background utilities
func (g *CSSGenerator) generateBackgroundCSS(parsed *ParsedClass) (CSSRule, error) {
	// This is handled by generateColorCSS for bg- classes
	return g.generateColorCSS(parsed)
}

// isEffectUtility checks if the class is an effect utility
func (g *CSSGenerator) isEffectUtility(base string) bool {
	effectUtilities := []string{"shadow", "opacity", "blur", "brightness", "contrast"}
	for _, util := range effectUtilities {
		if strings.HasPrefix(base, util) {
			return true
		}
	}
	return false
}

// generateEffectCSS generates CSS for effect utilities
func (g *CSSGenerator) generateEffectCSS(parsed *ParsedClass) (CSSRule, error) {
	// Simplified effect handling
	return CSSRule{}, fmt.Errorf("effect utilities not fully implemented yet")
}

// findStaticUtility finds a static utility definition
func (g *CSSGenerator) findStaticUtility(className string) (string, bool) {
	// Check all static utility maps
	allUtilities := []map[string]string{
		g.utilities.Display,
		g.utilities.Position,
		g.utilities.Float,
		g.utilities.Clear,
		g.utilities.Overflow,
		g.utilities.Visibility,
		g.utilities.FlexDirection,
		g.utilities.FlexWrap,
		g.utilities.JustifyContent,
		g.utilities.AlignItems,
		g.utilities.AlignSelf,
		g.utilities.TextAlign,
		g.utilities.FontWeight,
	}

	for _, utilityMap := range allUtilities {
		if value, exists := utilityMap[className]; exists {
			return value, true
		}
	}

	return "", false
}

// getPseudoClass returns the CSS pseudo-class for a variant
func (g *CSSGenerator) getPseudoClass(variant string) string {
	pseudoClasses := map[string]string{
		"hover":    ":hover",
		"focus":    ":focus",
		"active":   ":active",
		"visited":  ":visited",
		"disabled": ":disabled",
		"first":    ":first-child",
		"last":     ":last-child",
		"odd":      ":nth-child(odd)",
		"even":     ":nth-child(even)",
	}

	if pseudo, exists := pseudoClasses[variant]; exists {
		return pseudo
	}
	return ""
}

// getMediaQuery returns the media query for a responsive variant
func (g *CSSGenerator) getMediaQuery(variant string) string {
	mediaQueries := map[string]string{
		"sm":  "@media (min-width: 640px)",
		"md":  "@media (min-width: 768px)",
		"lg":  "@media (min-width: 1024px)",
		"xl":  "@media (min-width: 1280px)",
		"2xl": "@media (min-width: 1536px)",
	}

	if media, exists := mediaQueries[variant]; exists {
		return media
	}
	return ""
}

// FormatCSSRule formats a CSS rule as a string
func (g *CSSGenerator) FormatCSSRule(rule CSSRule) string {
	selector := rule.Selector
	if rule.PseudoClass != "" {
		selector += rule.PseudoClass
	}

	cssRule := selector + " { " + rule.Value + " }"

	if rule.Important {
		cssRule = strings.Replace(cssRule, ";", " !important;", -1)
	}

	if rule.MediaQuery != "" {
		cssRule = rule.MediaQuery + " { " + cssRule + " }"
	}

	return cssRule
}
