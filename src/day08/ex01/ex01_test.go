package ex01

import "testing"

func TestDescribePlant(t *testing.T) {
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
		result := describePlant(test.plant)
		if result != test.expected {
			t.Errorf("For plant %v; expected %v, but got %v", test.plant, test.expected, result)
		}
	}
}
