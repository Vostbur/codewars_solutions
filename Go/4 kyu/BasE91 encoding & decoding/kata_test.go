package kata

import "testing"

type testCase struct {
	d        []byte
	expected string
}

var testEncodeCases = []testCase{
	{[]byte("test"), "fPNKd"},
	{[]byte("Hello World!"), ">OwJh>Io0Tv!8PE"},
}

var testDecodeCases = []testCase{
	{[]byte("fPNKd"), "test"},
	{[]byte(">OwJh>Io0Tv!8PE"), "Hello World!"},
}

func TestEncode(t *testing.T) {
	for _, test := range testEncodeCases {
		if output := EncodeToString(test.d); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range testDecodeCases {
		if output := DecodeToString(test.d); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}

func EncodeToString(d []byte) string {
	return string(Encode(d))
}

func DecodeToString(d []byte) string {
	return string(Decode(d))
}
