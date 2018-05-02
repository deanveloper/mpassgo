# mpassgo
Golang port of [MasterPassword], a password manager that generates passwords on the spot,
rather than storing them.

## Installation
Because I'm lazy I do not have any binaries that are precompiled.
Luckily, this package is purely go, and therefore is go-gettable!

Library:
```
go get github.com/deanveloper/mpassgo
```

CLI App:
```
go get github.com/deanveloper/mpassgo/cli
mv $GOPATH/bin/cli $GOPATH/bin/cli mpw
```

## Algorithm
Since it is a port of [MasterPassword], it uses the exact same algorithm, which
can be found [here](http://masterpasswordapp.com/masterpassword-algorithm.pdf) and
[here](https://en.wikipedia.org/wiki/Master_Password).

## Version
I'm honestly not sure what version of masterpassword that this complies to...
I believe that it complies with masterpassword v3 although I could be wrong.

[MasterPassword]: https://masterpasswordapp.com