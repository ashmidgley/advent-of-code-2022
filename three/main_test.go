package main

import "testing"

func TestGetPrioritySum(t *testing.T) {
	tt := []struct {
		name     string
		fileName string
		expected int
	}{
		{
			name:     "test file",
			fileName: "test.txt",
			expected: 70,
		},
		{
			name:     "input file",
			fileName: "input.txt",
			expected: 2525,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sum, _ := getPrioritySum(tc.fileName)

			if sum != tc.expected {
				t.Errorf("expected %d; got %d", tc.expected, sum)
			}
		})
	}

}
