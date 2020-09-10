package modulus11

import "testing"

func TestMod11(t *testing.T) {
	data := []int{3, 1, 7, 8, 4, 7}
	expected := 1
	cs := Mod11(data)
	if cs != expected {
		t.Errorf("Expected %d got %d", expected, cs)
	}
}

func TestMod11s(t *testing.T) {
	data := "0317847"
	expected := 1
	got, err := Mod11s(data)
	if err != nil {
		t.Errorf("Failed to compute mod11 to %s. err: %v", data, err)
	}
	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}

func TestMod11sWithError(t *testing.T) {
	data := "X317847"
	_, err := Mod11s(data)
	if err == nil {
		t.Errorf("Should not be possible to get a number representation of '%s'", data)
	}
}

func TestMod11i(t *testing.T) {
	data := 317847
	expected := 1
	cs := Mod11i(data)
	if cs != expected {
		t.Errorf("Expected %d got %d", expected, cs)
	}
}
