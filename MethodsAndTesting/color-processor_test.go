package asciiart

import (
	"strings"
	"testing"
)

var validColors = []string{
	"red", "green", "yellow", "blue", "magenta", "cyan", "white", "black",
	"bright_red", "bright_green", "bright_yellow", "bright_blue", "bright_magenta", "bright_cyan", "bright_white",
}

func TestApplyColor_AllValidColors(t *testing.T) {
	for _, color := range validColors {
		result := ApplyColor(color, []string{"Hello"})
		if result == "" {
			t.Errorf("%s: expected non-empty output, got empty string", color)
		}
		if !strings.Contains(result, colorStorage[color]) {
			t.Errorf("%s: output missing color code", color)
		}
		if !strings.Contains(result, "\033[0m") {
			t.Errorf("%s: output missing reset code", color)
		}
	}
}

func TestApplyColor_NoSubstring_ColorsEveryLine(t *testing.T) {
	result := ApplyColor("red", []string{"Hello"})
	for _, line := range strings.Split(strings.TrimRight(result, "\n"), "\n") {
		if !strings.HasPrefix(line, colorStorage["red"]) {
			t.Errorf("expected every line to start with color code, got: %q", line)
		}
		if !strings.HasSuffix(line, "\033[0m") {
			t.Errorf("expected every line to end with reset code, got: %q", line)
		}
	}
}

func TestApplyColor_WithSubstring_ContainsColorAndReset(t *testing.T) {
	result := ApplyColor("blue", []string{"He", "Hello"})
	if !strings.Contains(result, colorStorage["blue"]) {
		t.Error("substring color: missing color code in output")
	}
	if !strings.Contains(result, "\033[0m") {
		t.Error("substring color: missing reset code in output")
	}
}

func TestApplyColor_WithSubstring_UncoloredPartsHaveNoColorCode(t *testing.T) {
	result := ApplyColor("green", []string{"He", "Hello"})
	// Strip all colored segments and check remaining text has no color code
	stripped := strings.ReplaceAll(result, colorStorage["green"], "")
	stripped = strings.ReplaceAll(stripped, "\033[0m", "")
	if strings.Contains(stripped, "\033[") {
		t.Error("uncolored portion should not contain any escape codes")
	}
}

func TestApplyColor_EmptySubstring_IsIgnored(t *testing.T) {
	result := ApplyColor("cyan", []string{"", "Hello"})
	if result == "" {
		t.Error("expected non-empty result when empty substring is passed")
	}
}
