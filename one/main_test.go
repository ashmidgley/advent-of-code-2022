package main

import "testing"

func TestGetMaxCalories(t *testing.T) {
	tt := []struct {
		name     string
		fileName string
		expected int
	}{
		{
			name:     "test file",
			fileName: "test.txt",
			expected: 45000,
		},
		{
			name:     "input file",
			fileName: "input.txt",
			expected: 209691,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := getMaxCalories(tc.fileName)
			if err != nil {
				t.Error(err)
			}

			if calories != tc.expected {
				t.Errorf("expected calories %d; got %d", tc.expected, calories)
			}
		})
	}
}
