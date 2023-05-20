package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	expected := 10.0
	actual := CalculateTax(1000)

	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}

func BatchTestCalculateTax(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{1000, 10.0},
		{500, 5.0},
		{2000, 10.0},
	}
	for _, test := range tests {
		if actual := CalculateTax(test.input); actual != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, actual)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{1000, 500, 2000, 10000, 100000}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount >= 1000 && result != 10.0 {
			t.Errorf("Expected 10.0, got %f", result)
		}
	})
}
