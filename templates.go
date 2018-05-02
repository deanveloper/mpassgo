package mpassgo

import (
	"sort"
	"strings"
)

type PasswordType [][]byte

var Maximum = PasswordType{
	[]byte("anoxxxxxxxxxxxxxxxxx"),
	[]byte("axxxxxxxxxxxxxxxxxno"),
}

var Long = PasswordType{
	[]byte("CvcvnoCvcvCvcv"),
	[]byte("CvcvCvcvnoCvcv"),
	[]byte("CvcvCvcvCvcvno"),
	[]byte("CvccnoCvcvCvcv"),
	[]byte("CvccCvcvnoCvcv"),
	[]byte("CvccCvcvCvcvno"),
	[]byte("CvcvnoCvccCvcv"),
	[]byte("CvcvCvccnoCvcv"),
	[]byte("CvcvCvccCvcvno"),
	[]byte("CvcvnoCvcvCvcc"),
	[]byte("CvcvCvcvnoCvcc"),
	[]byte("CvcvCvcvCvccno"),
	[]byte("CvccnoCvccCvcv"),
	[]byte("CvccCvccnoCvcv"),
	[]byte("CvccCvccCvcvno"),
	[]byte("CvcvnoCvccCvcc"),
	[]byte("CvcvCvccnoCvcc"),
	[]byte("CvcvCvccCvccno"),
	[]byte("CvccnoCvcvCvcc"),
	[]byte("CvccCvcvnoCvcc"),
	[]byte("CvccCvcvCvccno"),
}

var Medium = PasswordType{
	[]byte("CvcnoCvc"),
	[]byte("CvcCvcno"),
}

var Short = PasswordType{
	[]byte("Cvcn"),
}

var Basic = PasswordType{
	[]byte("aaanaaan"),
	[]byte("aannaaan"),
	[]byte("aaannaaa"),
}

var Pin = PasswordType{
	[]byte("nnnn"),
}

// All valid password types represented as a string->PasswordType map.
var PasswordTypes = map[string]PasswordType{
	"maximum": Maximum,
	"max":     Maximum,
	"long":    Long,
	"medium":  Medium,
	"med":     Medium,
	"short":   Short,
	"basic":   Basic,
	"pin":     Pin,
}

// Gets a password type based on the string pwType.
func GetPasswordType(pwType string) (PasswordType, bool) {
	typ, ok := PasswordTypes[strings.ToLower(pwType)]
	return typ, ok
}

// Gets the valid password types.
func ValidPasswordTypes() []string {
	validTypes := make([]string, len(PasswordTypes))
	for k := range PasswordTypes {
		validTypes = append(validTypes, k)
	}
	sort.Strings(validTypes)
	return validTypes
}

var runeMap = map[rune]string{
	'V': "AEIOU",
	'C': "BCDFGHJKLMNPQRSTVWXYZ",
	'v': "aeiou",
	'c': "bcdfghjklmnpqrstvwxyz",
	'A': "AEIOUBCDFGHJKLMNPQRSTVWXYZ",
	'a': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz",
	'n': "0123456789",
	'o': "@&%?,=[]_:-+*$#!'^~;()/.",
	'x': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()",
}
