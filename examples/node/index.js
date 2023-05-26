// Importing required modules
const fs = require('fs');
const { Configuration, OpenAIApi } = require("openai");

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

// Getting the API key from environment variables
const apiKey = process.env.OPENAI_API_KEY;

// Checking if the API key was provided
if (!apiKey) {
    console.log("Missing OpenAI API Key.");
    process.exit(1);
}

// Configuring the OpenAI API with the API key
const configuration = new Configuration({ apiKey });

// Creating an instance of the OpenAI API
const openai = new OpenAIApi(configuration);

// Function to generate the prompt for the AI model
function createPrompt() {
    return `Given the following git diff, which contains the lines changed
and filenames, please provide an appropriate commit message suggestion. Make
sure to highlight any breaking changes explicitly. The commit message should
consist of a subject and a body, separated by two newlines. The subject,
    written in the imperative mood (e.g., "Add", "Fix", "Change"), should be
brief, 50 characters or less. The body of the message should be wrapped at 72
characters.`;
}

// Main function
async function main() {
    try {
        // Reading the file's content
        const gitDiff = fs.readFileSync(filename, 'utf8');
        // Creating the AI prompt
        const prompt = createPrompt();
        // Combining the prompt and the file's content
        const content = `${prompt}\n\n${gitDiff}`;

        // Sending a request to the OpenAI API and waiting for the result
        const completion = await openai.createChatCompletion({
            model: "gpt-3.5-turbo",
            messages: [{role: "user", content: content}],
        });

        // Logging the AI's response
        console.log(completion.data.choices[0].message.content);

    } catch (err) {
        // Handling any errors that might occur
        console.error(err);
    }
}

// Running the main function
main()
