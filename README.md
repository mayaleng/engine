# Translator Engine 

[![Build Status](https://travis-ci.org/mayaleng/engine.svg?branch=master)](https://travis-ci.org/mayaleng/engine) [![codecov](https://codecov.io/gh/mayaleng/engine/branch/master/graph/badge.svg)](https://codecov.io/gh/mayaleng/engine)

This project aim to provide an HTTP interface to make Mayan translations.

We use a database that store:
1. Word translations
2. Languages
3. Rules

That data is used to determine the best translation of a sentence.

## Table of contets

- [Translator Engine](#translator-engine)
  - [Table of contets](#table-of-contets)
  - [Motivation](#motivation)
  - [Getting started](#getting-started)
    - [Installation](#installation)
  - [Documentation](#documentation)
  - [Examples](#examples)
  - [Contributing](#contributing)

## Motivation

Our vision is to help reducing the human communication gap.

We don't want to reinvent Google Translate or a similar tool, we want to offer an easy and generic way to incorporate languages that are dying because the lack of importance on them (e.g Mayan languages in Guatemala).

At the beginning of this project we tried to use Machine Learning to use some famous and effective algorithms to create models for language translation. But we didn't foudn enough data to train the models. So that we got the idea to create a platform capable to create translation based on rules to then generate huge amout of data that then could be used to train AI models.

This idea moved us to create this repo, and now we are working to make it happen.

## Getting started

This project was made in [`Golang`](https://golang.org), so, to run this locally you need to have at lest the version `1.14`.

Other important dependency is [`Linguakit`](https://github.com/citiususc/Linguakit) used for the sentence analysis in Spanish.

For some dependencies you will need [`Docker`](https://docs.docker.com/get-docker/) and [`Docker compose`](https://docs.docker.com/compose/). This will avoid the work of install locally dependencies like databases.

### Installation

As we mentioned we use Go for this project. So, to install locally first you need to install the dependencies.

```
go mod download
```

Other important thing is to install Linguakit. [`Here`](https://github.com/citiususc/Linguakit) is the complete guide.

```
git clone  https://github.com/citiususc/Linguakit
cd Linguakit
sudo make deps
sudo make install
sudo make test-me
```

Last thing is to run the provided `run` script.

```
./run.sh
```

You will see somethig like
```
2020/05/22 10:29:17 Version 1.0.0 built at: 2020-05-22 10:29:17.572036978 -0600 CST m=+0.000292569
2020/05/22 10:29:17 Trying to listen at http://localhost:8080
```

## Documentation

Here you have more documentation to deep dive into some specific topics.

1. [Deterministic translation algorithm](docs/algorithm.md)
2. [Rule examples](docs/rulesresources.md)
3. [Postman examples API](docs/MayaLeng.postman_collection.json): here you can import some examples for Postman and you only need to change the port if run in another.

## Examples
To make a translation you can execute an `HTTP POST`

```
curl -L -X POST 'http://localhost:8080/v1/translate' \
-H 'Content-Type: application/json' \
--data-raw '{
	"from": "espaol",
	"to": "kaqchikel",
	"phrase": "Estoy muy feliz."
}'
```

## Contributing

Contributions are welcome! We want to build a system that can help not only in Guatemalan, but in every place as possible.

We follw some some guidelines as standard:
1. **Code style**. The official style that go recommend. [More info](https://golang.org/doc/effective_go.html)
2. **Commints**: The [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) 
3. **Branches**. The [git-flow](https://danielkummer.github.io/git-flow-cheatsheet/) style
4. **Versions**: The [sem-ver](https://semver.org/) definition

You are free to make  `PR`s as you want and we will love to accept them.

---
Power by [mayaleng.org](https://docs.mayaleng.org)
