package mpassgo

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
