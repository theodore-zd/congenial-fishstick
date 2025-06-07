package main

import (
	"strings"
)

// TailwindConfig represents the complete Tailwind CSS configuration
type TailwindConfig struct {
	Colors       map[string]string
	Spacing      map[string]string
	Screens      map[string]string
	FontSize     map[string]FontSize
	FontFamily   map[string][]string
	BoxShadow    map[string]string
	BorderRadius map[string]string
	Utilities    map[string]UtilityDefinition
}

// FontSize represents font size with line height
type FontSize struct {
	Size       string
	LineHeight string
}

// UtilityDefinition represents a utility class definition
type UtilityDefinition struct {
	Property   string
	Values     map[string]string
	Responsive bool
	Hover      bool
	Focus      bool
	Active     bool
}

// TailwindUtilities represents all possible Tailwind utility classes
type TailwindUtilities struct {
	// Layout utilities
	Container          map[string]string
	Box                map[string]string
	Display            map[string]string
	Float              map[string]string
	Clear              map[string]string
	Isolation          map[string]string
	ObjectFit          map[string]string
	ObjectPosition     map[string]string
	Overflow           map[string]string
	OverscrollBehavior map[string]string
	Position           map[string]string
	TopRightBottomLeft map[string]string
	Visibility         map[string]string
	ZIndex             map[string]string

	// Flexbox & Grid
	FlexBasis           map[string]string
	FlexDirection       map[string]string
	FlexWrap            map[string]string
	Flex                map[string]string
	FlexGrow            map[string]string
	FlexShrink          map[string]string
	Order               map[string]string
	GridTemplateColumns map[string]string
	GridColumn          map[string]string
	GridTemplateRows    map[string]string
	GridRow             map[string]string
	GridAutoFlow        map[string]string
	GridAutoColumns     map[string]string
	GridAutoRows        map[string]string
	Gap                 map[string]string
	JustifyContent      map[string]string
	JustifyItems        map[string]string
	JustifySelf         map[string]string
	AlignContent        map[string]string
	AlignItems          map[string]string
	AlignSelf           map[string]string
	PlaceContent        map[string]string
	PlaceItems          map[string]string
	PlaceSelf           map[string]string

	// Spacing
	Padding map[string]string
	Margin  map[string]string
	SpaceX  map[string]string
	SpaceY  map[string]string

	// Sizing
	Width     map[string]string
	MinWidth  map[string]string
	MaxWidth  map[string]string
	Height    map[string]string
	MinHeight map[string]string
	MaxHeight map[string]string
	Size      map[string]string

	// Typography
	FontFamily              map[string]string
	FontSize                map[string]string
	FontSmoothing           map[string]string
	FontStyle               map[string]string
	FontWeight              map[string]string
	FontVariantNumeric      map[string]string
	LetterSpacing           map[string]string
	LineClamp               map[string]string
	LineHeight              map[string]string
	ListStyleImage          map[string]string
	ListStylePosition       map[string]string
	ListStyleType           map[string]string
	TextAlign               map[string]string
	TextColor               map[string]string
	TextDecoration          map[string]string
	TextDecorationColor     map[string]string
	TextDecorationStyle     map[string]string
	TextDecorationThickness map[string]string
	TextUnderlineOffset     map[string]string
	TextTransform           map[string]string
	TextOverflow            map[string]string
	TextIndent              map[string]string
	VerticalAlign           map[string]string
	Whitespace              map[string]string
	WordBreak               map[string]string
	Hyphens                 map[string]string
	Content                 map[string]string

	// Backgrounds
	BackgroundAttachment map[string]string
	BackgroundClip       map[string]string
	BackgroundColor      map[string]string
	BackgroundOrigin     map[string]string
	BackgroundPosition   map[string]string
	BackgroundRepeat     map[string]string
	BackgroundSize       map[string]string
	BackgroundImage      map[string]string
	GradientColorStops   map[string]string

	// Borders
	BorderRadius    map[string]string
	BorderWidth     map[string]string
	BorderColor     map[string]string
	BorderStyle     map[string]string
	DivideWidth     map[string]string
	DivideColor     map[string]string
	DivideStyle     map[string]string
	OutlineWidth    map[string]string
	OutlineColor    map[string]string
	OutlineStyle    map[string]string
	OutlineOffset   map[string]string
	RingWidth       map[string]string
	RingColor       map[string]string
	RingOffsetWidth map[string]string
	RingOffsetColor map[string]string

	// Effects
	BoxShadow           map[string]string
	BoxShadowColor      map[string]string
	Opacity             map[string]string
	MixBlendMode        map[string]string
	BackgroundBlendMode map[string]string

	// Filters
	Blur               map[string]string
	Brightness         map[string]string
	Contrast           map[string]string
	DropShadow         map[string]string
	Grayscale          map[string]string
	HueRotate          map[string]string
	Invert             map[string]string
	Saturate           map[string]string
	Sepia              map[string]string
	BackdropBlur       map[string]string
	BackdropBrightness map[string]string
	BackdropContrast   map[string]string
	BackdropGrayscale  map[string]string
	BackdropHueRotate  map[string]string
	BackdropInvert     map[string]string
	BackdropOpacity    map[string]string
	BackdropSaturate   map[string]string
	BackdropSepia      map[string]string

	// Tables
	BorderCollapse map[string]string
	BorderSpacing  map[string]string
	TableLayout    map[string]string
	CaptionSide    map[string]string

	// Transitions & Animation
	TransitionProperty map[string]string
	TransitionDuration map[string]string
	TransitionTiming   map[string]string
	TransitionDelay    map[string]string
	Animation          map[string]string

	// Transforms
	Scale           map[string]string
	Rotate          map[string]string
	Translate       map[string]string
	Skew            map[string]string
	TransformOrigin map[string]string

	// Interactivity
	AccentColor     map[string]string
	Appearance      map[string]string
	Cursor          map[string]string
	CaretColor      map[string]string
	PointerEvents   map[string]string
	Resize          map[string]string
	ScrollBehavior  map[string]string
	ScrollMargin    map[string]string
	ScrollPadding   map[string]string
	ScrollSnapAlign map[string]string
	ScrollSnapStop  map[string]string
	ScrollSnapType  map[string]string
	TouchAction     map[string]string
	UserSelect      map[string]string
	WillChange      map[string]string

	// SVG
	Fill        map[string]string
	Stroke      map[string]string
	StrokeWidth map[string]string

	// Accessibility
	ScreenReaders map[string]string

	// Official Tailwind color palette
	Colors map[string]map[string]string
}

