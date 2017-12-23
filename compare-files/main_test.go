package main

import "testing"

func TestCompareFiles(t *testing.T) {
	testCases := []struct {
		name      string
		filepath1 string
		filepath2 string
		isEqual   bool
	}{
		{
			"f1 != f2",
			"testdata/f1",
			"testdata/f2",
			false,
		},
		{
			"f1 == f2",
			"testdata/f1",
			"testdata/f1-copy",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isEqual, _, _, err := compareFiles(tc.filepath1, tc.filepath2)
			if err != nil {
				t.Fatal(err)
			}
			if isEqual != tc.isEqual {
				t.Errorf("expected isEqual to be %t, got: %t", tc.isEqual, isEqual)
			}
		})
	}
}
