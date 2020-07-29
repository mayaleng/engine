# Kaqchikel Rules

These are rules to construct kaqchikel sentences.

## `[VERB]`
When you have a verb, you need to apply this group of templates, where X in WordX is the number of word and Word is the VERB you want to apply the rule.

**output[0] - Aspect and Time**
```
IF Word.Tense == "S" THEN
    return "x"
ELSE IF Word.Tense == "P" THEN
    return "y"
ELSE IF Word.Tense == "F" THEN
    return "xk"
ELSE IF Word.Tense == "0" THEN
    return "tajin"
ELSE IF
```
**output[1] - Number and Person Indicator - Transitive**
```
IF Word.Transitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN
    return "w"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "nu"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN
    return "qa"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "a"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN
    return "aw"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "a"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN
    return "iw"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "i"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN    
    return "r"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "u"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN    
    return "k"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "ki"
END
```
**output[2] - Number and Person Indicator - Intransitive**
```
IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN
    return "in"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "i"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN
    return "oj"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "oj"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN
    return "at"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "a"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN
    return "ix"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "ix"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN    
    return "Ø"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "Ø"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN    
    return "e'"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "e"
END
```
Template
```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "VERB",
    "details" : [
        {
            "tag" : "VERB",
            "type" : "M"
        }
    ],
    "output" : [ 
        {
            "type":"literal",
            "value":"{{ if (eq .Word1.Properties.tense \"S\" ) }}x{{ else if (eq .Word1.Properties.tense \"P\")}}y{{ else if (eq .Word1.Properties.tense \"F\")}}xk{{ else if (eq .Word1.Properties.tense \"0\")}}tajin{{end}}"
        },
        {
            "type":"literal",
            "value":"{{ if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma ) }}w{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}nu{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}qa{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}aw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}iw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}i{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}r{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}u{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}k{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}ki{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma ) }}in{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}i{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}at{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}e'{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}e{{end}}"
        },
        {
            "type":"direct-translation",
            "value":"{{ .Word1.Lemma }}"
        }
    ]
}
```
___
## `[NOUN]`
When you have a possesive adjective + noun, you need to apply this group of templates, where X in WordX is the number of word and Word is the NOUN you want to apply the rule.

**[0] - Number and Person Indicator - Transitive**
```
IF Word.Person == "1" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN
    return "w"
ELSE IF Word.Person == "1" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "nu"
ELSE IF Word.Person == "1" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN
    return "qa"
ELSE IF Word.Person == "1" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "a"
ELSE IF Word.Person == "2" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN
    return "aw"
ELSE IF Word.Person == "2" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "a"
ELSE IF Word.Person == "2" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN
    return "iw"
ELSE IF Word.Person == "2" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "i"
ELSE IF Word.Person == "3" AND Word.Number == "S" AND Word.StartWithVowel(Word.Lemma) THEN    
    return "r"
ELSE IF Word.Person == "3" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "u"
ELSE IF Word.Person == "3" AND Word.Number == "P" AND Word.StartWithVowel(Word.Lemma) THEN    
    return "k"
ELSE IF Word.Person == "3" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Lemma) THEN
    return "ki"
END
```
Template
```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "DET,NOUN",
    "details" : [ 
        {
            "tag" : "DET",
            "type" : "P"
        },
        {
            "tag" : "NOUN",
            "type" : "C"
        }
    ],
    "output" : [
        {
            "type":"literal",
            "value":"{{ if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma ) }}w{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}nu{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}qa{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}aw{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}iw{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}i{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}r{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}u{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}k{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}ki{{end}}"
        },
        {
            "type":"direct-translation",
            "value":"{{ .Word2.Lemma }}"
        }
    ]
}
```

## `[DET] - Articles`
When you have an article + noun, you need to apply this group of templates, where X in WordX is the number of word and Word is the NOUN you want to apply the rule.

**Article**
```
IF Word.Type == "I" THEN
    return "jun"
ELSE IF Word.Type == "D" THEN
    return "ri jun"
ELSE IF Word.Type == "A" THEN
    return "ri"
END
```
Template
```
/*1*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "DET,NOUN",
    "details" : [ 
        {
            "tag" : "DET",
            "type": "I"
        },
        {
            "tag" : "NOUN",
            "type" : "C"
        }
    ],
    "output" : [
        {
           "type":"literal",
            "value":"{{ if (eq .Word1.Type \"I\" ) }}jun{{end}}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
           "type":"direct-translation",
            "value":"{{ .Word2.Lemma }}"
        }
    ]
}
/*2*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "DET,NOUN",
    "details" : [ 
        {
            "tag" : "DET",
            "type": "D"
        },
        {
            "tag" : "NOUN",
            "type" : "C"
        }
    ],
    "output" : [
        {
           "type":"literal",
            "value":"{{ if (eq .Word1.Type \"D\")}}ri jun{{end}}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
           "type":"direct-translation",
            "value":"{{ .Word2.Lemma }}"
        }
    ]
}
/*3*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "DET,NOUN",
    "details" : [ 
        {
            "tag" : "DET",
            "type": "A"
        },
        {
            "tag" : "NOUN",
            "type" : "C"
        }
    ],
    "output" : [
        {
           "type":"literal",
            "value":"{{ if (eq .Word1.Type \"A\")}}ri{{end}}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
           "type":"direct-translation",
            "value":"{{ .Word2.Lemma }}"
        }
    ]
}
```