// ParsedClass represents a parsed Tailwind class
type ParsedClass struct {
	Original  string
	Base      string
	Modifier  string
	Value     string
	Variant   string
	Important bool
	Negative  bool
}

// CSSRule represents a CSS rule
type CSSRule struct {
	Selector    string
	Property    string
	Value       string
	Important   bool
	MediaQuery  string
	PseudoClass string
}

// NewTailwindUtilities creates a new instance with all utility definitions
func NewTailwindUtilities() *TailwindUtilities {
	return &TailwindUtilities{
		// Initialize all utility maps
		Container: map[string]string{
			"container": "width: 100%; margin-left: auto; margin-right: auto;",
		},

		Display: map[string]string{
			"block":              "display: block;",
			"inline-block":       "display: inline-block;",
			"inline":             "display: inline;",
			"flex":               "display: flex;",
			"inline-flex":        "display: inline-flex;",
			"table":              "display: table;",
			"inline-table":       "display: inline-table;",
			"table-caption":      "display: table-caption;",
			"table-cell":         "display: table-cell;",
			"table-column":       "display: table-column;",
			"table-column-group": "display: table-column-group;",
			"table-footer-group": "display: table-footer-group;",
			"table-header-group": "display: table-header-group;",
			"table-row-group":    "display: table-row-group;",
			"table-row":          "display: table-row;",
			"flow-root":          "display: flow-root;",
			"grid":               "display: grid;",
			"inline-grid":        "display: inline-grid;",
			"contents":           "display: contents;",
			"list-item":          "display: list-item;",
			"hidden":             "display: none;",
		},

		Position: map[string]string{
			"static":   "position: static;",
			"fixed":    "position: fixed;",
			"absolute": "position: absolute;",
			"relative": "position: relative;",
			"sticky":   "position: sticky;",
		},

		Float: map[string]string{
			"float-left":  "float: left;",
			"float-right": "float: right;",
			"float-none":  "float: none;",
		},

		Clear: map[string]string{
			"clear-left":  "clear: left;",
			"clear-right": "clear: right;",
			"clear-both":  "clear: both;",
			"clear-none":  "clear: none;",
		},

		Overflow: map[string]string{
			"overflow-auto":      "overflow: auto;",
			"overflow-hidden":    "overflow: hidden;",
			"overflow-clip":      "overflow: clip;",
			"overflow-visible":   "overflow: visible;",
			"overflow-scroll":    "overflow: scroll;",
			"overflow-x-auto":    "overflow-x: auto;",
			"overflow-y-auto":    "overflow-y: auto;",
			"overflow-x-hidden":  "overflow-x: hidden;",
			"overflow-y-hidden":  "overflow-y: hidden;",
			"overflow-x-clip":    "overflow-x: clip;",
			"overflow-y-clip":    "overflow-y: clip;",
			"overflow-x-visible": "overflow-x: visible;",
			"overflow-y-visible": "overflow-y: visible;",
			"overflow-x-scroll":  "overflow-x: scroll;",
			"overflow-y-scroll":  "overflow-y: scroll;",
		},

		Visibility: map[string]string{
			"visible":   "visibility: visible;",
			"invisible": "visibility: hidden;",
			"collapse":  "visibility: collapse;",
		},

		FlexDirection: map[string]string{
			"flex-row":         "flex-direction: row;",
			"flex-row-reverse": "flex-direction: row-reverse;",
			"flex-col":         "flex-direction: column;",
			"flex-col-reverse": "flex-direction: column-reverse;",
		},

		FlexWrap: map[string]string{
			"flex-wrap":         "flex-wrap: wrap;",
			"flex-wrap-reverse": "flex-wrap: wrap-reverse;",
			"flex-nowrap":       "flex-wrap: nowrap;",
		},

		JustifyContent: map[string]string{
			"justify-normal":  "justify-content: normal;",
			"justify-start":   "justify-content: flex-start;",
			"justify-end":     "justify-content: flex-end;",
			"justify-center":  "justify-content: center;",
			"justify-between": "justify-content: space-between;",
			"justify-around":  "justify-content: space-around;",
			"justify-evenly":  "justify-content: space-evenly;",
			"justify-stretch": "justify-content: stretch;",
		},

		AlignItems: map[string]string{
			"items-start":    "align-items: flex-start;",
			"items-end":      "align-items: flex-end;",
			"items-center":   "align-items: center;",
			"items-baseline": "align-items: baseline;",
			"items-stretch":  "align-items: stretch;",
		},

		AlignSelf: map[string]string{
			"self-auto":     "align-self: auto;",
			"self-start":    "align-self: flex-start;",
			"self-end":      "align-self: flex-end;",
			"self-center":   "align-self: center;",
			"self-stretch":  "align-self: stretch;",
			"self-baseline": "align-self: baseline;",
		},

		TextAlign: map[string]string{
			"text-left":    "text-align: left;",
			"text-center":  "text-align: center;",
			"text-right":   "text-align: right;",
			"text-justify": "text-align: justify;",
			"text-start":   "text-align: start;",
			"text-end":     "text-align: end;",
		},

		FontWeight: map[string]string{
			"font-thin":       "font-weight: 100;",
			"font-extralight": "font-weight: 200;",
			"font-light":      "font-weight: 300;",
			"font-normal":     "font-weight: 400;",
			"font-medium":     "font-weight: 500;",
			"font-semibold":   "font-weight: 600;",
			"font-bold":       "font-weight: 700;",
			"font-extrabold":  "font-weight: 800;",
			"font-black":      "font-weight: 900;",
		},

		Colors: map[string]map[string]string{
			"slate": {
				"50":  "#f8fafc",
				"100": "#f1f5f9",
				"200": "#e2e8f0",
				"300": "#cbd5e1",
				"400": "#94a3b8",
				"500": "#64748b",
				"600": "#475569",
				"700": "#334155",
				"800": "#1e293b",
				"900": "#0f172a",
				"950": "#020617",
			},
			"gray": {
				"50":  "#f9fafb",
				"100": "#f3f4f6",
				"200": "#e5e7eb",
				"300": "#d1d5db",
				"400": "#9ca3af",
				"500": "#6b7280",
				"600": "#4b5563",
				"700": "#374151",
				"800": "#1f2937",
				"900": "#111827",
				"950": "#030712",
			},
			"zinc": {
				"50":  "#fafafa",
				"100": "#f4f4f5",
				"200": "#e4e4e7",
				"300": "#d4d4d8",
				"400": "#a1a1aa",
				"500": "#71717a",
				"600": "#52525b",
				"700": "#3f3f46",
				"800": "#27272a",
				"900": "#18181b",
				"950": "#09090b",
			},
			"red": {
				"50":  "#fef2f2",
				"100": "#fee2e2",
				"200": "#fecaca",
				"300": "#fca5a5",
				"400": "#f87171",
				"500": "#ef4444",
				"600": "#dc2626",
				"700": "#b91c1c",
				"800": "#991b1b",
				"900": "#7f1d1d",
				"950": "#450a0a",
			},
			"blue": {
				"50":  "#eff6ff",
				"100": "#dbeafe",
				"200": "#bfdbfe",
				"300": "#93c5fd",
				"400": "#60a5fa",
				"500": "#3b82f6",
				"600": "#2563eb",
				"700": "#1d4ed8",
				"800": "#1e40af",
				"900": "#1e3a8a",
				"950": "#172554",
			},
			"green": {
				"50":  "#f0fdf4",
				"100": "#dcfce7",
				"200": "#bbf7d0",
				"300": "#86efac",
				"400": "#4ade80",
				"500": "#22c55e",
				"600": "#16a34a",
				"700": "#15803d",
				"800": "#166534",
				"900": "#14532d",
				"950": "#052e16",
			},
			// Add more colors as needed...
		},
	}
}

