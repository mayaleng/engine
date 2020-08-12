# Deterministic Algorithm

Nowaday we can use Machine Learning to translate sentences from one language to other, using a the `sequence to sequence` model. But in some cases there are a little problem: _not enough training data_.

That was what happened to us. We tried to translate _Spanish_ sentences to _Kaqchikel_ (a _Mayan_ language). But we could not find enough data to train a model. So we make a decision: 

1. Use translation based on rules.

This approach will allow to us generate training data to then be able to build a _se2seq_ model.

In this section we will describe our translation rules.

## Rule structure

This is an example of a rule definition.

```json
{
    "pattern": "VERB,ADV,ADJ",
    "details": [
        {
            "tag": "VERB",
            "type": "",
            "properties": {
                "number": "s"
            }
        },
        {
            "tag": "ADV",
            "type": "",
            "properties": {
                
            }
        },
        {
            "tag": "ADJ",
            "type": "",
            "properties": {
                "number": "S"
            }
        }
    ],
    "output": [
        {
            "type": "literal",
            "value": "{{if and (eq .Word1.Properties.number \"S\") (eq .Word3.Properties.number \"S\")}} {{- \"a\" -}} {{else}} {{- \"r\" -}} {{end}}"
        },
        {
            "type": "literal",
            "value": " "
        },
        {
            "type": "literal",
            "value": "{{ .Word1.Translation }}",
        },
        {
            "type": "literal",
            "value": " "
        },
        {
            "type": "literal",
            "value": "{{ .Word2.Translation }}",
        },
    ]
}
```

## `pattern`

It is a `string` that represents ordered words that senteces MUST contain to be able to use this rule.

e.g
1. `DET,NOUN` e.g _El sol_
2. `DET,NOUN,VERB` e.g _El perro salta_
3. `VERB` e.g _Cantando_

## `details`

It is an `array` of `objects`. Each object represents a member of the `pattern`. Those objects will be used to add extra filter capabilities. For example you could be able to filter a verb based on its _tense_. 

The structure of a single object maps directly with the results of the [`DepPattern`](https://github.com/gamallo/DepPattern/blob/master/doc/tutorialGrammar.pdf) module of `Linguakit`.

Here the known schema:

```json
{
    "tag": "ADJ|ADV|DT|NOUN|VERB|PRO|CONJ|I|P|CARD|DATE",
    "type": "",
    "properties": {
        "mode": "I",
        "tense": "P",
        "<key>": "<value>",
        ...
    }
}
```

## `output`

Represents the rules to follow to generate the translated sentence.

Each ouput element always will have two properties: `type` and `value`, based on those we will generate a new word.

### **type**

There are 2 possible values:
- `literal`: will output exactly the value provieded.
- `.WordX.Translation`: will output the _direct translation_ of the word given in the property `value`.

### **value**

As we mentioned above here you put the value to be used to generate the translated word.

This property has the power to contains **templates** and not only literals.

A literal value could be:
- `Helo!`
- `'`
- `prefix-`
- `-sufix`

Any hard-coded string. 

Using only literal values are not enough. Some languages uses prior word to generate next words in the sentence. So it is strictly needed to have a way to access randomly to the words in the sentences to translate. In the next section **templates** is described.

## Templates

This _syntax_ is used to access to words within the sentence.

A template usage looks like:

`value: "hard coded prefix: {{ .Word1.Lemma }} suffix"`

### Accessors

To use words in the value contet, you only need to use `{{ <property path> }}`.
You can notice that the property path starts with a dot. Immediately after the dot we can use `Word<n>` where `n` is a number from `1..length(details)`. It means that all the words in the array `details` could be accessed via templates.

### Conditional

This is the most useful feature of templates.

You can use logical conditions to generate outputs.

e.g 
```
{{if and (eq .Word1.Lemma "yo") (eq .Word2.Lemma "feliz")}} 😀 {{else}} 🥺 {{end}}
```

The most useful operators are: `if`, `eq`, `or` and `and`. 

This capability is provided by Golang templates, for more information read the offcial [documentation](https://golang.org/pkg/text/template/).
