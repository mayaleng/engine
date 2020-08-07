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
**output[1.1] - Number and Person Indicator - Transitive**
```
IF Word.Transitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN
    return "w"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "nu"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN
    return "qa"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN
    return "aw"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN
    return "iw"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
    return "i"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN    
    return "r"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "u"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN    
    return "k"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
    return "ki"
END
```
**output[1.2] - Number and Person Indicator - Intransitive**
```
IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN
    return "in"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "i"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN
    return "oj"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
    return "oj"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN
    return "at"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN
    return "ix"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
    return "ix"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN    
    return "Ø"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "Ø"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN    
    return "e'"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
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
            "value":"{{ if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel ) }}w{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant )}}nu{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel )}}qa{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel )}}aw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel )}}iw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant )}}i{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel )}}r{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant )}}u{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel )}}k{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant )}}ki{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel ) }}in{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant )}}i{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel )}}at{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant )}}a{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel )}}e'{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant )}}e{{end}}"
        },
        {
            "type":"direct-translation",
            "value":"{{ .Word1.Translation }}"
        }
    ]
}
```
To express verbs we have other form, a main verb and an auxiliary verb, we use the auxiliary verb to analyze the phrase and translate the second verb, on this scenario we don't use the time aspect; like this (for these examples the phrase will be in spanish because the translator works with spanish to kaqchikel):

* He ayudado        (1st person singular)
* Le he ayudado     (1st person singular)
* Han ayudado       (3rd person plural)
* Le han ayudado    (3rd person plural)

There are two patterns to apply this case but the output is the same.

Templates:
```
/*1*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "PRO,VERB,VERB",
    "details" : [
        {
            "tag" : "PRO",
            "type" : "P"
        },
        {
            "tag" : "VERB",
            "type" : "A"
        },
        {
            "tag" : "VERB",
            "type" : "M"
        }
    ],
    "output" : [
        {
            "type":"literal",
            "value":"{{ if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"S\") ( .Word3.StartWithVowel ) }}w{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"S\") ( .Word3.StartWithConsonant )}}nu{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"P\") ( .Word3.StartWithVowel )}}qa{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"P\") ( .Word3.StartWithConsonant )}}a{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"S\") ( .Word3.StartWithVowel )}}aw{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"S\") ( .Word3.StartWithConsonant )}}a{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"P\") ( .Word3.StartWithVowel )}}iw{{ else if and ( .Word2.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"P\") ( .Word3.StartWithConsonant )}}i{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"S\") ( .Word3.StartWithVowel )}}r{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"S\") ( .Word3.StartWithConsonant )}}u{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"P\") ( .Word3.StartWithVowel )}}k{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"P\") ( .Word3.StartWithConsonant )}}ki{{end}}"
        },
        {
            "type":"direct-translation",
            "value":"{{ .Word3.Translation }}"
        }
    ]
}
```
```
/*2*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "VERB,VERB",
    "details" : [
        {
            "tag" : "VERB",
            "type" : "A"
        },
        {
            "tag" : "VERB",
            "type" : "M"
        }
    ],
    "output" : [ 
        {
            "type":"literal",
            "value":"{{ if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithVowel ) }}w{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithConsonant )}}nu{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithVowel )}}qa{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithConsonant )}}a{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithVowel )}}aw{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithConsonant )}}a{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithVowel )}}iw{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithConsonant )}}i{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithVowel )}}r{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithConsonant )}}u{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithVowel )}}k{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithConsonant )}}ki{{end}}"
        },
        {
            "type":"direct-translation",
            "value":"{{ .Word2.Translation }}"
        }
    ]
}
```
___
## `[NOUN]`
When you have a possesive adjective + noun, you need to apply this group of templates, where X in WordX is the number of word and Word is the NOUN you want to apply the rule.