// GenerateSpacingValues generates spacing values for Tailwind
func GenerateSpacingValues() map[string]string {
	spacing := make(map[string]string)

	// Standard spacing values
	spacingMap := map[string]string{
		"0":   "0px",
		"px":  "1px",
		"0.5": "0.125rem",
		"1":   "0.25rem",
		"1.5": "0.375rem",
		"2":   "0.5rem",
		"2.5": "0.625rem",
		"3":   "0.75rem",
		"3.5": "0.875rem",
		"4":   "1rem",
		"5":   "1.25rem",
		"6":   "1.5rem",
		"7":   "1.75rem",
		"8":   "2rem",
		"9":   "2.25rem",
		"10":  "2.5rem",
		"11":  "2.75rem",
		"12":  "3rem",
		"14":  "3.5rem",
		"16":  "4rem",
		"20":  "5rem",
		"24":  "6rem",
		"28":  "7rem",
		"32":  "8rem",
		"36":  "9rem",
		"40":  "10rem",
		"44":  "11rem",
		"48":  "12rem",
		"52":  "13rem",
		"56":  "14rem",
		"60":  "15rem",
		"64":  "16rem",
		"72":  "18rem",
		"80":  "20rem",
		"96":  "24rem",
	}

	for k, v := range spacingMap {
		spacing[k] = v
	}

	return spacing
}

