# First Coding Session

~ 15 minutes.

Goals to achieve:

1. Get an OPENAI KEY. If you do now wish to create an account, send me a slack
   message and I'll send you one.
   It is recommended to set a hard limit on your API usage. Do so
   [here](https://platform.openai.com/account/billing/limits).

2. Decide on a strategy for how to interact with the tool you are creating

    2.1 Read filename

    2.2 Read entire diff

    2.3 Get diff using a git library

3. Implement strategy - get the diff into your tool

4. Send a request from your tool to the OpenAI API.
    Use a [library](https://platform.openai.com/docs/libraries) or build your
    own.

## 2. Decide on strategy

You can do this however you want, but I would recommend one of:

* 2.1 Pass filename that contains git diff

E.g.
```sh
$ commit-msg --filename=.git/COMMIT_EDITTMSG
```

Probably easiest ⭐

* 2.2 Read entire diff

E.g.
```sh
$ commit-msg --changes=$(git diff)
```

* 2.3 Get diff using a git library

This would use a git library (or call git as a shell command) to get a git diff.

## 3. Implement strategy

### Example Go

Using 2.1, passing in a filename and parsing the file.

```go

var (
	filename           string // this is populated from a CLI flag --file
)

func main() {

	// ...

	gitDiff, err := readFile()
	if err != nil {
		//nolint:gocritic
		log.Fatalf("could not read file %q: %s", filename, err)
	}

	// ...
}

// readFile will read the file contents, ignoring lines starting with #
// (comments) and return a string.
func readFile() (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("open file %q: %w", filename, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	sb := strings.Builder{}

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if strings.HasPrefix(currentLine, "#") {
			continue
		}
		sb.WriteString(currentLine)
	}

	return sb.String(), nil
}
```
