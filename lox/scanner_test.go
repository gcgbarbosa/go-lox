package lox

import "testing"

func TestScanTokens(t *testing.T) {
	scanner := ScanTokens("()")

  source := "()"

	if scanner.source != source {
		t.Errorf("expected source to be %q, got %q", source, scanner.source)
	}
	if scanner.start != 0 {
		t.Errorf("expected start to be 0, got %d", scanner.start)
	}
	if scanner.current != 0 {
		t.Errorf("expected current to be 0, got %d", scanner.current)
	}
	if scanner.line != 1 {
		t.Errorf("expected line to be 1, got %d", scanner.line)
	}
	if len(scanner.tokens) != 3 {
    t.Errorf("expected three tokens: ( ) EOL")  
	}
}
