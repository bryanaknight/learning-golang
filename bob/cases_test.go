package bob

// Source: exercism/problem-specifications
// Commit: 6dc2014 bob: apply "input" policy
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		"stating something",
		"Tom-ay-to, tom-aaaah-to.",
		"Whatever.",
	},
	{
		"shouting",
		"WATCH OUT!",
		"Whoa, chill out!",
	},
	{
		"shouting gibberish",
		"FCECDFCAAB",
		"Whoa, chill out!",
	},
	{
		"asking a question",
		"Does this cryogenic chamber make me look fat?",
		"Sure.",
	},
	{
		"asking a numeric question",
		"You are, what, like 15?",
		"Sure.",
	},
	{
		"asking gibberish",
		"fffbbcbeab?",
		"Sure.",
	},
	{
		"talking forcefully",
		"Let's go make out behind the gym!",
		"Whatever.",
	},
	{
		"using acronyms in regular speech",
		"It's OK if you don't want to go to the DMV.",
		"Whatever.",
	},
	{
		"forceful question",
		"WHAT THE HELL WERE YOU THINKING?",
		"Calm down, I know what I'm doing!",
	},
	{
		"shouting numbers",
		"1, 2, 3 GO!",
		"Whoa, chill out!",
	},
}
