package gocolor

import (
	"testing"
)

var colorMap = map[int]string{
	1:"green",
}
func colorFunc(color, s string) string {
	result := ""
	switch color {
	case "green":
		result = SGreen(s)
	case "white":
		result = SWhite(s)
	case "yellow":
		result = SYellow(s)
	case "red":
		result = SRed(s)
	case "blue":
		result = SBlue(s)
	case "magenta":
		result = SMagenta(s)
	case "cyan":
		result = SCyan(s)
	}
	return result
}

func TestColor(t *testing.T) {
	tests := []struct {
		s string
		color string
		expected string
	}{
		{"g", "green", "\x1b[32mg\n\x1b[0m"},
		{"w", "white", "\x1b[37mw\n\x1b[0m"},
		{"y", "yellow","\x1b[33my\n\x1b[0m"},
		{"r", "red",   "\x1b[31mr\n\x1b[0m"},
		{"b", "blue",  "\x1b[34mb\n\x1b[0m"},
		{"m", "magenta","\x1b[35mm\n\x1b[0m"},
		{"c", "cyan",  "\x1b[36mc\n\x1b[0m"},
	}
	for _, row := range tests {
		actual := colorFunc(row.color, row.s)
		if actual != row.expected {
			t.Errorf("color: %s, actual: %s, expected: %s", row.color, actual, row.expected)
		}
	}
}

func TestBackgroundColor(t *testing.T) {
	GreenBG("绿色背景")
	WhiteBG("白色背景")
	YellowBG("黄色背景")
	RedBG("红色背景")
	BlueBG("蓝色背景")
	MagentaBG("洋红色背景")
	CyanBG("蓝绿色背景")
}

func TestGroupColor(t *testing.T) {
	BlueBG(SRed("红字蓝底色"))
	CyanBG(SYellow("黄字蓝绿底色"))
	MagentaBG(SRed("红字洋底色"))
	GreenBG(SMagentaBG("洋红字绿底"))
}
