package logo_test

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/ebiiim/logo"
)

func TestNew(t *testing.T) {
	Log := logo.New(logo.INFO, nil)
	if Log.Level != logo.INFO {
		t.Error("Level")
	}
	if Log.Logger != nil {
		t.Error("Logger")
	}
	if Log.LogMessageMap == nil {
		t.Error("LogMessageMap")
	}
}

func TestLogo_DefaultOutput(t *testing.T) {
	Log := logo.New(logo.INFO, nil)
	Log.I("test")
}

func TestLogo_X(t *testing.T) {
	logMessageMapJP := map[logo.LogLevel]string{logo.DEBUG: "詳細", logo.INFO: "情報", logo.WARNING: "警告", logo.ERROR: "例外"}
	cases := []struct {
		name    string
		lv      logo.LogLevel
		printLv logo.LogLevel
		msm     map[logo.LogLevel]string
		prefix  string
		txt     string
		wantFmt string
	}{
		{"DD", logo.DEBUG, logo.DEBUG, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"DI", logo.DEBUG, logo.INFO, logo.DefaultLogMessageMap, "", "test", ""},
		{"DW", logo.DEBUG, logo.WARNING, logo.DefaultLogMessageMap, "", "test", ""},
		{"DE", logo.DEBUG, logo.ERROR, logo.DefaultLogMessageMap, "", "test", ""},

		{"ID", logo.INFO, logo.DEBUG, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"II", logo.INFO, logo.INFO, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"IW", logo.INFO, logo.WARNING, logo.DefaultLogMessageMap, "", "test", ""},
		{"IE", logo.INFO, logo.ERROR, logo.DefaultLogMessageMap, "", "test", ""},

		{"WD", logo.WARNING, logo.DEBUG, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"WI", logo.WARNING, logo.INFO, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"WW", logo.WARNING, logo.WARNING, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"WE", logo.WARNING, logo.ERROR, logo.DefaultLogMessageMap, "", "test", ""},

		{"ED", logo.ERROR, logo.DEBUG, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"EI", logo.ERROR, logo.INFO, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"EW", logo.ERROR, logo.WARNING, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},
		{"EE", logo.ERROR, logo.ERROR, logo.DefaultLogMessageMap, "", "test", "%s%stest\n"},

		{"DDm", logo.DEBUG, logo.DEBUG, logMessageMapJP, "", "test", "%s%stest\n"},
		{"IIm", logo.INFO, logo.INFO, logMessageMapJP, "", "test", "%s%stest\n"},
		{"WWm", logo.WARNING, logo.WARNING, logMessageMapJP, "", "test", "%s%stest\n"},
		{"EEm", logo.ERROR, logo.ERROR, logMessageMapJP, "", "test", "%s%stest\n"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var b bytes.Buffer
			Log := logo.New(c.printLv, log.New(&b, "", 0))
			Log.LogMessageMap = c.msm
			switch c.lv {
			case logo.DEBUG:
				Log.D(c.txt)
			case logo.INFO:
				Log.I(c.txt)
			case logo.WARNING:
				Log.W(c.txt)
			case logo.ERROR:
				Log.E(c.txt)
			}
			var want string
			if c.wantFmt != "" {
				want = fmt.Sprintf(c.wantFmt, c.prefix, c.msm[c.lv])
			}
			if got := b.String(); got != want {
				t.Errorf("want [%s] but got [%s]", want, got)
			}
		})
	}
}
