### THIS IS NOW ARCHIVED
I found my code messy and wrote a [short perl script](painseeker.pl) that does most of what I wanted this program to do.

# painseeker
Filter comments out of source files on the off chance you might find some expressing the developer's confusion towards the weird behaviour of their code.


### TODO
- Allow reading various files (or search through files recursively if given a directory).
- Detect languages by file name and set the test strings to match the style of comments in that language.
- Properly check for lines with `// /*` or `/* //` kinds of conflicts instead of ignoring them.
- Check if they're inside of a string or multiline string (it is very unlikely that I'll implement this)
