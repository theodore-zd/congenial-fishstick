package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
	cornflowerblue: [100, 149, 237],
	cornsilk: [255, 248, 220],
	crimson: [220, 20, 60],
	cyan: [0, 255, 255],
	darkblue: [0, 0, 139],
	darkcyan: [0, 139, 139],
	darkgoldenrod: [184, 134, 11],
	darkgray: [169, 169, 169],
	darkgreen: [0, 100, 0],
	darkgrey: [169, 169, 169],
	darkkhaki: [189, 183, 107],
	darkmagenta: [139, 0, 139],
	darkolivegreen: [85, 107, 47],
	darkorange: [255, 140, 0],
	darkorchid: [153, 50, 204],
	darkred: [139, 0, 0],
	darksalmon: [233, 150, 122],
	darkseagreen: [143, 188, 143],
	darkslateblue: [72, 61, 139],
	darkslategray: [47, 79, 79],
	darkslategrey: [47, 79, 79],
	darkturquoise: [0, 206, 209],
	darkviolet: [148, 0, 211],
	deeppink: [255, 20, 147],
	deepskyblue: [0, 191, 255],
	dimgray: [105, 105, 105],
	dimgrey: [105, 105, 105],
	dodgerblue: [30, 144, 255],
	firebrick: [178, 34, 34],
	floralwhite: [255, 250, 240],
	forestgreen: [34, 139, 34],
	fuchsia: [255, 0, 255],
	gainsboro: [220, 220, 220],
	ghostwhite: [248, 248, 255],
	gold: [255, 215, 0],
	goldenrod: [218, 165, 32],
	gray: [128, 128, 128],
	green: [0, 128, 0],
	greenyellow: [173, 255, 47],
	grey: [128, 128, 128],
	honeydew: [240, 255, 240],
	hotpink: [255, 105, 180],
	indianred: [205, 92, 92],
	indigo: [75, 0, 130],
	ivory: [255, 255, 240],
	khaki: [240, 230, 140],
	lavender: [230, 230, 250],
	lavenderblush: [255, 240, 245],
	lawngreen: [124, 252, 0],
	lemonchiffon: [255, 250, 205],
	lightblue: [173, 216, 230],
	lightcoral: [240, 128, 128],
	lightcyan: [224, 255, 255],
	lightgoldenrodyellow: [250, 250, 210],
	lightgray: [211, 211, 211],
	lightgreen: [144, 238, 144],
	lightgrey: [211, 211, 211],
	lightpink: [255, 182, 193],
	lightsalmon: [255, 160, 122],
	lightseagreen: [32, 178, 170],
	lightskyblue: [135, 206, 250],
	lightslategray: [119, 136, 153],
	lightslategrey: [119, 136, 153],
	lightsteelblue: [176, 196, 222],
	lightyellow: [255, 255, 224],
	lime: [0, 255, 0],
	limegreen: [50, 205, 50],
	linen: [250, 240, 230],
	magenta: [255, 0, 255],
	maroon: [128, 0, 0],
	mediumaquamarine: [102, 205, 170],
	mediumblue: [0, 0, 205],
	mediumorchid: [186, 85, 211],
	mediumpurple: [147, 112, 219],
	mediumseagreen: [60, 179, 113],
	mediumslateblue: [123, 104, 238],
	mediumspringgreen: [0, 250, 154],
	mediumturquoise: [72, 209, 204],
	mediumvioletred: [199, 21, 133],
	midnightblue: [25, 25, 112],
	mintcream: [245, 255, 250],
	mistyrose: [255, 228, 225],
	moccasin: [255, 228, 181],
	navajowhite: [255, 222, 173],
	navy: [0, 0, 128],
	oldlace: [253, 245, 230],
	olive: [128, 128, 0],
	olivedrab: [107, 142, 35],
	orange: [255, 165, 0],
	orangered: [255, 69, 0],
	orchid: [218, 112, 214],
	palegoldenrod: [238, 232, 170],
	palegreen: [152, 251, 152],
	paleturquoise: [175, 238, 238],
	palevioletred: [219, 112, 147],
	papayawhip: [255, 239, 213],
	peachpuff: [255, 218, 185],
	peru: [205, 133, 63],
	pink: [255, 192, 203],
	plum: [221, 160, 221],
	powderblue: [176, 224, 230],
	purple: [128, 0, 128],
	rebeccapurple: [102, 51, 153],
	red: [255, 0, 0],
	rosybrown: [188, 143, 143],
	royalblue: [65, 105, 225],
	saddlebrown: [139, 69, 19],
	salmon: [250, 128, 114],
	sandybrown: [244, 164, 96],
	seagreen: [46, 139, 87],
	seashell: [255, 245, 238],
	sienna: [160, 82, 45],
	silver: [192, 192, 192],
	skyblue: [135, 206, 235],
	slateblue: [106, 90, 205],
	slategray: [112, 128, 144],
	slategrey: [112, 128, 144],
	snow: [255, 250, 250],
	springgreen: [0, 255, 127],
	steelblue: [70, 130, 180],
	tan: [210, 180, 140],
	teal: [0, 128, 128],
	thistle: [216, 191, 216],
	tomato: [255, 99, 71],
	turquoise: [64, 224, 208],
	violet: [238, 130, 238],
	wheat: [245, 222, 179],
	white: [255, 255, 255],
	whitesmoke: [245, 245, 245],
	yellow: [255, 255, 0],
	yellowgreen: [154, 205, 50]
}

