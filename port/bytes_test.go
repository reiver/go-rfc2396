package port_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2396/port"
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
			ExpectedOK: true,
		},
		{
			Value: []byte(nil),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value: []byte{},
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value: []byte(""),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value:        []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b'}, // == "🙂"
			ExpectedResult: nil,
			ExpectedRest: []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b'},
			ExpectedOK: true,
		},



		{
			Value:        []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f',   ' '},
			ExpectedResult: nil,
			ExpectedRest: []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f',   ' '},
			ExpectedOK: true,
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
			Value:          []byte("8080"),
			ExpectedResult: []byte("8080"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("8080 "),
			ExpectedResult: []byte("8080"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("8080/"),
			ExpectedResult: []byte("8080"),
			ExpectedRest:       []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("8080\n"),
			ExpectedResult: []byte("8080"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
	}

	for _, b1 := range []byte{'0','1','2','3','4','5','6','7','8','9'} {
		for _, b2 := range []byte{'0','1','2','3','4','5','6','7','8','9'} {

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2},
					ExpectedResult: []byte{b1, b2},
					ExpectedRest:  nil,
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, ' '},
					ExpectedResult: []byte{b1, b2},
					ExpectedRest:           []byte{' '},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, '/'},
					ExpectedResult: []byte{b1, b2},
					ExpectedRest:           []byte{'/'},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, '\n'},
					ExpectedResult: []byte{b1, b2},
					ExpectedRest:           []byte{'\n'},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, 'A'},
					ExpectedResult: []byte{b1, b2},
					ExpectedRest:           []byte{'A'},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}
		}
	}

	for _, b1 := range []byte{'0','1','2','3','4','5','6','7','8','9'} {
		for _, b2 := range []byte{'0','1','2','3','4','5','6','7','8','9'} {
			for _, b3 := range []byte{'0','1','2','3','4','5','6','7','8','9'} {

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, b3},
					ExpectedResult: []byte{b1, b2, b3},
					ExpectedRest:  nil,
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, b3, ' '},
					ExpectedResult: []byte{b1, b2, b3},
					ExpectedRest:               []byte{' '},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, b3, '/'},
					ExpectedResult: []byte{b1, b2, b3},
					ExpectedRest:               []byte{'/'},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, b3, '\n'},
					ExpectedResult: []byte{b1, b2, b3},
					ExpectedRest:               []byte{'\n'},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			{
				test := struct{
					Value []byte
					ExpectedResult []byte
					ExpectedRest []byte
					ExpectedOK bool
				}{
					Value:          []byte{b1, b2, b3, 'A'},
					ExpectedResult: []byte{b1, b2, b3},
					ExpectedRest:               []byte{'A'},
					ExpectedOK: true,
				}

				tests = append(tests, test)
			}

			}
		}
	}

	for testNumber, test := range tests {

		actualResult, actualRest, actualOK := port.Bytes(test.Value)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual ok-result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedResult
			actual   := actualResult

			if !bytes.Equal(expected,actual) {
				t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				t.Logf("ACTUAL-REST:   %q (%#v)", actualRest, actualRest)
				continue
			}
		}

		{
			expected := test.ExpectedRest
			actual   := actualRest

			if !bytes.Equal(expected,actual) {
				t.Errorf("For test #%d, the actual rest-result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				continue
			}
		}
	}
}
