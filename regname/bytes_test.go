package regname_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2396/regname"
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
			Value:          []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b'}, // == "🙂"
			ExpectedResult: []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b'},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b',   ' '}, // == "🙂"
			ExpectedResult: []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b'},
			ExpectedRest:                                                             []byte{' '},
			ExpectedOK: true,
		},



		{
			Value:          []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f',   ' '},
			ExpectedResult: []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f'},
			ExpectedRest:                                                                                                                                                                              []byte{' '},
			ExpectedOK: true,
		},
		{
			Value:          []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f',   '%'},
			ExpectedResult: []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f'},
			ExpectedRest:                                                                                                                                                                              []byte{'%'},
			ExpectedOK: true,
		},
		{
			Value:          []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f',   '%','0'},
			ExpectedResult: []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f'},
			ExpectedRest:                                                                                                                                                                              []byte{'%','0'},
			ExpectedOK: true,
		},
		{
			Value:          []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f',   '%','0','D'},
			ExpectedResult: []byte{'$',',',';',':','@','&','=','+',   '-','_','.','!','~','*','\'','(',')',   '0','1','2','3','4','5','6','7','8','9',   'A','B','C','D','E','F',  'a','b','c','d','e','f',   '%','0','D'},
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value:          []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b',   ' '}, // == "🙂"
			ExpectedResult: []byte{'%','f','0',  '%','9','f',  '%','9','9',   '%','8','b'},
			ExpectedRest:                                                             []byte{' '},
			ExpectedOK: true,
		},



		{
			Value:        []byte{'%','%','%','%'},
			ExpectedResult: nil,
			ExpectedRest: []byte{'%','%','%','%'},
			ExpectedOK: false,
		},
		{
			Value:        []byte{'%','%','%'},
			ExpectedResult: nil,
			ExpectedRest: []byte{'%','%','%'},
			ExpectedOK: false,
		},
		{
			Value:        []byte{'%','%'},
			ExpectedResult: nil,
			ExpectedRest: []byte{'%','%'},
			ExpectedOK: false,
		},
		{
			Value:        []byte{'%'},
			ExpectedResult: nil,
			ExpectedRest: []byte{'%'},
			ExpectedOK: false,
		},
	}

	for _, b := range []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'} {
		test := struct{
			Value []byte
			ExpectedResult []byte
			ExpectedRest []byte
			ExpectedOK bool
		}{
			Value:        []byte{'%', b},
			ExpectedResult:  nil,
			ExpectedRest: []byte{'%', b},
			ExpectedOK: false,
		}

		tests = append(tests, test)
	}

	for _, b1 := range []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'} {
		for _, b2 := range []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'} {
			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{'%', b1, b2},
				ExpectedResult: []byte{'%', b1, b2},
				ExpectedRest:  nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for _, b1 := range []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'} {
		for _, b2 := range []byte{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'} {
			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{'%', b1, b2},
				ExpectedResult: []byte{'%', b1, b2},
				ExpectedRest:  nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for _, b1 := range []byte{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'} {
		for _, b2 := range []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'} {
			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{'%', b1, b2},
				ExpectedResult: []byte{'%', b1, b2},
				ExpectedRest:  nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for _, b1 := range []byte{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'} {
		for _, b2 := range []byte{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'} {
			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte{'%', b1, b2},
				ExpectedResult: []byte{'%', b1, b2},
				ExpectedRest:  nil,
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actualResult, actualRest, actualOK := regname.Bytes(test.Value)

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
