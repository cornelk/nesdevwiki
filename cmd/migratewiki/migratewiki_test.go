package main

import "testing"

func TestAnchorID(t *testing.T) {
	tests := map[string]string{
		"Banks":                           "Banks",
		"Bank_select_.28.248000.29":       "Bank_select_.28.248000.29",
		"Bank_select_($8000,_even)":       "Bank_select_.28.248000.2C_even.29",
		"cite_note-.244017_lfsr-7":        "cite_note-.244017_lfsr-7",
		"Café":                            "Caf.C3.A9",
		"a_%28b":                          "a_.2528b",
		"Register_overview#Status":        "Register_overview.23Status",
		"2A03_register_map:with-colon.ok": "2A03_register_map:with-colon.ok",
	}
	for in, want := range tests {
		if got := anchorID(in); got != want {
			t.Errorf("anchorID(%q) = %q, want %q", in, got, want)
		}
		if again := anchorID(anchorID(in)); again != anchorID(in) {
			t.Errorf("anchorID is not idempotent for %q: %q", in, again)
		}
	}
}

func TestFragmentIDMatchesHeadingID(t *testing.T) {
	// A heading id and a link to it must agree, whether the export used decoded, escaped or percent-encoded text.
	want := anchorID("Bank_select_($8000,_even)")
	for _, fragment := range []string{"Bank_select_($8000,_even)", "Bank_select_.28.248000.2C_even.29", "Bank_select_%28%248000%2C_even%29"} {
		if got := fragmentID(fragment); got != want {
			t.Errorf("fragmentID(%q) = %q, want %q", fragment, got, want)
		}
	}
}

func TestWikiURL(t *testing.T) {
	tests := map[string]string{
		"MMC3":                       "/wiki/MMC3/",
		"Category:APU":               "/category/APU/",
		"Category:Mappers with IRQs": "/category/Mappers_with_IRQs/",
		"Nesdev wiki:About":          "/wiki/Nesdev_wiki_About/",
		"VT02+ Registers":            "/wiki/VT02+_Registers/",
	}
	for in, want := range tests {
		if got := wikiURL(in); got != want {
			t.Errorf("wikiURL(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestURLScheme(t *testing.T) {
	for href, want := range map[string]bool{
		"https://example.com/a":       true,
		"mailto:a@b.c":                true,
		"Status_flags.xhtml#I:_Flag":  false,
		"Foo.xhtml#a:b":               false,
		"../wiki-images/a.png":        false,
		"NES_2_0_submappers.xhtml#0:": false,
	} {
		if got := urlScheme.MatchString(href); got != want {
			t.Errorf("urlScheme.MatchString(%q) = %v, want %v", href, got, want)
		}
	}
}

func TestUnescapeHeadingIDs(t *testing.T) {
	in := "## Outer (\\$D003) {#Outer_Bank\\_.28.24D003.29}\n\nText with \\_ stays {#not_a_heading\\_}\n"
	want := "## Outer (\\$D003) {#Outer_Bank_.28.24D003.29}\n\nText with \\_ stays {#not_a_heading\\_}\n"
	if got := unescapeHeadingIDs(in); got != want {
		t.Errorf("unescapeHeadingIDs() = %q, want %q", got, want)
	}
}