var tailwindColors = TailwindColors{
		inherit: "inherit",
		current: "currentColor",
		transparent: "transparent",
		black: "#000",
		white: "#fff",
		slate: {
			50: "#f8fafc",
			100: "#f1f5f9",
			200: "#e2e8f0",
			300: "#cbd5e1",
			400: "#94a3b8",
			500: "#64748b",
			600: "#475569",
			700: "#334155",
			800: "#1e293b",
			900: "#0f172a",
			950: "#020617"
		},
		gray: {
			50: "#f9fafb",
			100: "#f3f4f6",
			200: "#e5e7eb",
			300: "#d1d5db",
			400: "#9ca3af",
			500: "#6b7280",
			600: "#4b5563",
			700: "#374151",
			800: "#1f2937",
			900: "#111827",
			950: "#030712"
		},
		zinc: {
			50: "#fafafa",
			100: "#f4f4f5",
			200: "#e4e4e7",
			300: "#d4d4d8",
			400: "#a1a1aa",
			500: "#71717a",
			600: "#52525b",
			700: "#3f3f46",
			800: "#27272a",
			900: "#18181b",
			950: "#09090b"
		},
		neutral: {
			50: "#fafafa",
			100: "#f5f5f5",
			200: "#e5e5e5",
			300: "#d4d4d4",
			400: "#a3a3a3",
			500: "#737373",
			600: "#525252",
			700: "#404040",
			800: "#262626",
			900: "#171717",
			950: "#0a0a0a"
		},
		stone: {
			50: "#fafaf9",
			100: "#f5f5f4",
			200: "#e7e5e4",
			300: "#d6d3d1",
			400: "#a8a29e",
			500: "#78716c",
			600: "#57534e",
			700: "#44403c",
			800: "#292524",
			900: "#1c1917",
			950: "#0c0a09"
		},
		red: {
			50: "#fef2f2",
			100: "#fee2e2",
			200: "#fecaca",
			300: "#fca5a5",
			400: "#f87171",
			500: "#ef4444",
			600: "#dc2626",
			700: "#b91c1c",
			800: "#991b1b",
			900: "#7f1d1d",
			950: "#450a0a"
		},
		orange: {
			50: "#fff7ed",
			100: "#ffedd5",
			200: "#fed7aa",
			300: "#fdba74",
			400: "#fb923c",
			500: "#f97316",
			600: "#ea580c",
			700: "#c2410c",
			800: "#9a3412",
			900: "#7c2d12",
			950: "#431407"
		},
		amber: {
			50: "#fffbeb",
			100: "#fef3c7",
			200: "#fde68a",
			300: "#fcd34d",
			400: "#fbbf24",
			500: "#f59e0b",
			600: "#d97706",
			700: "#b45309",
			800: "#92400e",
			900: "#78350f",
			950: "#451a03"
		},
		yellow: {
			50: "#fefce8",
			100: "#fef9c3",
			200: "#fef08a",
			300: "#fde047",
			400: "#facc15",
			500: "#eab308",
			600: "#ca8a04",
			700: "#a16207",
			800: "#854d0e",
			900: "#713f12",
			950: "#422006"
		},
		lime: {
			50: "#f7fee7",
			100: "#ecfccb",
			200: "#d9f99d",
			300: "#bef264",
			400: "#a3e635",
			500: "#84cc16",
			600: "#65a30d",
			700: "#4d7c0f",
			800: "#3f6212",
			900: "#365314",
			950: "#1a2e05"
		},
		green: {
			50: "#f0fdf4",
			100: "#dcfce7",
			200: "#bbf7d0",
			300: "#86efac",
			400: "#4ade80",
			500: "#22c55e",
			600: "#16a34a",
			700: "#15803d",
			800: "#166534",
			900: "#14532d",
			950: "#052e16"
		},
		emerald: {
			50: "#ecfdf5",
			100: "#d1fae5",
			200: "#a7f3d0",
			300: "#6ee7b7",
			400: "#34d399",
			500: "#10b981",
			600: "#059669",
			700: "#047857",
			800: "#065f46",
			900: "#064e3b",
			950: "#022c22"
		},
		teal: {
			50: "#f0fdfa",
			100: "#ccfbf1",
			200: "#99f6e4",
			300: "#5eead4",
			400: "#2dd4bf",
			500: "#14b8a6",
			600: "#0d9488",
			700: "#0f766e",
			800: "#115e59",
			900: "#134e4a",
			950: "#042f2e"
		},
		cyan: {
			50: "#ecfeff",
			100: "#cffafe",
			200: "#a5f3fc",
			300: "#67e8f9",
			400: "#22d3ee",
			500: "#06b6d4",
			600: "#0891b2",
			700: "#0e7490",
			800: "#155e75",
			900: "#164e63",
			950: "#083344"
		},
		sky: {
			50: "#f0f9ff",
			100: "#e0f2fe",
			200: "#bae6fd",
			300: "#7dd3fc",
			400: "#38bdf8",
			500: "#0ea5e9",
			600: "#0284c7",
			700: "#0369a1",
			800: "#075985",
			900: "#0c4a6e",
			950: "#082f49"
		},
		blue: {
			50: "#eff6ff",
			100: "#dbeafe",
			200: "#bfdbfe",
			300: "#93c5fd",
			400: "#60a5fa",
			500: "#3b82f6",
			600: "#2563eb",
			700: "#1d4ed8",
			800: "#1e40af",
			900: "#1e3a8a",
			950: "#172554"
		},
		indigo: {
			50: "#eef2ff",
			100: "#e0e7ff",
			200: "#c7d2fe",
			300: "#a5b4fc",
			400: "#818cf8",
			500: "#6366f1",
			600: "#4f46e5",
			700: "#4338ca",
			800: "#3730a3",
			900: "#312e81",
			950: "#1e1b4b"
		},
		violet: {
			50: "#f5f3ff",
			100: "#ede9fe",
			200: "#ddd6fe",
			300: "#c4b5fd",
			400: "#a78bfa",
			500: "#8b5cf6",
			600: "#7c3aed",
			700: "#6d28d9",
			800: "#5b21b6",
			900: "#4c1d95",
			950: "#2e1065"
		},
		purple: {
			50: "#faf5ff",
			100: "#f3e8ff",
			200: "#e9d5ff",
			300: "#d8b4fe",
			400: "#c084fc",
			500: "#a855f7",
			600: "#9333ea",
			700: "#7e22ce",
			800: "#6b21a8",
			900: "#581c87",
			950: "#3b0764"
		},
		fuchsia: {
			50: "#fdf4ff",
			100: "#fae8ff",
			200: "#f5d0fe",
			300: "#f0abfc",
			400: "#e879f9",
			500: "#d946ef",
			600: "#c026d3",
			700: "#a21caf",
			800: "#86198f",
			900: "#701a75",
			950: "#4a044e"
		},
		pink: {
			50: "#fdf2f8",
			100: "#fce7f3",
			200: "#fbcfe8",
			300: "#f9a8d4",
			400: "#f472b6",
			500: "#ec4899",
			600: "#db2777",
			700: "#be185d",
			800: "#9d174d",
			900: "#831843",
			950: "#500724"
		},
		rose: {
			50: "#fff1f2",
			100: "#ffe4e6",
			200: "#fecdd3",
			300: "#fda4af",
			400: "#fb7185",
			500: "#f43f5e",
			600: "#e11d48",
			700: "#be123c",
			800: "#9f1239",
			900: "#881337",
			950: "#4c0519"
		},
}


