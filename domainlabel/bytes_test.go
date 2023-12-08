package domainlabel_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2396/domainlabel"
)

func TestBytes(t *testing.T) {

	tests := []struct{
		Value []byte
		ExpectedResult []byte
		ExpectedRest []byte
		ExpectedOK bool
	}{
		{
			Value: nil,
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},
		{
			Value: []byte(nil),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value: []byte{},
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},
		{
			Value: []byte(""),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value:          []byte("0"),
			ExpectedResult: []byte("0"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("1"),
			ExpectedResult: []byte("1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("2"),
			ExpectedResult: []byte("2"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("3"),
			ExpectedResult: []byte("3"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("4"),
			ExpectedResult: []byte("4"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("5"),
			ExpectedResult: []byte("5"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("6"),
			ExpectedResult: []byte("6"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("7"),
			ExpectedResult: []byte("7"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("8"),
			ExpectedResult: []byte("8"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("9"),
			ExpectedResult: []byte("9"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value:          []byte("com"),
			ExpectedResult: []byte("com"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("com/"),
			ExpectedResult: []byte("com"),
			ExpectedRest:      []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("com/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("com"),
			ExpectedRest:      []byte("/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("com."),
			ExpectedResult: []byte("com"),
			ExpectedRest:      []byte("."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("com./"),
			ExpectedResult: []byte("com"),
			ExpectedRest:      []byte("./"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("co"),
			ExpectedResult: []byte("co"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("co/"),
			ExpectedResult: []byte("co"),
			ExpectedRest:     []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("co/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("co"),
			ExpectedRest:     []byte("/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("co."),
			ExpectedResult: []byte("co"),
			ExpectedRest:     []byte("."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("co./"),
			ExpectedResult: []byte("co"),
			ExpectedRest:     []byte("./"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("c"),
			ExpectedResult: []byte("c"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("c/"),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c."),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c./"),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("./"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("c-"),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("-"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-/"),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("-/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("-/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-."),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("-."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-./"),
			ExpectedResult: []byte("c"),
			ExpectedRest:    []byte("-./"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("c1"),
			ExpectedResult: []byte("c1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1/"),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1."),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1./"),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("./"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("c-1"),
			ExpectedResult: []byte("c-1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-1/"),
			ExpectedResult: []byte("c-1"),
			ExpectedRest:      []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-1/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("c-1"),
			ExpectedRest:      []byte("/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-1."),
			ExpectedResult: []byte("c-1"),
			ExpectedRest:      []byte("."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c-1./"),
			ExpectedResult: []byte("c-1"),
			ExpectedRest:      []byte("./"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("c1-"),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("-"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1-/"),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("-/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1-/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("-/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1-."),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("-."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("c1-./"),
			ExpectedResult: []byte("c1"),
			ExpectedRest:     []byte("-./"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("1c-"),
			ExpectedResult: []byte("1c"),
			ExpectedRest:     []byte("-"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1c-/"),
			ExpectedResult: []byte("1c"),
			ExpectedRest:     []byte("-/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1c-/once/twice/thrice/fource.php"),
			ExpectedResult: []byte("1c"),
			ExpectedRest:     []byte("-/once/twice/thrice/fource.php"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1c-."),
			ExpectedResult: []byte("1c"),
			ExpectedRest:     []byte("-."),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1c-./"),
			ExpectedResult: []byte("1c"),
			ExpectedRest:     []byte("-./"),
			ExpectedOK: true,
		},
	}

	for b0:=byte('0'); b0 <= '9'; b0++ {

		{
			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0},
				ExpectedResult: []byte{b0},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('0'); b1 <= '9'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('A'); b1 <= 'Z'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('a'); b1 <= 'z'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for b0:=byte('a'); b0 <= 'z'; b0++ {

		{
			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0},
				ExpectedResult: []byte{b0},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('0'); b1 <= '9'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('A'); b1 <= 'Z'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('a'); b1 <= 'z'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for b0:=byte('A'); b0 <= 'Z'; b0++ {

		{
			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0},
				ExpectedResult: []byte{b0},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('0'); b1 <= '9'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('A'); b1 <= 'Z'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		for b1:=byte('a'); b1 <= 'z'; b1++ {

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{b0,b1},
				ExpectedResult: []byte{b0,b1},
				ExpectedRest: nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actualResult, actualRest, actualOK := domainlabel.Bytes(test.Value)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual ok-result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedResult
			actual   := actualResult

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q", test.Value)
				t.Logf("ACTUAL-REST: %q (%#v)", actualRest, actualRest)
				continue
			}
		}

		{
			expected := test.ExpectedRest
			actual   := actualRest

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual rest is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
