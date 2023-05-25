# Readme

This folder contains example commits and the file diffs that can be used for
testing the git commit suggestion tool.

  - [angular.md](./angular.md)
  - [git.md](./git.md)
  - [go.md](./go.md)
  - [react.md](./react.md)

This is how I've created them:

```sh
# checkout any interesting commit
$ git checkout f8de255e94540f9018d8196b3a34da500707c39b

# store files changed in file
$ git diff HEAD~ > diff.txt


# alternatively you can copy them to clipboard directly (use 'xclip -selection clipboard'
# on linux)
$ git diff HEAD~ | pbcopy
```