// Main function to extract Tailwind utility definitions and generate CSS from HTML
func main() {
	if len(os.Args) != 3 { // Expecting two arguments now
		fmt.Println("Usage: go run tailwind-parser-enhanced.go <javascript-file> <html-file>")
		os.Exit(1)
	}

	jsFilePath := os.Args[1]
	htmlFilePath := os.Args[2]

	definitions, err := extractTailwindDefinitions(jsFilePath)
	if err != nil {
		fmt.Printf("Error extracting definitions: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d unique Tailwind utility definitions.\n", len(definitions))

	// Write all extracted definitions to a file for inspection
	err = writeFullDefinitionsToFile(definitions, "tailwind-definitions.json")
	if err != nil {
		fmt.Printf("Error writing full definitions to file: %v\n", err)
		// Continue execution even if this fails
	} else {
		fmt.Println("All extracted Tailwind definitions written to tailwind-definitions.json")
	}

	// Generate CSS from HTML
	generatedCSS, err := generateCSSFromHTML(htmlFilePath, definitions)
	if err != nil {
		fmt.Printf("Error generating CSS from HTML: %v\n", err)
		os.Exit(1)
	}

	// Write generated CSS to output.css
	outputCSSFile := "output.css"
	err = os.WriteFile(outputCSSFile, []byte(generatedCSS), 0644)
	if err != nil {
		fmt.Printf("Error writing generated CSS to %s: %v\n", outputCSSFile, err)
		os.Exit(1)
	}

	fmt.Printf("\nSuccessfully generated CSS from %s and saved to %s\n", htmlFilePath, outputCSSFile)
}

// writeFullDefinitionsToFile writes the complete extracted definitions to a JSON file
func writeFullDefinitionsToFile(definitions TailwindUtilityDefinitions, filename string) error {
	jsonData, err := json.MarshalIndent(definitions, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal definitions to JSON: %w", err)
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
	}
	return nil
}

// extractClassesFromHTML reads an HTML file and extracts all unique class names from class attributes.
func extractClassesFromHTML(htmlFilePath string) (map[string]bool, error) {
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTML file %s: %w", htmlFilePath, err)
	}

	classSet := make(map[string]bool)
	// Regex to find class attributes and capture their values
	// Handles class="class1 class2" and class='class1 class2'
	classAttrPattern := regexp.MustCompile(`class=(?:\"([^\"]*)\"|\'([^\']*)\')`)
	matches := classAttrPattern.FindAllSubmatch(htmlContent, -1)

	for _, match := range matches {
		var classValue string
		if len(match[1]) > 0 { // Matched double quotes
			classValue = string(match[1])
		} else if len(match[2]) > 0 { // Matched single quotes
			classValue = string(match[2])
		}

		classesInAttr := strings.Fields(classValue) // Split by whitespace
		for _, class := range classesInAttr {
			if strings.TrimSpace(class) != "" {
				classSet[class] = true
			}
		}
	}
	return classSet, nil
}

