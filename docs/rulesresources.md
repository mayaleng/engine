# Rules Examples

We provide some simple rules you can use to try the translator:

---

*  Rule with direct translation

**Phrase:** Estoy muy feliz

Linguakit response:

```
Estoy: 
    tag:  VERB
    type: A
muy: 
    tag:  ADV
    type: G
feliz: 
    tag:  ADJ
    type: Q
```

Rule to translate that kind of phrase
```json
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "VERB,ADV,ADJ",
    "details" : [ 
        {
            "tag" : "VERB",
            "type" : "A"
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
            "type" : "literal",
            "value" : "{{ .Word3.Translation }}"
        }, 
        {
            "type" : "literal",
            "value" : " "
        }, 
        {
            "type" : "literal",
            "value" : "{{ .Word2.Translation }}"
        }, 
        {
            "type" : "literal",
            "value" : "{{ .Word1.Translation }}"
        }
    ]
}
```
---
* Rule with direct and conditional translation
  
**Phrase:** Salta mucho

Linguakit response:

```
Salta:
    tag:  VERB
    type: M
mucho: 
    tag:  ADV
    type: G
```

Rule to translate that kind of phrase
```json
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "VERB,ADV",
    "details" : [ 
        {
            "tag" : "VERB",
            "type" : "M"
        }, 
        {
            "tag" : "ADV",
            "type" : "G"
        }
    ],
    "output" : [ 
        {
            "type" : "literal",
            "value" : "{{ .Word1.Translation }}"
        }, 
        {
            "type" : "literal",
            "value" : "{{if (eq .Word2.Lemma \"mucho\") }} q'uiy {{- else}} _ {{end}}"
        }
    ]
}
```
---
* Rule with direct, conditional translation and literal value

**Phrase:** Jorge juega con mi pelota

Linguakit response:

```
Jorge:
    tag:  NOUN
    type: P
juega: 
    tag:  VERB
    type: M
con:
    tag:  PRP
    type: P
mi:
    tag:  DET
    type: P
pelota:
    tag:  NOUN
    type: C
```

Rule to translate that kind of phrase
```json
{
    "source_language" : "espaol",
    "target_language" : "kaqchikel",
    "pattern" : "NOUN,VERB,PRP,DET,NOUN",
    "details" : [ 
        {
            "tag" : "NOUN",
            "type" : "P"
        }, 
        {
            "tag" : "VERB",
            "type" : "M"
        },
        {
            "tag" : "PRP",
            "type" : "P"
        },
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
            "type" : "literal",
            "value" : "{{ .Word2.Translation }}"
        }, 
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{if (eq .Word5.Properties.number \"S\")nu{{end}}"
        },
        {
            "type" : "literal",
            "value" : "{{{ .Word5.Translation }}"
        },
        {
            "type" : "literal",
            "value" : " "
        },
        {
            "type" : "literal",
            "value" : "{{ .Word1.Lemma }}"
        }
    ]
```

You can see in that case we don't need all the words in the original sentence to construct the output.

---