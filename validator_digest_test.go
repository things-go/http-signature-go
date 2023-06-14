package httpsign

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/things-go/httpsign/digest"
)

func Test_Validator_Digest(t *testing.T) {
	t.Run("digest mismatch", func(t *testing.T) {
		r, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/a", bytes.NewReader([]byte([]byte("hello world!!"))))

		r.Header.Set(DigestHeader, "invalid body")
		err := NewDigestValidator(digest.DigestHashSha256).Validate(r, &Parameter{})
		require.Equal(t, ErrDigestMismatch, err)
	})
	t.Run("validate digest", func(t *testing.T) {
		tests := []struct {
			name    string
			digest  digest.Digest
			body    []byte
			wantErr error
		}{
			{
				name:    "empty body",
				digest:  digest.DigestHashSha256,
				body:    []byte{},
				wantErr: nil,
			},
			{
				name:    "not empty body",
				digest:  digest.DigestHashSha256,
				body:    []byte("hello world!!"),
				wantErr: nil,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/a", bytes.NewReader([]byte(tt.body)))

				bodyDigest, err := tt.digest.Sign(tt.body)
				require.NoError(t, err)
				r.Header.Set(DigestHeader, bodyDigest)
				err = NewDigestValidator(tt.digest).Validate(r, &Parameter{})
				require.Equal(t, tt.wantErr, err)
			})
		}
	})
}
