package morse

import (
	"errors"
	"strings"
)

func New() *Morse {
	return &Morse{ ".", "-", " "}
}

type Morse struct {
	// dots and dashes or dits and dahs
	dot      string // .
	dash     string // -
	splitter string // spaces and words by "/" or "|"
	// wordSplitter string  // default is space
	// codeSplitter string // default is space
}

// Input the words and convert it to the Morse code
func (m *Morse) Encode(words string) string {
	upWords := strings.ToUpper(words)
	// TODO Split word by any character not only space
	splitWords := strings.Split(upWords, " ")

	var morseCode string
	for _, w := range splitWords {
		var wordCode string
		for _, v := range string(w) {
			if c, ok := morseMap[string(v)]; ok {
				// TODO Split morse code by any character not only space
				wordCode += c + " "
			}
		}
		// morseCode += wordCode[:len(wordCode)-1] + m.splitter
		morseCode += " " + wordCode + string(m.splitter)
	}

	if m.dot != "." {
		morseCode = strings.Replace(morseCode, ".", string(m.dot), -1)
	}
	if m.dash != "-" {
		morseCode = strings.Replace(morseCode, "-", string(m.dash), -1)
	}

	return morseCode[1:len(morseCode)-1]
}

// Input the Morse code and convert it to the words
func (m *Morse) Decode(codes string) string {
    // Replace 3 spaces to 2 spaces
    codes = strings.Replace(codes, " " + string(m.splitter) + " ", "  ", -1)
	// TODO Split word by any character not only space
	splitCodes := strings.Split(codes, " ")
	mapping := codeMap()
	var words string
	for _, code := range splitCodes {
		// Not conver Space
		if len(code) == 0  || code == m.splitter {
			words += " "
			continue
		}

		if ch, ok := mapping[string(code)]; ok {
            words += ch
		}
	}

	return words
}

func (m *Morse) Reset() *Morse {
	nm := New()
	m.SetDot(nm.dot)
	m.SetDash(nm.dash)
	m.SetSplitter(nm.splitter)
	return m
}

func (m *Morse) SetDot(dot string) error {
	if dot == m.dash || dot == m.splitter {
		return errors.New("dot should be unique!")
	}
	m.dot = dot
	return nil
}

func (m *Morse) SetDash(dash string) error {
	if dash == m.dot || dash == m.splitter {
		return errors.New("dash should be unique!")
	}
	m.dash = dash
	return nil
}

func (m *Morse) SetSplitter(splitter string) error {
	if splitter == m.dot || splitter == m.dash {
		return errors.New("splitter should be unique!")
	}
	m.splitter = splitter
	return nil
}

var morseMap = map[string]string{
	"A":  ".-",   // ascii 65
	"B":  "-...", // ascii 66
	"C":  "-.-.",
	"D":  "-..",
	"E":  ".",
	"F":  "..-.",
	"G":  "--.",
	"H":  "....",
	"I":  "..",
	"J":  ".---",
	"K":  "-.-",
	"L":  ".-..",
	"M":  "--",
	"N":  "-.",
	"O":  "---",
	"P":  ".--.",
	"Q":  "--.-",
	"R":  ".-.",
	"S":  "...",
	"T":  "-",
	"U":  "..-",
	"V":  "...-",
	"W":  ".--",
	"X":  "-..-",
	"Y":  "-.--",
	"Z":  "--..",
	"1":  ".----", // ascii 1
	"2":  "..---", // ascii 2
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	"0":  "-----",
	".":  ".-.-.-", // ascii 46
	":":  "---...",
	",":  "--..--",
	";":  "-.-.-",
	"?":  "..--..",
	"=":  "-...-",
	"'":  ".----.",
	"/":  "-..-.",
	"!":  "-.-.--",
	"-":  "-....-",
	"_":  "..--.-",
	"\"": ".-..-.",
	"(":  "-.--.",
	")":  "-.--.-",
	"()": "-.--.-",
	"$":  "...-..-",
	"&":  ".-...",
	"@":  ".--.-.",
	"+":  ".-.-.",
	// " ":  ".......",
}

var codeMap = func() map[string]string {
	newMap := make(map[string]string)
	for k, v := range morseMap {
		newMap[v] = k
	}
	return newMap
}