// generateCSSFromHTML takes an HTML file path and Tailwind definitions,
// extracts classes from HTML, and generates the corresponding CSS.
func generateCSSFromHTML(htmlFilePath string, definitions TailwindUtilityDefinitions) (string, error) {
	htmlClasses, err := extractClassesFromHTML(htmlFilePath)
	if err != nil {
		return "", fmt.Errorf("could not extract classes from HTML: %w", err)
	}

	var cssBuilder strings.Builder
	var foundClasses []string // To sort them for consistent output

	for className := range htmlClasses {
		foundClasses = append(foundClasses, className)
	}
	sort.Strings(foundClasses) // Sort for deterministic output

	for _, className := range foundClasses {
		cssProps, exists := definitions[className]
		if exists {
			if len(cssProps) == 0 { // Class exists but has no props (e.g. special case or parse error)
				// Optionally add a comment or skip
				// cssBuilder.WriteString(fmt.Sprintf("/* .%s - No specific CSS rules found */\n", className))
				continue
			}
			// Skip placeholder types if they don\'t map to direct CSS
			if _, isMatchUtil := cssProps["_type"]; isMatchUtil && cssProps["_type"] == "matchUtility" {
				cssBuilder.WriteString(fmt.Sprintf("/* .%s - (matchUtility, CSS dynamically generated, not included directly) */\n", className))
				continue
			}
			if _, isSpecialCase := cssProps["_type"]; isSpecialCase && cssProps["_type"] == "specialCase" {
				cssBuilder.WriteString(fmt.Sprintf("/* .%s - (specialCase, CSS likely core/preflight, not included directly) */\n", className))
				continue
			}

			cssBuilder.WriteString(fmt.Sprintf(".%s {\n", className))
			// Sort properties for consistent output
			propKeys := make([]string, 0, len(cssProps))
			for k := range cssProps {
				propKeys = append(propKeys, k)
			}
			sort.Strings(propKeys)
			for _, key := range propKeys {
				cssBuilder.WriteString(fmt.Sprintf("  %s: %s;\n", key, cssProps[key]))
			}
			cssBuilder.WriteString("}\n\n")
		} else {
			cssBuilder.WriteString(fmt.Sprintf("/* Class .%s not found in definitions */\n", className))
		}
	}

	return cssBuilder.String(), nil
}

