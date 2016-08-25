package config

import "testing"

func TestLoad(t *testing.T) {
	c, err := Load([]byte(`{"url":"http://localhost","debug":true}`))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if c.URL != "http://localhost" {
		t.Errorf("expected %q got %q", "http://localhost", c.URL)
	}
	if c.LibConfig.URL != "http://localhost" {
		t.Errorf("expected %q got %q", "http://localhost", c.LibConfig.URL)
	}
	if !c.Debug {
		t.Errorf("expected true")
	}
	// this one should be zero value, which is false
	if c.LibConfig.Debug {
		t.Errorf("expected false")
	}
}

func TestDumpFunc(t *testing.T) {
	b, err := Dump(new(AppConfig))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expect := `{"url":"","insecure_tls":false,"source":"","debug":false}`
	if string(b) != expect {
		t.Errorf("expected %q got %q", expect, string(b))
	}
}

func TestDumpMethod(t *testing.T) {
	b, err := new(AppConfig).Dump()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expect := `{"url":"","insecure_tls":false,"source":"","debug":false}`
	if string(b) != expect {
		t.Errorf("expected %q got %q", expect, string(b))
	}
}
