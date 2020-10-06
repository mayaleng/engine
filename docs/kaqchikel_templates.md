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
            "type" : "literal",
            "value" : "{{ .Word3.Translation }}"
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
            "type" : "literal",
            "value" : "{{ .Word3.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{ .Word4.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{ .Word5.Translation }}"
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
            "type" : "literal",
            "value" : "{{ .Word3.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{ .Word4.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{ .Word5.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{ .Word1.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{ .Word2.Translation }}"
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
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.tr ) }}{{if ( StartWithVowel .WordX.Translation ) }}w{{else}}nu{{end}}{{end}}"
}
```
___
Second (2nd) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.tr ) }}{{if ( StartWithVowel .WordX.Translation ) }}aw{{else}}a{{end}}{{end}}"
}
```
___
Third (3rd) Person and Singular Number.
IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.tr ) }}{{if ( StartWithVowel .WordX.Translation ) }}r{{else}}u{{end}}{{end}}
}
```
___
First (1er) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.tr ) }}{{if ( StartWithVowel .WordX.Translation ) }}q{{else}}qa{{end}}{{end}}
}
```
___
Second (2nd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.tr ) }}{{if ( StartWithVowel .WordX.Translation ) }}iw{{else}}i{{end}}{{end}}
}
```
___
Third (3rd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.tr ) }}{{if ( StartWithVowel .WordX.Translation ) }}k{{else}}ki{{end}}{{end}}
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
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.intr ) }}{{if ( StartWithVowel .WordX.Translation) }}in{{else}}i{{end}}{{end}}
}
```
___
Second (2nd) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.intr ) }}{{if ( StartWithVowel .WordX.Translation) }}at{{else}}a{{end}}{{end}}
}
```
___
Third (3rd) Person and Singular Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"S\") ( .WordX.Properties.intr ) }}{{if ( StartWithVowel .WordX.Translation) }}Ø{{else}}Ø{{end}}{{end}}
}
```
___
First (1er) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"1\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.intr ) }}{{if ( StartWithVowel .WordX.Translation) }}oj{{else}}oj{{end}}{{end}}
}
```
___
Second (2nd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"2\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.intr ) }}{{if ( StartWithVowel .WordX.Translation) }}ix{{else}}ix{{end}}{{end}}
}
```
___
Third (3rd) Person and Plural Number.

IF starts with Vowel ELSE starts with Consonant
```
{
    "type":"literal",
    "value":"{{ if and (eq .WordX.Properties.person \"3\") (eq .WordX.Properties.number \"P\") ( .WordX.Properties.intr ) }}{{if ( StartWithVowel .WordX.Translation ) }}e'{{else}}e{{end}}{{end}}
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
## `Adjectives [ADV+ADJ]`
An adjective modified the noun, and only is going to be applied if we write the next ajectives.

MUY - to express a lot
```
{
    "type":"literal",
    "value":"{{ if (eq .WordX.Lemma \"muy\")}}.WordX.Translation .WordX.Translation{{end}}"
}
```
MEDIO - to express more or less
```
{
    "type":"literal",
    "value":"{{ if (eq .WordX.Lemma \"medio\")}}.WordX.Translation+[.WordX.FirstLetter]+oj{{end}}"
}
```
MUCHÍSIMO - to express more than a lot
```
{
    "type":"literal",
    "value":"{{ if (eq .WordX.Lemma \"muchísimo\")}}.WordX.Translation+iläj{{end}}"
}
```
___
## `Numbers [CARD]`
Numbers from 1 to 10 are, and the number root is in bold:
* 1  = **ju**n
* 2  = **kab'**, this have an exception, the number is **ka'i'**
* 3  = **ox**'i
* 4  = **kaj**i'
* 5  = **wo**'o
* 6  = **waq**i'
* 7  = **wuq**u'
* 8  = **waqxaq**i'
* 9  = **b'elej**e'
* 10 = **lajuj**

From 11 to 19 you use the number root from 1 to 9 + number 10

* 11  = **julaj**uj
* 12  = **kab'laj**uj
* 13  = **oxlaj**uj
* 14  = **kajlaj**uj
* 15  = **wolaj**uj
* 16  = **waqlaj**uj
* 17  = **wuqlaj**uj
* 18  = **waqxaqlaj**uj
* 19  = **b'elejlaj**uj

For numbers higher than 20, we apply other prefixes.

The mayan numbers works in base 20, every potencial of 20, we have a different prefix to numbers higher than 19.

* 20 to 399 (20^2 - 1) = **k'al**
* 400 to 7,999 (20^3 - 1) = **q'o'**
* 8,000 to 159,999 (20^4 - 1) = **chuy**
* 160,000 to 3,^199,999 (20^5 - 1) = **k'ala'**

The costruction to numbers greater than 19 are with those prefix and using the 19 roots.

The 19 roots
```
1  = ju
2  = kab
3  = ox
4  = kaj
5  = wo
6  = waq
7  = wuq
8  = waqxaq
9  = b'elej
10 = laj
11 = julaj
12 = kab'laj
13 = oxlaj
14 = kajlaj
15 = wolaj
16 = waqlaj
17 = wuqlaj
18 = waqxaqlaj
19 = b'elejlaj
```

Examples
```
20 = juk'al

ju = 20 / 20 = 1, the root to use is the first
k'al = 20 <= 20 < 400, the prefix to use is the first
```
```
120 = waqk'al

waq = 120 / 20 = 6, the root to use is the sixth
k'al = 20 <= 120 < 400
```
```
3200 = waqxaqk'alq'o

waqxaq = 3200 / 20 / 20 = 8, the root to use is the eigth
ka'l = 3200 > 400, we have to concat the previous prefix
q'o = 400 <= 3200 < 8000
```
```
9600 = kab'lajk'alq'o'chuy

kab'laj = 9600 / 20 / 20 / 20 = 12, the root so use is the twelveth
k'al = 9600 > 400, we have to concat the previous prefix
q'o = 9600 > 8000, we have to concat the previous prefix
chuy = 8000 <= 9600 < 160000
```
Cardinal Numbers
```
{
           "type":"literal",
            "value":"{{ GetKaqchikelNumber .Word1.Lemma \"jun,ka'i',oxi',kaji',wo'o,waqi',wuqu',waqxaqi',b'eleje'\" \"ju,ka,ox,kaj,o,waq,wuq,waqxaq,b'elej,laj,julaj,kab'laj,oxlaj,kajlaj,wolaj,waqlaj,wuqlaj,waqxaqlaj,b'elejlaj\" \"k'al,q'o',chuy,k'ala'\" }}"
 }
```
___
You know how to write individual templates to a rule, now we are going to see how to apply it together to translate short sentences.

These are rules you can load in your database.

[Kaqchikel Rules](kaqchikel_rules.md)