// writeDefinitionsToFile is deprecated, use writeFullDefinitionsToFile for JSON output
// For simple class name list, can be kept or removed.
// For now, let\'s comment it out to avoid confusion.
/*
// writeDefinitionsToFile writes the extracted class names to a file
// TODO: Enhance to write full definitions if needed, for now, just class names or a summary
func writeDefinitionsToFile(definitions TailwindUtilityDefinitions, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	classNames := make([]string, 0, len(definitions))
	for name := range definitions {
		classNames = append(classNames, name)
	}
	sort.Strings(classNames)

	for _, className := range classNames {
		// For now, just write class name. Later, can format and write full CSS.
		// cssProps := definitions[className]
		// propsStr := "" // Format cssProps into a string
		// _, err := file.WriteString(fmt.Sprintf(".%s { %s }\\n", className, propsStr))
		_, err := file.WriteString(className + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}
*/

// extractTailwindDefinitions parses the JavaScript file and returns all Tailwind utility definitions
func extractTailwindDefinitions(filePath string) (TailwindUtilityDefinitions, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	definitions := make(TailwindUtilityDefinitions)

	extractDirectClassDefinitions(content, definitions)
	extractUtilityDefinitions(content, definitions)  // Will need update for full parsing
	extractAddUtilitiesCalls(content, definitions)   // Will need update for full parsing
	extractMatchUtilitiesCalls(content, definitions) // Will need update for full parsing
	extractSpecialCases(content, definitions)        // Will need update for full parsing

	return definitions, nil
}

