# Markov Chain Text Generator

A text generator that reads text from `stdin` and generates new text using the **Markov Chain algorithm**.

The program predicts the most probable next word based on the previous word, similar to predictive text on phone keyboards.

## Features

- Reads input text from `stdin`.
- Generates text using the Markov Chain algorithm.
- Uses a suffix length of **1 word**.
- Uses a default prefix length of **2 words**.
- Uses the first two words of the input as the default starting prefix.
- Generates up to **100 words** by default.
- Allows the maximum number of generated words to be configured.
- Allows a custom starting prefix to be specified.
- Allows the prefix length to be configured from **1 to 5 words**.
- Stops generating when the maximum word limit is reached.
- Stops generating when the last word of the input is reached.
- Provides error messages for invalid input and arguments.
- Provides usage information with the `--help` option.

## Usage

The program reads text from standard input:

```bash
cat the_great_gatsby.txt | ./markovchain
