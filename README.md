# json

Parse JSON objects from shell scripts.

**Install**: `go get -u github.com/karimsa/json`

## Usage

To read a json file from the shell, just pipe it to the json command and eval the result.

```shell
$ cat sample_file.json
{
  "message": "hello world"
}
$ eval "$(json < sample_file.json)"
$ echo ${message}
hello world
```

To scope all keys to a given key, pass that key as a parameter to json.

```shell
$ cat sample_file.json
{
  "message": "hello world"
}
$ eval "$(json obj < sample_file.json)"
$ echo ${obj_message}
hello world
```

Object paths use `_` instead of `.` - so `path.to.key` becomes `path_to_key`.

Arrays are a bit weird though. My goal was to keep this utility POSIX compliant so it is
as universal as possible. The element index will be a suffix to the variable. To retrieve
the value dynamically, you need to use a bit of `eval` hacking. Here's a pretty function that
can do it for you:

```shell
function at () {
  echo "$(eval echo "\${${1}_${2}${3}}")"
}
```

For a more complete example, see [this](example/test.sh).

## License

Licensed under MIT license.

Copyright (C) 2017-present Karim Alibhai.