// parseCSSProperties parses a JSON-like string of CSS properties into a map
func parseCSSProperties(propsStr string) map[string]string {
	props := make(map[string]string)
	// Ensure the string is a valid JSON object by wrapping with {} if not already
	if !strings.HasPrefix(propsStr, "{") {
		propsStr = "{" + propsStr + "}"
	}
	// Replace single quotes with double quotes for valid JSON
	propsStr = strings.ReplaceAll(propsStr, "'", "\"")
	// Attempt to remove trailing commas before a closing brace
	propsStr = regexp.MustCompile(`,(\s*})`).ReplaceAllString(propsStr, "$1")

	err := json.Unmarshal([]byte(propsStr), &props)
	if err != nil {
		// Fallback or error logging if JSON parsing fails
		// This can happen with complex values like CSS variables or functions
		// For now, we'll return an empty map or a map with the raw string
		// fmt.Printf("Warning: Could not parse CSS properties string: %s. Error: %v\\n", propsStr, err)
		// A simple regex fallback for simple key:value pairs
		// This is a basic fallback and might not cover all cases.
		propPattern := regexp.MustCompile(`"([^"]+)":\s*"([^"]+)"`)
		matches := propPattern.FindAllStringSubmatch(propsStr, -1)
		for _, match := range matches {
			if len(match) == 3 {
				props[match[1]] = match[2]
			}
		}
		if len(props) == 0 {
			//  fmt.Printf("Warning: Could not parse CSS properties string using primary or fallback: %s\\n", propsStr)
		}
	}
	return props
}

// extractDirectClassDefinitions finds class definitions like ".bg-fixed":{"background-attachment":"fixed"}
func extractDirectClassDefinitions(content []byte, definitions TailwindUtilityDefinitions) {
	// Regex to find patterns like ".className": { "prop": "value", ... }
	// It captures the class name and the properties block
	pattern := regexp.MustCompile(`"\.([a-zA-Z0-9\\.:/%\[\]_-]+)":\s*({[^}]+})`)
	matches := pattern.FindAllSubmatch(content, -1)

	for _, match := range matches {
		if len(match) == 3 {
			className := string(match[1])
			if !isValidClassName(className) {
				continue
			}
			propsStr := string(match[2])
			cssProperties := parseCSSProperties(propsStr)

			// If parsing resulted in empty properties but the string was not empty,
			// it might indicate a more complex structure not yet handled,
			// or it could be an empty definition.
			if len(cssProperties) > 0 || propsStr == "{}" {
				definitions[className] = cssProperties
			} else if propsStr != "{}" {
				// Store raw if parsing fails and it's not an empty object
				// definitions[className] = map[string]string{"_raw_css": propsStr}
				// For now, if parsing fails for non-empty, we'll add with empty map
				// to avoid polluting with raw strings if not desired.
				// Or, one could log this.
				// fmt.Printf("Note: Storing empty props for class '%s' due to parsing issue with: %s\\n", className, propsStr)
				definitions[className] = make(map[string]string)
			}
		}
	}
}

// extractUtilityDefinitions finds utility definitions in code blocks
// TODO: Update to parse CSS properties similar to extractDirectClassDefinitions
func extractUtilityDefinitions(content []byte, definitions TailwindUtilityDefinitions) {
	blocks := extractCodeBlocks(content)
	for _, block := range blocks {
		if len(block) < 10 || !bytes.Contains(block, []byte(":{")) && !bytes.Contains(block, []byte(`".`)) {
			continue
		}
		// Updated pattern to capture properties block
		pattern := regexp.MustCompile(`"\.([a-zA-Z0-9\\.:/%\[\]_-]+)":\s*({[^}]+})`)
		matches := pattern.FindAllSubmatch(block, -1)
		for _, match := range matches {
			if len(match) == 3 {
				className := string(match[1])
				if !isValidClassName(className) {
					continue
				}
				propsStr := string(match[2])
				cssProperties := parseCSSProperties(propsStr)
				if len(cssProperties) > 0 || propsStr == "{}" {
					definitions[className] = cssProperties
				} else if propsStr != "{}" {
					// fmt.Printf("Note: Storing empty props for class '%s' (from utility block) due to parsing issue with: %s\\n", className, propsStr)
					definitions[className] = make(map[string]string)
				}
			}
		}
	}
}

