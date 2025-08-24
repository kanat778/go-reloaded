# GO-RELOADED

## Objectives

This program represents a simple **text completion / editing / auto-correction tool** written in Go.

The tool will receive as arguments the name of a file containing a text that needs some modifications (**input**) and the name of the file the modified text should be placed in (**output**).

---

## Supported Modifications

* Every instance of `(hex)` should replace the word before with the **decimal version** of the word (in this case the word will always be a hexadecimal number).
    - Ex: `"1E (hex) files were added"` -> `"30 files were added"`

* Every instance of `(bin)` should replace the word before with the **decimal version** of the word (in this case the word will always be a binary number).
    - Ex: `"It has been 10 (bin) years"` -> `"It has been 2 years"`

* Every instance of `(up)` converts the previous word into **uppercase**.
    - Ex: `"Ready, set, go (up) !"` -> `"Ready, set, GO!"`

* Every instance of `(low)` converts the previous word into **lowercase**.
    - Ex: `"I should stop SHOUTING (low)"` -> `"I should stop shouting"`

* Every instance of `(cap)` capitalizes the previous word.
    - Ex: `"Welcome to the Brooklyn bridge (cap)"` -> `"Welcome to the Brooklyn Bridge"`

    - For `(low)`, `(up)`, `(cap)` if a number appears next to it, like so: `(low, <number>)` it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly.
        - Ex: `"This is so exciting (up, 2)"` -> `"This is SO EXCITING"`

* Every instance of the punctuations `.`, `,`, `!`, `?`, `:` and `;` should be close to the **previous word** and separated from the **next one** by a space.
    - Ex: `"I was sitting over there ,and then BAMM !!"` -> `"I was sitting over there, and then BAMM!!"`

    - Except if there are groups of punctuation like: `...` or `!?`.
        - Ex: `"I was thinking ... You were right"` -> `"I was thinking... You were right"`

* The punctuation mark `'` will always be found with another instance of it and they should be placed to the right and left of the word(s) inside them, without spaces.
    - Ex: `"I am exactly how they describe me: ' awesome '"` -> `"I am exactly how they describe me: 'awesome'"`
    - Ex: `"There is no greater agony than    '    bearing a untold story inside you."` -> `"There is no greater agony than' bearing a untold story inside you."`

* Every instance of `a` should be turned into `an` if the next word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or an `h`.
    - Ex: `"There it was. A amazing rock!"` -> `"There it was. An amazing rock!"`

---

## Usage

```bash
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

## Feedback

If you liked my project, please do not shy to add a ⭐ to the repository.

Made by **kkanat**

