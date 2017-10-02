package mpassgo

type TemplateSet []string

var Maximum = TemplateSet {
    "anoxxxxxxxxxxxxxxxxx",
    "axxxxxxxxxxxxxxxxxno",
}

var Long = TemplateSet {
    "CvcvnoCvcvCvcv",
    "CvcvCvcvnoCvcv",
    "CvcvCvcvCvcvno",
    "CvccnoCvcvCvcv",
    "CvccCvcvnoCvcv",
    "CvccCvcvCvcvno",
    "CvcvnoCvccCvcv",
    "CvcvCvccnoCvcv",
    "CvcvCvccCvcvno",
    "CvcvnoCvcvCvcc",
    "CvcvCvcvnoCvcc",
    "CvcvCvcvCvccno",
    "CvccnoCvccCvcv",
    "CvccCvccnoCvcv",
    "CvccCvccCvcvno",
    "CvcvnoCvccCvcc",
    "CvcvCvccnoCvcc",
    "CvcvCvccCvccno",
    "CvccnoCvcvCvcc",
    "CvccCvcvnoCvcc",
    "CvccCvcvCvccno",
}

var Medium = TemplateSet {
    "CvcnoCvc",
    "CvcCvcno",
}

var Short = TemplateSet {
    "Cvcn",
}

var Basic = TemplateSet {
    "aaanaaan",
    "aannaaan",
    "aaannaaa",
}

var Pin = TemplateSet {
    "nnnn",
}

var charTemplates = map[rune] string {
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