**[0] - Number and Person Indicator**
```
IF Word.Person == "1" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN
    return "w"
ELSE IF Word.Person == "1" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "nu"
ELSE IF Word.Person == "1" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN
    return "qa"
ELSE IF Word.Person == "1" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Person == "2" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN
    return "aw"
ELSE IF Word.Person == "2" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Person == "2" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN
    return "iw"
ELSE IF Word.Person == "2" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
    return "i"
ELSE IF Word.Person == "3" AND Word.Number == "S" AND Word.StartWithVowel(Word.Translation) THEN    
    return "r"
ELSE IF Word.Person == "3" AND Word.Number == "S" AND Word.StartWithConsonant(Word.Translation) THEN
    return "u"
ELSE IF Word.Person == "3" AND Word.Number == "P" AND Word.StartWithVowel(Word.Translation) THEN    
    return "k"
ELSE IF Word.Person == "3" AND Word.Number == "P" AND Word.StartWithConsonant(Word.Translation) THEN
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
            "value":"{{ if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithVowel ) }}w{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithConsonant )}}nu{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithVowel )}}qa{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithConsonant )}}a{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithVowel )}}aw{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithConsonant )}}a{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithVowel )}}iw{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithConsonant )}}i{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithVowel )}}r{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word2.StartWithConsonant )}}u{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithVowel )}}k{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word2.StartWithConsonant )}}ki{{end}}"
        },
        {
            "type":"direct-translation",
            "value":"{{ .Word2.Translation }}"
        }
    ]
}
```
___
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
            "value":"{{ .Word2.Translation }}"
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
            "value":"{{ .Word2.Translation }}"
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
            "value":"{{ .Word2.Translation }}"
        }
    ]
}
```
___
## `[ADJ] - Adjectives`
An adjective has its own significant and we are going to apply these rules for 3 expressions, `MUY`, `MEDIO` and `MUCHÍSIMO`, and the patterns we have is `ADV+ADJ` to `MUY` and `MEDIO`, `NOUN+ADV+ADJ` to `MUCHÍSIMO`. If you are saying something about the noun, the adjective precedes the noun.

**Adjective**
```
IF Word == MUY Then
    return {Adjective}+{Adjective}
ELSE IF Word == MEDIO Then
    return {Adjective}+{FirstLetter}+oj
ELSE IF Word == MUCHÍSIMO Then
    return {Adjective}+iläj
END
```
Template
```
/*MUY - MEDIO*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "ADV,ADJ",
    "details" : [ 
        {
            "tag" : "ADV",
            "type": "G"
        },
        {
            "tag" : "ADJ",
            "type" : "Q"
        }
    ],
    "output" : [
        {
           "type":"literal",
            "value":"{{ if (eq .Word1.ToLower \"muy\") }}{{.Word2.Translation}} {{.Word2.Translation}}{{ else if (eq .Word1.ToLower \"medio\") }}{{ .Word2.Translation }}{{ .Word2.FirstLetter }}oj{{ else if (eq .Word1.ToLower \"muchísimo\")}}{{end}}"
        }
    ]
}
/*MUCHÍSIMO*/
/*accent*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "NOUN,ADV,ADJ",
    "details" : [ 
        {
            "tag" : "NOUN",
            "type": "P",
            "properties":{
                "lemma":"muchísimo"
            }
        },
        {
            "tag" : "ADV",
            "type" : "G"
        },
        {
            "tag" : "ADJ",
            "type" : "Q"
        }
    ],
    "output" : [
        {
           "type":"literal",
           "value":"{{ .Word3.Translation }}iläj"
        }
    ]
}
/*without accent*/
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "NOUN,ADV,ADJ",
    "details" : [ 
        {
            "tag" : "NOUN",
            "type": "P",
            "properties":{
                "lemma":"muchisimo"
            }
        },
        {
            "tag" : "ADV",
            "type" : "G"
        },
        {
            "tag" : "ADJ",
            "type" : "Q"
        }
    ],
    "output" : [
        {
           "type":"literal",
           "value":"{{ .Word3.Translation }}iläj"
        }
    ]
}
```