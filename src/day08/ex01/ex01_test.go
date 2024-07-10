package ex01

import "testing"

func TestDescribePlantNormalValues(t *testing.T) {
	tests := []struct {
		plant    interface{}
		expected string
	}{
		{
			plant: UnknownPlant{
				FlowerType: "Rose",
				LeafType:   "Oval",
				Color:      255,
			},
			expected: "FlowerType:Rose\nLeafType:Oval\nColor(color_scheme=rgb):255\n",
		},
		{
			plant: AnotherUnknownPlant{
				FlowerColor: 100,
				LeafType:    "Lanceolate",
				Height:      15,
			},
			expected: "FlowerColor:100\nLeafType:Lanceolate\nHeight(unit=inches):15\n",
		},
		{
			plant: UnknownPlant{
				FlowerType: "Lily",
				LeafType:   "Linear",
				Color:      128,
			},
			expected: "FlowerType:Lily\nLeafType:Linear\nColor(color_scheme=rgb):128\n",
		},
	}

	for _, test := range tests {
		result, err := describePlant(test.plant)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != test.expected {
			t.Errorf("For plant %v; expected %v, but got %v", test.plant, test.expected, result)
		}
	}
}

func TestDescribePlantError(t *testing.T) {
	test := struct {
		plant    interface{}
		expected string
	}{
		plant:    1,
		expected: "Error: input is not a struct, it's a int\n",
	}

	_, err := describePlant(test.plant)
	if err.Error() != "Error: input is not a struct, it's a int\n" {
		t.Errorf("For plant %v; expected %v, but got %v", test.plant, test.expected, err)
	}
}
