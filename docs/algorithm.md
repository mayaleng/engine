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
    "pattern": "VERB,ADV,ADJ", // Generated based on the rules
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
            "type": "direct-translation",
            "value": "{{Word1.Lemma}}",
        },
        {
            "type": "literal",
            "value": " "
        },
        {
            "type": "direct-translation",
            "value": "{{Word2.Lemma}}",
        },
    ]
}
```

### `pattern`

It is used to represent the `details` array. Basically it represents the form of the sentences.

e.g
1. `DET+NOUN`
2. `DET+NOUN+VERB`
3. `VERB`

### `details`


The structure of a single word mapping directly with the results of the `DepPattern` module of `Linguakit`.

[Here](https://github.com/gamallo/DepPattern/blob/master/doc/tutorialGrammar.pdf) you have a better understading of the result.

```json
{
    "type": "ADJ|ADV|DT|NOUN|VERB|PRO|CONJ|I|P|CARD|DATE",
    "properties": {
        "mode": "I",
        "<key>": "<value>",
        ...
    }
}
```

**properties** depends on the type of the word.

### `output`

Represents the rules to follow to generate the translated sentence.

Each ouput element always will have one property: `type`, based on the type other properties will be required.

**type**

There are 3 possible values:
- `literal`: will output exactly the value provieded in the property `value`.
- `direct-translation`: will output the direct translation of the word given in the property `value`. You can use accessor to dynamic values as well (`${...}` will be explained later).

**Dynamic accessors**

We provide a simple way to access to dyanmic values.

For example to the word values. As you can see, the `details` property is a list, so the most natural way to access to them is via indexes.

```
${0.lemma}
```
Will be replace with the value of the word dynamically.