// extractAddUtilitiesCalls finds classes defined in addUtilities calls
// TODO: Update to parse CSS properties
func extractAddUtilitiesCalls(content []byte, definitions TailwindUtilityDefinitions) {
	pattern := regexp.MustCompile(`addUtilities\s*\(\s*({[^}]+})\s*(?:,\s*[^)]*)?\)`)
	matches := pattern.FindAllSubmatch(content, -1)

	for _, outerMatch := range matches {
		if len(outerMatch) > 1 {
			utilityBlockContent := outerMatch[1]
			// Regex for individual class definitions within the addUtilities object
			classPattern := regexp.MustCompile(`"\.([a-zA-Z0-9\\.:/%\[\]_-]+)":\s*({[^}]+})`)
			classMatches := classPattern.FindAllSubmatch(utilityBlockContent, -1)

			for _, classMatch := range classMatches {
				if len(classMatch) == 3 {
					className := string(classMatch[1])
					if !isValidClassName(className) {
						continue
					}
					propsStr := string(classMatch[2])
					cssProperties := parseCSSProperties(propsStr)

					if len(cssProperties) > 0 || propsStr == "{}" {
						definitions[className] = cssProperties
					} else if propsStr != "{}" {
						// fmt.Printf("Note: Storing empty props for class '%s' (from addUtilities) due to parsing issue with: %s\\n", className, propsStr)
						definitions[className] = make(map[string]string)
					}
				}
			}
		}
	}
}

// extractMatchUtilitiesCalls finds classes defined in matchUtilities calls
// TODO: This is more complex as it often defines dynamic utilities.
// For now, we'll try to extract the base names, but proper CSS generation would require handling the functions.
func extractMatchUtilitiesCalls(content []byte, definitions TailwindUtilityDefinitions) {
	pattern := regexp.MustCompile(`matchUtilities\s*\(\s*{([^}]+)}`)
	matches := pattern.FindAllSubmatch(content, -1)

	for _, match := range matches {
		if len(match) > 1 {
			utilityBlock := match[1]
			classPattern := regexp.MustCompile(`(?:"([a-zA-Z0-9_-]+)":|([a-zA-Z0-9_-]+):)\s*function`)
			classMatches := classPattern.FindAllSubmatch(utilityBlock, -1)

			for _, classMatch := range classMatches {
				var className string
				if len(classMatch[1]) > 0 {
					className = string(classMatch[1])
				} else if len(classMatch) > 2 && len(classMatch[2]) > 0 {
					className = string(classMatch[2])
				}

				if className != "" && isValidClassName(className) {
					// For matchUtilities, the actual CSS is generated by a function.
					// We'll add the base name with a placeholder or note.
					if _, exists := definitions[className]; !exists {
						definitions[className] = map[string]string{"_type": "matchUtility", "_note": "CSS generated by function"}
					}
				}
			}
		}
	}
}

// extractCodeBlocks finds JavaScript code blocks (between { and })
func extractCodeBlocks(content []byte) [][]byte {
	var blocks [][]byte
	var stack []int

	for i := 0; i < len(content); i++ {
		if content[i] == '{' {
			stack = append(stack, i)
		} else if content[i] == '}' && len(stack) > 0 {
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				// This is a top-level block
				blocks = append(blocks, content[start:i+1])
			}
		}
	}

	return blocks
}

