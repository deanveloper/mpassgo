package mpassgo

var Maximum = []string {
    "anoxxxxxxxxxxxxxxxxx",
    "axxxxxxxxxxxxxxxxxno",
}

var Long = []string {
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

var Medium = []string {
    "CvcnoCvc",
    "CvcCvcno",
}

var Short = []string {
    "Cvcn",
}

var Basic = []string {
    "aaanaaan",
    "aannaaan",
    "aaannaaa",
}

var Pin = []string {
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