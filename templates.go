package mpassgo

type TemplateSet [][]byte

var Maximum = TemplateSet{
	[]byte("anoxxxxxxxxxxxxxxxxx"),
	[]byte("axxxxxxxxxxxxxxxxxno"),
}

var Long = TemplateSet{
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

var Medium = TemplateSet{
	[]byte("CvcnoCvc"),
	[]byte("CvcCvcno"),
}

var Short = TemplateSet{
	[]byte("Cvcn"),
}

var Basic = TemplateSet{
	[]byte("aaanaaan"),
	[]byte("aannaaan"),
	[]byte("aaannaaa"),
}

var Pin = TemplateSet{
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
	'X': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()",
}
