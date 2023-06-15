package httpsign

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParserSignatureValue(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		params map[string]string
		err    error
	}{
		{
			name:  `Missing = character`,
			input: `keyId="rsa-key-1",algorithm"rsa-sha256",headers="(request-target) host date digest",signature="Hello world"`,
			err:   ErrMissingEqualCharacter,
		},
		{
			name:  `Missing = character`,
			input: `keyId="rsa-key-1",algorithm"rsa-sha256",headers="(request-target) host date digest",signature="Hello world"`,
			err:   ErrMissingEqualCharacter,
		},
		{
			name:  `Missing " at end value`,
			input: `keyId="rsa-key-1",algorithm="rsa-sha256,headers="(request-target) host date digest",signature="Hello world"`,
			err:   ErrMissingDoubleQuote,
		},
		{
			name:  `Missing " at begin value`,
			input: `keyId="rsa-key-1",algorithm=rsa-sha256",headers="(request-target) host date digest",signature="Hello world"`,
			err:   ErrMissingDoubleQuote,
		},
		{
			name:  `number value`,
			input: `keyId="",algorithm="rsa-sha256",created=1686625874,headers="(request-target) host date digest",signature="Hello world"`,
			params: map[string]string{
				"keyId":     "",
				"algorithm": "rsa-sha256",
				"created":   "1686625874",
				"headers":   "(request-target) host date digest",
				"signature": "Hello world",
			},
			err: nil,
		},
		{
			name:  `empty value`,
			input: `keyId="",algorithm="rsa-sha256",headers="(request-target) host date digest",signature="Hello world"`,
			params: map[string]string{
				"keyId":     "",
				"algorithm": "rsa-sha256",
				"headers":   "(request-target) host date digest",
				"signature": "Hello world",
			},
			err: nil,
		},
		{
			name:  `correct test`,
			input: `keyId="rsa-key-1",algorithm="rsa-sha256",headers="(request-target) host date digest",signature="70AaN3BDO0XC9QbtgksgCy2jJvmOvshq8VmjSthdXC+sgcgrKrl9WME4DbZv4W7UZKElvCemhDLHQ1Nln9GMkQ=="`,
			params: map[string]string{
				"keyId":     "rsa-key-1",
				"algorithm": "rsa-sha256",
				"headers":   "(request-target) host date digest",
				"signature": "70AaN3BDO0XC9QbtgksgCy2jJvmOvshq8VmjSthdXC+sgcgrKrl9WME4DbZv4W7UZKElvCemhDLHQ1Nln9GMkQ==",
			},
			err: nil,
		},
	}
	for _, tc := range tests {
		results, err := parseSignatureValue(tc.input)
		require.Equal(t, tc.err, err, tc.name)
		if err != nil {
			continue
		}
		assert.Equal(t, tc.params, results, tc.name)
	}
}
