package stdlib

// Generated by 'goexports crypto'. Do not edit!

import (
	"crypto"
	"reflect"
)

func init() {
	Value["crypto"] = map[string]reflect.Value{
		"BLAKE2b_256":  reflect.ValueOf(crypto.BLAKE2b_256),
		"BLAKE2b_384":  reflect.ValueOf(crypto.BLAKE2b_384),
		"BLAKE2b_512":  reflect.ValueOf(crypto.BLAKE2b_512),
		"BLAKE2s_256":  reflect.ValueOf(crypto.BLAKE2s_256),
		"MD4":          reflect.ValueOf(crypto.MD4),
		"MD5":          reflect.ValueOf(crypto.MD5),
		"MD5SHA1":      reflect.ValueOf(crypto.MD5SHA1),
		"RIPEMD160":    reflect.ValueOf(crypto.RIPEMD160),
		"RegisterHash": reflect.ValueOf(crypto.RegisterHash),
		"SHA1":         reflect.ValueOf(crypto.SHA1),
		"SHA224":       reflect.ValueOf(crypto.SHA224),
		"SHA256":       reflect.ValueOf(crypto.SHA256),
		"SHA384":       reflect.ValueOf(crypto.SHA384),
		"SHA3_224":     reflect.ValueOf(crypto.SHA3_224),
		"SHA3_256":     reflect.ValueOf(crypto.SHA3_256),
		"SHA3_384":     reflect.ValueOf(crypto.SHA3_384),
		"SHA3_512":     reflect.ValueOf(crypto.SHA3_512),
		"SHA512":       reflect.ValueOf(crypto.SHA512),
		"SHA512_224":   reflect.ValueOf(crypto.SHA512_224),
		"SHA512_256":   reflect.ValueOf(crypto.SHA512_256),
	}
	Type["crypto"] = map[string]reflect.Type{
		"Decrypter":     reflect.TypeOf((*crypto.Decrypter)(nil)).Elem(),
		"DecrypterOpts": reflect.TypeOf((*crypto.DecrypterOpts)(nil)).Elem(),
		"Hash":          reflect.TypeOf((*crypto.Hash)(nil)).Elem(),
		"PrivateKey":    reflect.TypeOf((*crypto.PrivateKey)(nil)).Elem(),
		"PublicKey":     reflect.TypeOf((*crypto.PublicKey)(nil)).Elem(),
		"Signer":        reflect.TypeOf((*crypto.Signer)(nil)).Elem(),
		"SignerOpts":    reflect.TypeOf((*crypto.SignerOpts)(nil)).Elem(),
	}
}