package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
var key = []byte("1234567890123456")

func TestCryptCFB(t *testing.T) {
	type args struct {
		v []byte
		k []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				v: []byte("crypt"),
				k: key,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CryptCFB(tt.args.v, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptCFB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			assert.NotEmpty(t, got)
			d, err := DecryptCFB(got, tt.args.k)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, d, tt.args.v)
		})
	}
}

func TestDecryptCFB(t *testing.T) {
	type args struct {
		v []byte
		k []byte
	}
	v := "crypt"
	d, err := CryptCFB([]byte(v), key)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				v: d,
				k: key,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptCFB(tt.args.v, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptCFB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			assert.NotEmpty(t, got)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, v, string(got))
		})
	}
}

func TestCryptCBC(t *testing.T) {
	type args struct {
		v []byte
		k []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				v: []byte("crypt"),
				k: key,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CryptCBC(tt.args.v, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptCFB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			assert.NotEmpty(t, got)
			d, err := DecryptCBC(got, tt.args.k)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, d, tt.args.v)
		})
	}
}

func TestDecryptCBC(t *testing.T) {
	type args struct {
		v []byte
		k []byte
	}
	v := "crypt"
	d, err := CryptCBC([]byte(v), key)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				v: d,
				k: key,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptCBC(tt.args.v, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptCFB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			assert.NotEmpty(t, got)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, v, string(got))
		})
	}
}
