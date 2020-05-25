# Deterministic algorithm

Nowaday we can use Machine Learning to translate sentences from one language to other, using a the `sequence to sequence` model. But in some cases there are a small problem: not enough training data.

This was our situation. We tried to translate _Spanish_ sentences to _Kaqchikel_. But we could not find enough data to train a model. So we make a decision: 

1. Use translation based on rule.

This will allow to us generate training data to then be able to build a model _se2seq_.

In this section we will describe our translation rules.

## Rule structure

This is an example of a definiton rule structure.

```js
{
    "pattern": "DET+NOUN+VERB" // Generated based on the rules
    "words": [
        {
            "type": "DET",
            "properties": {
            }
        },
        {
            "type": "NOUN",
            "properties": {
                "gender": "M"
            }
        },
        {
            "type": "VERB",
            "properties": {
                "tense": "P"
            }
        }
    ],
    "output": [
        {
            "type": "conditional",
            "when": "${0.number} == \"S\" && ${0.person} == \"3\" || ${0.person} == \"3\"" // 0 is the index in the array
            "then": "ri",
            "else": "re"
        },
        {
            "type": "literal",
            "value": "<blank>"
        },
        {
            "type": "direct-translation",
            "value": "${1.lemma}",
        },
        {
            "type": "literal",
            "value": "<blank>"
        },
        {
            "type": "direct-translation",
            "value": "${2.lemma}",
        },
    ]
}
```

### `pattern`

It is used to represent the `words` array. Basically it represents the form of the sentences.

e.g
1. `DET+NOUN`
2. `DET+NOUN+VERB`
3. `VERB`

### `words`


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
- `literal`: will output exactly the value provieded in the property `value`. Special value `<blank>` is exactly the same that ` `.
- `direct-translation`: will output the direct translation of the word given in the property `value`. You can use accessor to dynamic values as well (`${...}` will be explained later).
- `conditional`: this one is more complex. It will generate an output based on `logical` conditions. The logical condition will be provided in the property `when`. Based on the result of the logical operation you can have two possible ways: the consequence (`then`) and the alternative (`alternative`).

**Dynamic accessors**

We provide a simple way to access to dyanmic values.

For example to the word values. As you can see, the `words` property is a list, so the most natural way to access to them is via indexes.

```
${0.lemma}
```
Will be replace with the value of the word dynamically.
