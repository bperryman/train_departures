package train_departures

import "testing"

func TestCleanIgnoresCleanText(t *testing.T) {
	if cleanText("hello world") != "hello world" {
		t.Errorf("cleanText broke correct text")
	}
}

func TestCleanTrimsWSBefore(t *testing.T) {
	if cleanText("   hello world") != "hello world" {
		t.Errorf("cleanText not triming spaces before text")
	}
	if cleanText("\t\thello world") != "hello world" {
		t.Errorf("cleanText not triming tabs before text")
	}
	if cleanText("\t \t hello world") != "hello world" {
		t.Errorf("cleanText not triming mixed tabs and spaces before text")
	}
}

func TestCleanTrimsWSAfter(t *testing.T) {
	if cleanText("hello world    ") != "hello world" {
		t.Errorf("cleanText not triming spaces after text")
	}
	if cleanText("hello world\t\t") != "hello world" {
		t.Errorf("cleanText not triming tabs after text")
	}
	if cleanText("hello world\t \t ") != "hello world" {
		t.Errorf("cleanText not triming mixed tabs and spaces after text")
	}
}

func TestCleanTrimExcessSpacesInWords(t *testing.T) {
	if cleanText("hello     world") != "hello world" {
		t.Errorf("cleanText not removing duplicate white space in words")
	}
	if cleanText("hello\t\t\tworld") != "hello world" {
		t.Errorf("cleanText not removing duplicate white space in words")
	}
	if cleanText("hello \t \t \t \t world") != "hello world" {
		t.Errorf("cleanText not removing duplicate white space in words")
	}
}

func TestCleanRemovesUnusualSpecificUnusualChar(t *testing.T) {
	if cleanText("hello\xff\xfdworld") != "hello world" {
		t.Errorf("cleanText not replacing odd character when expected")
	}
	if cleanText("hello \xff\xfd world") != "hello world" {
		t.Errorf("cleanText not replacing odd character when expected")
	}
}
