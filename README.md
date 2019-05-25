```
  ____        ____            _   _
 / ___| ___  / ___| ___ _ __ | \ | | __ _ _ __ ___   ___  ___
| |  _ / _ \| |  _ / _ \ '_ \|  \| |/ _` | '_ ` _ \ / _ \/ __|
| |_| | (_) | |_| |  __/ | | | |\  | (_| | | | | | |  __/\__ \
 \____|\___/ \____|\___|_| |_|_| \_|\__,_|_| |_| |_|\___||___/
                                                                         
```

Golang generate password

# Prerequisites

I assume you have a clean Go installation, that includes `$GOPATH` defined:

```bash
# for macos
echo $GOPATH
/Users/<login>/go

# for linux
echo $GOPATH
/home/<login>/go
```

# Build

For Linux and MacOS only:

```bash
# clone repository
$ git clone git@github.com:bzhtux/go-gennames.git
$ cd go-gennames

# build binary
$ go build -o gen_names
$ chmod u+x gen_names
$ sudo cp gen_names ~/bin/

# test gen_names
$ gen_names
joyful-goose
```

# Docker

If you need to generate password remotly you may want to use a docker container (e.g for concourse-ci):

```bash
# clone repository
$ git clone git@github.com:bzhtux/go-gennames.git

# build docker image
$ cd go-gennames
$ docker build . -t gen-names:1.0.0

# test gen_names binary
$ docker run -ti --rm gen-names:1.0.0
```