// ParseClass parses a Tailwind class string into its components
func ParseClass(className string) *ParsedClass {
	parsed := &ParsedClass{
		Original: className,
	}

	// Handle important modifier
	if strings.HasSuffix(className, "!") {
		parsed.Important = true
		className = strings.TrimSuffix(className, "!")
	}

	// Handle negative values
	if strings.HasPrefix(className, "-") {
		parsed.Negative = true
		className = strings.TrimPrefix(className, "-")
	}

	// Handle variants (hover:, focus:, etc.)
	if strings.Contains(className, ":") {
		parts := strings.SplitN(className, ":", 2)
		parsed.Variant = parts[0]
		className = parts[1]
	}

	// Handle modifiers (/50, /25, etc.)
	if strings.Contains(className, "/") {
		parts := strings.SplitN(className, "/", 2)
		className = parts[0]
		parsed.Modifier = parts[1]
	}

	// Handle arbitrary values [...]
	if strings.Contains(className, "[") && strings.Contains(className, "]") {
		start := strings.Index(className, "[")
		end := strings.Index(className, "]")
		if start != -1 && end != -1 && end > start {
			parsed.Base = className[:start]
			parsed.Value = className[start+1 : end]
			return parsed
		}
	}

	// Extract base and value
	parts := strings.Split(className, "-")
	if len(parts) > 0 {
		parsed.Base = parts[0]
		if len(parts) > 1 {
			parsed.Value = strings.Join(parts[1:], "-")
		}
	}

	return parsed
}