// extractSpecialCases handles specific patterns known to define Tailwind classes
// For now, these are treated as class names without specific CSS rules,
// as their rules might be complex or part of Tailwind's core/preflight.
func extractSpecialCases(content []byte, definitions TailwindUtilityDefinitions) {
	// Extract class prefixes from common Tailwind patterns
	prefixPatterns := []string{
		"bg-", "text-", "border-", "p-", "m-", "w-", "h-", "flex-", "grid-",
		"rounded-", "shadow-", "opacity-", "z-", "gap-", "space-", "translate-",
		"scale-", "rotate-", "skew-", "transform-", "transition-", "duration-",
		"ease-", "delay-", "animate-", "cursor-", "outline-", "ring-", "fill-",
		"stroke-", "container", "box-", "float-", "clear-", "object-", "overflow-",
		"overscroll-", "position-", "top-", "right-", "bottom-", "left-", "visible",
		"invisible", "static", "fixed", "absolute", "relative", "sticky", "font-",
		"tracking-", "leading-", "list-", "decoration-", "underline", "line-through",
		"no-underline", "uppercase", "lowercase", "capitalize", "normal-case",
		"truncate", "overflow-ellipsis", "whitespace-", "break-", "placeholder-", "align-",
	}

	// Instead of using Scanner which has line size limits, process the content directly
	contentStr := string(content)

	for _, prefix := range prefixPatterns {
		// Look for patterns like 'prefix-value'
		if strings.Contains(contentStr, prefix) {
			// Try to extract class names that match this pattern
			pattern := regexp.MustCompile(`['".]` + regexp.QuoteMeta(prefix) + `([a-zA-Z0-9_-]+)['"}\s:,\.]`)
			matches := pattern.FindAllStringSubmatch(contentStr, -1)

			for _, match := range matches {
				if len(match) > 1 {
					className := prefix + match[1]
					className = strings.TrimLeft(className, `'".`)
					if isValidClassName(className) {
						if _, exists := definitions[className]; !exists {
							// Mark as special case, actual CSS might be complex/core
							definitions[className] = map[string]string{"_type": "specialCase", "_note": "CSS likely core or complex"}
						}
					}
				}
			}
		}
	}
}

// isValidClassName checks if a class name is valid
func isValidClassName(name string) bool {
	// Skip JavaScript keywords, function names, and invalid class names
	invalidNames := map[string]bool{
		"function": true, "return": true, "const": true, "var": true, "let": true,
		"if": true, "else": true, "for": true, "while": true, "do": true,
		"switch": true, "case": true, "break": true, "continue": true,
		"new": true, "delete": true, "typeof": true, "instanceof": true,
		"prototype": true, "constructor": true, "default": true,
		"null": true, "undefined": true, "NaN": true, "Infinity": true,
		"true": true, "false": true,
		"theme": true, "variants": true, "plugins": true,
		"values": true, "addUtilities": true, "addComponents": true, "addBase": true,
		"addVariant": true, "matchUtilities": true, "matchVariant": true,
		"type": true, "color": true, "lookup": true, "any": true,
		"length": true, "percentage": true, "position": true, "url": true, "number": true,
		"filter": true, "reduce": true, "map": true, "fill": true, "Object": true,
		"String": true, "Number": true, "Array": true, "Boolean": true, "Function": true,
		"Math": true, "Date": true, "RegExp": true, "JSON": true, "Error": true,
	}

	// Check if it's an invalid name
	if invalidNames[name] {
		return false
	}

	// Class names shouldn't contain certain characters
	if strings.Contains(name, "(") || strings.Contains(name, ")") ||
		strings.Contains(name, "{") || strings.Contains(name, "}") ||
		strings.Contains(name, "=>") || strings.Contains(name, "==") ||
		strings.Contains(name, "\"") || strings.Contains(name, "'") ||
		strings.Contains(name, ":") || strings.Contains(name, "?") ||
		name == "." {
		return false
	}

	// Class names should not be too short or too long
	if len(name) < 2 || len(name) > 50 {
		return false
	}

	return true
}
