I have created a program that reads a text input from stdin and generates new text using the Markov 
Chain algorithm. The program can generate text by predicting the most probable next word based on 
the previous word, similar to the way predictive text functions on phone keyboards. 
It follows these key points:

By default, the program processes the entire input text and generates output based on the Markov Chain algorithm.
The suffix length is always set to 1 word.
The default prefix length is 2 words, with the starting prefix being the first two words of the input text.
The generated text is limited to a maximum of 100 words, unless otherwise specified.
The program has error handling to notify if any problems occur, and it stops generating text after reaching the maximum word limit or when encountering the last word in the input text.

In addition, the program allows users to:

Set a maximum number of words to generate (with constraints: non-negative and ≤ 10,000).
Example: $ cat the_great_gatsby.txt | ./markovchain -w 10 | cat -e

Specify a starting prefix (which must exist in the input text).
Example: $ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to play" | cat -e


Adjust the prefix length (allowed between 1 and 5 words).
Example: $ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to something funny" -l 3

If any issues arise, such as invalid input or constraints being violated, the program will print an error message indicating the problem.

Additionally, the program is capable of printing usage information. This includes instructions 
on how to use the text generator, such as command-line options, input requirements, and any additional 
features the user needs to know.
Example: ./markovchain --help
