# Kaqchikel Templates

Here, you can find how to write kaqchikel templates to mayaleng translator rules.

We are going to start with an example, we have a sentence and 3 forms that this sentence can be aplied, and depending which rule you configure, it will be the matched rule and the result. The translator have to choose the 3rd rule because its has the majority of details to match with that rule.

Spanish sentence: **Doña Ixim regó una planta**

Kaqchikel result: **xukiranïk jun q'ayïs don Ixim**

* Rule only with TAG on details
```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "NOUN,NOUN,VERB,DET,NOUN",
    "details" : [ 
        {
            "tag" : "NOUN"
        },
        {
            "tag" : "NOUN"
        },
        {
            "tag" : "VERB"
        },
        {
            "tag" : "DET"
        },
        {
            "tag" : "NOUN"
        }
    ],
    "output" : [ 
        {
            "type" : "literal",
            "value" : "{{ if (eq .Word3.Properties.tense \"S\") }}x{{end}}"
        },
        {
            "type" : "literal",
            "value" : "{{ if and (eq .Word3.Properties.person \"3\") (eq .Word3.Properties.number \"S\") }}u{{end}}"
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word3.Lemma }}"
        }
    ]
}
```
* Rule with TAG and TYPE
```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "NOUN,NOUN,VERB,DET,NOUN",
    "details" : [ 
        {
            "tag" : "NOUN",
            "type": "C"
        },
        {
            "tag" : "NOUN",
            "type": "P"
        },
        {
            "tag" : "VERB",
            "type": "M"
        },
        {
            "tag" : "DET",
            "type": "I"
        },
        {
            "tag" : "NOUN",
            "type": "C"
        }
    ],
    "output" : [ 
        {
            "type" : "literal",
            "value" : "{{ if (eq .Word3.Properties.tense \"S\") }}x{{end}}"
        },
        {
            "type" : "literal",
            "value" : "{{ if and (eq .Word3.Properties.person \"3\") (eq .Word3.Properties.number \"S\") }}u{{end}}"
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word3.Lemma }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word4.Lemma }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word5.Lemma }}"
        }
    ]
}
```
* Rule with TAG, TYPE and some PROPERTIES
```
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "NOUN,NOUN,VERB,DET,NOUN",
    "details" : [ 
        {
            "tag" : "NOUN",
            "type": "C"
        },
        {
            "tag" : "NOUN",
            "type": "P"
        },
        {
            "tag" : "VERB",
            "type": "M",
            "Properties":{
                "tense":"S",
                "person":"3"
            }
        },
        {
            "tag" : "DET",
            "type": "I"
        },
        {
            "tag" : "NOUN",
            "type": "C"
        }
    ],
    "output" : [ 
        {
            "type" : "literal",
            "value" : "{{ if (eq .Word3.Properties.tense \"S\") }}x{{end}}"
        },
        {
            "type" : "literal",
            "value" : "{{ if and (eq .Word3.Properties.person \"3\") (eq .Word3.Properties.number \"S\") ( .Word3.Properties.tr ) }}u{{end}}"
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word3.Lemma }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word4.Lemma }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word5.Lemma }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word1.Lemma }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "direct-translation",
            "value" : "{{ .Word2.Lemma }}"
        }
    ]
}
```
___

## **How to write kaqchikel rules with literal type and their format values**

## `Aspect and Time [VERB]`

NOTE: Where X in WordX is the number of word you want to apply the rule.

These rule apply for verbs, to indicate the tense and subject (using Number and Person indicators). It is mandatory use it to construct a verb.
___
Past tense
```
{
    "type":"literal":
    "value":"{{ if (eq .WordX.Properties.tense \"S\")}}x{{end}}"
}
```
___
Present and Immediate Future
```
{
    "type":"literal":
    "value":"{{ if (eq .WordX.Properties.tense \"P\")}}y{{end}}"
}
```
___
Future
```
{
    "type":"literal":
    "value":"{{ if (eq .WordX.Properties.tense \"F\")}}xk{{end}}"
}
```
___
Progressive
```
{
    "type":"literal":
    "value":"{{ if (eq .WordX.Properties.tense \"0\")}}tajin{{end}}"
}
```
___


## `Number and Person indicators to NOUNS and VERBS [VERB]-[NOUN]`

NOTE: Where X in WordX is the number of word you want to apply the rule.

These templates has the format
```
IF Word.Person == A AND Word.Number == B AND Word.Transitive THEN
    IF Word.Vowel THEN
        RESULT = Y
    ELSE
        RESULT = Z
    END
END
```

* **JUEGO A:** Transitive Verbs and any Possesive Adjective + Noun

These rules apply when you use a Possesive Adjective english: (my, your, his, etc), español: (mi, tu, suyo, etc) and always for transitive verbs, to indicate the tense (using Aspect and Time) and subject of the verb; there are rules when the word starts with consonant and vowel. It is mandatory use it to construct a verb.
___
First (1er) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.tr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}w{{else}}nu{{end}}{{end}}"
}
```
___
Second (2nd) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.tr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}aw{{else}}a{{end}}{{end}}"
}
```
___
Third (3rd) Person and Singular Number.
IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.tr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}r{{else}}u{{end}}{{end}}
}
```
___
First (1er) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.tr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}q{{else}}qa{{end}}{{end}}
}
```
___
Second (2nd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.tr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}iw{{else}}i{{end}}{{end}}
}
```
___
Third (3rd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.tr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}k{{else}}ki{{end}}{{end}}
}
```


* **JUEGO B:** Intransitive Verbs

These rules apply always for intransitive verbs, to indicate the tense (using Aspect and Time) and subject of the verb; there are rules when the verb starts with consonant and vowel. It is mandatory use it to construct a verb.
___
First (1er) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.intr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}in{{else}}i{{end}}{{end}}
}
```
___
Second (2nd) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.intr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}at{{else}}a{{end}}{{end}}
}
```
___
Third (3rd) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.intr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}Ø{{else}}Ø{{end}}{{end}}
}
```
___
First (1er) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.intr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}oj{{else}}oj{{end}}{{end}}
}
```
___
Second (2nd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.intr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}ix{{else}}ix{{end}}{{end}}
}
```
___
Third (3rd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.intr ) }}{{if ( .WordX.StartWithVowel .WordX.Lemma ) }}e'{{else}}e{{end}}{{end}}
}
```
___

## `Articles [DET]`
The articles are aplied through the TYPE property.

Indefinite - unknown
```
{
    "type":"literal",
    "value":"{{ if (eq .WordX.Type \"I\")}}jun{{end}}"
}
```
Demonstrative - identified but unknown
```
{
    "type":"literal",
    "value":"{{ if (eq .WordX.Type \"D\")}}ri jun{{end}}"
}
```
Known Article 
```
{
    "type":"literal",
    "value":"{{ if (eq .WordX.Type \"A\")}}ri{{end}}"
}
```
___
You know how to write individual templates to a rule, now we are going to see how to apply it together to translate short sentences.

These are rules you can load in your database.

[Kaqchikel Rules](kaqchikel_rules.md)