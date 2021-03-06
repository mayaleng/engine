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
IF Word.Transitive AND Word.Person == "1" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN
    return "w"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "nu"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN
    return "qa"
ELSE IF Word.Transitive AND Word.Person == "1" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN
    return "aw"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN
    return "iw"
ELSE IF Word.Transitive AND Word.Person == "2" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "i"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN    
    return "r"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "u"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN    
    return "k"
ELSE IF Word.Transitive AND Word.Person == "3" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "ki"
END
```
**output[1.2] - Number and Person Indicator - Intransitive**
```
IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN
    return "in"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "i"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN
    return "oj"
ELSE IF Word.Intransitive AND Word.Person == "1" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "oj"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN
    return "at"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN
    return "ix"
ELSE IF Word.Intransitive AND Word.Person == "2" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "ix"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN    
    return "Ø"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "Ø"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN    
    return "e'"
ELSE IF Word.Intransitive AND Word.Person == "3" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "e"
END
```

* Example: **camina**
* Result: **yupalaqinïk**

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
            "value":"{{ if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word1.Translation ) }}w{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word1.Translation )}}nu{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word1.Translation )}}qa{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word1.Translation )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word1.Translation )}}aw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word1.Translation )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word1.Translation )}}iw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word1.Translation )}}i{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word1.Translation )}}r{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word1.Translation )}}u{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word1.Translation )}}k{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word1.Translation )}}ki{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word1.Translation ) }}in{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word1.Translation )}}i{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word1.Translation )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word1.Translation )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word1.Translation )}}at{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word1.Translation )}}a{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word1.Translation )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word1.Translation )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word1.Translation )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word1.Translation )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word1.Translation )}}e'{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word1.Translation )}}e{{end}}"
        },
        {
            "type":"literal",
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

* Example: **le he regado / he regado**
* Result: **nukiranïk**

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
            "value":"{{ if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"S\") ( StartsWithVowel .Word3.Translation ) }}w{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"S\") ( StartsWithConsonant .Word3.Translation )}}nu{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"P\") ( StartsWithVowel .Word3.Translation )}}qa{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"1\") (eq .Word2.Properties.number \"P\") ( StartsWithConsonant .Word3.Translation )}}a{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"S\") ( StartsWithVowel .Word3.Translation )}}aw{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"S\") ( StartsWithConsonant .Word3.Translation )}}a{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"P\") ( StartsWithVowel .Word3.Translation )}}iw{{ else if and ( .Word2.Properties.tr ) (eq .Word2.Properties.person \"2\") (eq .Word2.Properties.number \"P\") ( StartsWithConsonant .Word3.Translation )}}i{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"S\") ( StartsWithVowel .Word3.Translation )}}r{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"S\") ( StartsWithConsonant .Word3.Translation )}}u{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"P\") ( StartsWithVowel .Word3.Translation )}}k{{ else if and ( .Word3.Properties.tr ) (eq .Word2.Properties.person \"3\") (eq .Word2.Properties.number \"P\") ( StartsWithConsonant .Word3.Translation )}}ki{{end}}"
        },
        {
            "type":"literal",
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
            "value":"{{ if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word2.Translation ) }}w{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word2.Translation )}}nu{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word2.Translation )}}qa{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word2.Translation )}}a{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word2.Translation )}}aw{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word2.Translation )}}a{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word2.Translation )}}iw{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word2.Translation )}}i{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word2.Translation )}}r{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word2.Translation )}}u{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word2.Translation )}}k{{ else if and ( .Word2.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word2.Translation )}}ki{{end}}"
        },
        {
            "type":"literal",
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
IF Word.Person == "1" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN
    return "w"
ELSE IF Word.Person == "1" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "nu"
ELSE IF Word.Person == "1" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN
    return "qa"
ELSE IF Word.Person == "1" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Person == "2" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN
    return "aw"
ELSE IF Word.Person == "2" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "a"
ELSE IF Word.Person == "2" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN
    return "iw"
ELSE IF Word.Person == "2" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "i"
ELSE IF Word.Person == "3" AND Word.Number == "S" AND StartsWithVowel(Word.Translation) THEN    
    return "r"
ELSE IF Word.Person == "3" AND Word.Number == "S" AND StartsWithConsonant(Word.Translation) THEN
    return "u"
ELSE IF Word.Person == "3" AND Word.Number == "P" AND StartsWithVowel(Word.Translation) THEN    
    return "k"
ELSE IF Word.Person == "3" AND Word.Number == "P" AND StartsWithConsonant(Word.Translation) THEN
    return "ki"
END
```
* Example: **su casa**
* Result: **uja**

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
            "value":"{{ if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word2.Translation ) }}w{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word2.Translation )}}nu{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word2.Translation )}}qa{{ else if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word2.Translation )}}a{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word2.Translation )}}aw{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word2.Translation )}}a{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word2.Translation )}}iw{{ else if and (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word2.Translation )}}i{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word2.Translation )}}r{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word2.Translation )}}u{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithVowel .Word2.Translation )}}k{{ else if and (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( StartsWithConsonant .Word2.Translation )}}ki{{end}}"
        },
        {
            "type":"literal",
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
* Example: **ese perro**
* Result: **ri jun tz'i'**

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
           "type":"literal",
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
           "type":"literal",
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
           "type":"literal",
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
    return {Adjective}+{GetFirstLetter}+oj
ELSE IF Word == MUCHÍSIMO Then
    return {Adjective}+iläj
END
```

* Example: **muy caliente**
* Result: **k'atän k'atän**

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
            "value":"{{ if (eq ( ToLower .Word1.Lemma ) \"muy\") }}{{.Word2.Translation}} {{.Word2.Translation}}{{ else if (eq ( ToLower .Word1.Lemma ) \"medio\") }}{{ .Word2.Translation }}{{ GetFirstLetter .Word2.Translation }}oj{{end}}"
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
___
## `[CARD] - Numbers`
We use specific text to numbers from 1 to 10, from 11 to onwards, we are going to apply these rules.

```
IF Word >= 1 AND Word <= 9 THEN
    number
ELSE IF Word >= 10 AND Word <= 19 THEN
    [root]+uj
ELSE IF Word >= 20 THEN
    GetUnity = MOD_10(number)

    While number >= 1 Then
        number = number / 20
        concat [base_20_root]
    End While

    concat [root]
    concat GetUnity
END IF
```

* Example: **3200**
* Result: **waqxaqk'alq'o'**

Template
```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "CARD",
    "details" : [
        {
            "tag" : "CARD"
        }
    ],
    "output" : [
        {
           "type":"literal",
            "value":"{{ GetKaqchikelNumber .Word1.Lemma \"jun,ka'i',oxi',kaji',wo'o,waqi',wuqu',waqxaqi',b'eleje'\" \"ju,ka,ox,kaj,o,waq,wuq,waqxaq,b'elej,laj,julaj,kab'laj,oxlaj,kajlaj,wolaj,waqlaj,wuqlaj,waqxaqlaj,b'elejlaj\" \"k'al,q'o',chuy,k'ala'\" }}"
        }
    ]
}
```
___
## `Re-Use Rules`

Using VERB rule.

* Example: **Camina**
* Result: **yisik'inïk**

```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "VERB",
    "details" : [
        {
            "tag" : "VERB",
            "type" : "M",
            "properties":{
                "f":"true"
            }
        }
    ],
    "output" : [ 
        {
            "type":"predefined",
            "value":"5f696a6d84b7dc8b08728565",
            "start_word":"0"
        }
    ]
}
```
Using VERB and Possesive NOUN rule.

* Example: **Mi perro camina**
* Result: **yupalaqinïk nutz'i'**

Where the first ID is the ObjectID from `[VERB]` rule and the second one is from `[NOUN]` rule.
```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "DET,NOUN,VERB",
    "details" : [
        {
            "tag" : "DET",
            "type" : "P"
        },
        {
            "tag" : "NOUN",
            "type" : "C"
        },
        {
            "tag" : "VERB",
            "type" : "M"
        }
    ],
    "output" : [ 
        {
            "type":"predefined",
            "value":"5f696a6d84b7dc8b08728565",
            "start_word":"2"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type":"predefined",
            "value":"5f696aa384b7dc8b08728568",
            "start_word":"0"
        }
    ]
}
```

* Example: **Mi casa es roja**
* Result: **käq nuja**

Where the ID is the ObjectID from `[NOUN]` rule and a normal `literal` type template to translate.

```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "DET,NOUN,VERB,ADJ",
    "details" : [
        {
            "tag" : "DET",
            "type" : "P"
        },
        {
            "tag" : "NOUN",
            "type" : "C"
        },
        {
            "tag" : "VERB",
            "type" : "S"
        },
        {
            "tag" : "ADJ",
            "type" : "Q"
        }
    ],
    "output" : [ 
        {
            "type" : "literal",
            "value" : "{{ .Word4.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type":"predefined",
            "value":"5f696aa384b7dc8b08728568",
            "start_word":"0"
        }
    ]
}
```
