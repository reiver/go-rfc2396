package ipv4address_test

import (
	"testing"

	"bytes"
	"fmt"
	"math/rand"
	"time"

	"sourcecode.social/reiver/go-rfc2396/ipv4address"
)

func TestBytes(t *testing.T) {

	var randomness *rand.Rand = rand.New(rand.NewSource( time.Now().Unix() ))

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
			Value:        []byte("apple banana cherry"),
			ExpectedResult: nil,
			ExpectedRest: []byte("apple banana cherry"),
			ExpectedOK: false,
		},



		{
			Value:        []byte("once twice thrice fource"),
			ExpectedResult: nil,
			ExpectedRest: []byte("once twice thrice fource"),
			ExpectedOK: false,
		},



		{
			Value:        []byte("🙂"),
			ExpectedResult: nil,
			ExpectedRest: []byte("🙂"),
			ExpectedOK: false,
		},



		{
			Value:        []byte("1"),
			ExpectedResult: nil,
			ExpectedRest: []byte("1"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("19"),
			ExpectedResult: nil,
			ExpectedRest: []byte("19"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192"),
			ExpectedResult: nil,
			ExpectedRest: []byte("192"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192."),
			ExpectedResult: nil,
			ExpectedRest: []byte("192."),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192.1"),
			ExpectedResult: nil,
			ExpectedRest: []byte("192.1"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192.16"),
			ExpectedResult: nil,
			ExpectedRest: []byte("192.16"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192.168"),
			ExpectedResult: nil,
			ExpectedRest: []byte("192.168"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192.168."),
			ExpectedResult: nil,
			ExpectedRest: []byte("192.168."),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192.168.0"),
			ExpectedResult: nil,
			ExpectedRest: []byte("192.168.0"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("192.168.0."),
			ExpectedResult: nil,
			ExpectedRest: []byte("192.168.0."),
			ExpectedOK: false,
		},
		{
			Value:          []byte("192.168.0.1"),
			ExpectedResult: []byte("192.168.0.1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("192.168.0.1/"),
			ExpectedResult: []byte("192.168.0.1"),
			ExpectedRest:              []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("192.168.0.1/path/to/file.txt"),
			ExpectedResult: []byte("192.168.0.1"),
			ExpectedRest:              []byte("/path/to/file.txt"),
			ExpectedOK: true,
		},



		{
			Value:        []byte("1"),
			ExpectedResult: nil,
			ExpectedRest: []byte("1"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("1."),
			ExpectedResult: nil,
			ExpectedRest: []byte("1."),
			ExpectedOK: false,
		},
		{
			Value:        []byte("1.3"),
			ExpectedResult: nil,
			ExpectedRest: []byte("1.3"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("1.3."),
			ExpectedResult: nil,
			ExpectedRest: []byte("1.3."),
			ExpectedOK: false,
		},
		{
			Value:        []byte("1.3.5"),
			ExpectedResult: nil,
			ExpectedRest: []byte("1.3.5"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("1.3.5."),
			ExpectedResult: nil,
			ExpectedRest: []byte("1.3.5."),
			ExpectedOK: false,
		},
		{
			Value:          []byte("1.3.5.7"),
			ExpectedResult: []byte("1.3.5.7"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("1.3.5.7/"),
			ExpectedResult: []byte("1.3.5.7"),
			ExpectedRest:          []byte("/"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1.3.5.7/apple/banana/cherry.php"),
			ExpectedResult: []byte("1.3.5.7"),
			ExpectedRest:          []byte("/apple/banana/cherry.php"),
			ExpectedOK: true,
		},
	}

	for i:=0; i<50; i++ {

		b0 := randomness.Intn(256)
		b1 := randomness.Intn(256)
		b2 := randomness.Intn(256)
		b3 := randomness.Intn(256)

		ipv4addr := fmt.Sprintf("%d.%d.%d.%d", b0,b1,b2,b3)

		{
			value  := ipv4addr
			result := ipv4addr
			rest   := ""

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte(value),
				ExpectedResult: []byte(result),
				ExpectedRest:   []byte(rest),
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		{
			value  := ipv4addr + "/"
			result := ipv4addr
			rest   :=            "/"

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte(value),
				ExpectedResult: []byte(result),
				ExpectedRest:   []byte(rest),
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		{
			value  := ipv4addr + "/once/twice/thrice/fource"
			result := ipv4addr
			rest   :=            "/once/twice/thrice/fource"

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte(value),
				ExpectedResult: []byte(result),
				ExpectedRest:   []byte(rest),
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}

		{
			value  := ipv4addr + " "
			result := ipv4addr
			rest   :=            " "

			test := struct{
				Value []byte
				ExpectedResult []byte
				ExpectedRest []byte
				ExpectedOK bool
			}{
				Value:          []byte(value),
				ExpectedResult: []byte(result),
				ExpectedRest:   []byte(rest),
				ExpectedOK: true,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actualResult, actualRest, actualOK := ipv4address.Bytes(test.Value)

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
