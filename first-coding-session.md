# First Coding Session

Lets get the ball rolling during this first coding session. Get an API KEY,
decide on how you want to interact with your CLI, store the git diff in memory
and try to send any request to OpenAI API.

This should take about ~ 15 minutes.

Goals to achieve:

1. Get an OPENAI KEY. If you do now wish to create an account, send me a slack
   message and I'll send you one.
   It is recommended to set a hard limit on your API usage. Do so
   [here](https://platform.openai.com/account/billing/limits).

2. Decide on a strategy for how to interact with the tool you are creating
	We somehow need to get the git diff into our CLI tool. We can do this in
	multiple ways, some of which are:

    2.1 Read filename

    2.2 Read entire diff

    2.3 Get diff using a git library

3. Implement strategy - get the diff into your tool

4. Send a request from your tool to the OpenAI API.
    Use a [library](https://platform.openai.com/docs/libraries) or build your
    own.

## 1. Get the API KEY

I would recommend you store your API key in a variable.
If you use `zsh` you can store it in `~/.zshenv`:

```sh
──────────────────────────────────────────────────────────────────────────────────────────
       │ File: /Users/philiplinell/.zshenv
──────────────────────────────────────────────────────────────────────────────────────────
   1   │ # -------------------------
   2   │ # This file should not contain any slow instructions as this file is always sourced.
   3   │ # .zshenv should not contain commands that produce output or assume the shell is attached to a tty.
   4   │ # $PATH, $EDITOR, and $PAGER are often set in .zshenv.
   5   │ # -------------------------
   6   │
   7   │
   7   │ export OPENAI_API_KEY=sk-abc123
──────────────────────────────────────────────────────────────────────────────────────────
```

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

### Example node

Using 2.1, passing in a filename and parsing the file.

```javascript
const fs = require('fs');

// Getting the filename from the command line arguments
const filename = process.argv.slice(2)[0];

// Checking if a filename was provided
if (!filename || filename.length === 0) {
    console.log("Please pass a filename as a single argument.");
    process.exit(1);
}

// Checking if the provided file exists
if (!fs.existsSync(filename)) {
    console.log(`File does not exist: ${filename}`);
    process.exit(1);
}

try {
  // Reading the file's content
  const gitDiff = fs.readFileSync(filename, 'utf8');

  console.log(gitDiff);
} catch (err) {
  // Handling any errors that might occur
  console.error(err);
}
```

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

