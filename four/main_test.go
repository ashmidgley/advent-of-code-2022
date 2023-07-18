package main

import "testing"

func TestGetRangeOverlapCount(t *testing.T) {
	tt := []struct {
		name     string
		fileName string
		expected int
	}{
		{
			name:     "test file",
			fileName: "test.txt",
			expected: 4,
		},
		{
			name:     "input file",
			fileName: "input.txt",
			expected: 888,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			count, err := getRangeOverlapCount(tc.fileName)

			if err != nil {
				t.Error(err)
			}

			if count != tc.expected {
				t.Errorf("expected %d; got %d", tc.expected, count)
			}
		})
	}
}
