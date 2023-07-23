package main

import "testing"

func TestGetStackTops(t *testing.T) {
	tt := []struct {
		name     string
		fileName string
		expected string
	}{
		{
			name:     "test file",
			fileName: "test.txt",
			expected: "MCD",
		},
		{
			name:     "input file",
			fileName: "input.txt",
			expected: "GCFGLDNJZ",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getStackTops(tc.fileName)

			if err != nil {
				t.Error(err)
			}

			if result != tc.expected {
				t.Errorf("expected %s; got %s", tc.expected, result)
			}
		})
	}
}
