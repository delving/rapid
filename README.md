## NOTICE: this repository is no longer actively maintained.

This functionality is now being developed in the [Hub3](https://github.com/delving/hub3) repository.

<p align="center">
  <!--<img alt="GoReleaser Logo" src="https://avatars2.githubusercontent.com/u/24697112?v=3&s=200" height="140" />-->
  <h3 align="center">Delving RAPID</h3>
  <p align="center">Linked Open Data microservice</p>
  <p align="center">
    <a href="https://github.com/delving/rapid/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/delving/rapid.svg?style=flat-square"></a>
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/delving/rapid"><img alt="Travis" src="https://img.shields.io/travis/delving/rapid/master.svg?style=flat-square"></a>
    <a href='https://coveralls.io/github/delving/rapid?branch=develop'><img src='https://coveralls.io/repos/github/delving/rapid/badge.svg?branch=develop' alt='Coverage Status' /></a>

    <a href="https://goreportcard.com/report/github.com/delving/rapid"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/delving/rapid?style=flat-square"></a>
    <!--<a href="http://godoc.org/github.com/delving/rapid"><img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>-->
    <a href="https://github.com/goreleaser"><img alt="Powered By: GoReleaser" src="https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square"></a>
  </p>
</p>


RAPID is a recursive acronym that stands for *Rapid API Delving*. 

The goal of RAPID is to provide *an API Framework that makes it easy and predictable for webdevelopers to work with arbitrarily structured RDF and leverage semantic network technology*.

The core functionality that it aims to provide can be summarised by the acronym *SILAS*:

* **S**PARQL proxy
* **I**ndex RDF
* **L**inked Open Data Resolver
* **A**ggregate and transform RDF
* **S**earch RDF

Part of the design is to require as little external dependencies outside the compiled *Golang* binary as possible. 

## Project Status

Rapid v1.0 will be released in July 2018. Check the [Changelog] for the full details.

[Changelog]:https://github.com/delving/rapid/blob/master/CHANGELOG.md

## Table of Contents

* [Getting Started](#getting-started)
* [Developing](#developing)
    - [Branching Model](#branching-model)
    - [Dependency Management](#dependency-management)
* [License](#license)

## Getting Started 

RAPID is written in Golang, so you have to setup your Golang environment first, see [Golang Installation].

After that you can glone it from bitbucket:

    $ git clone git@github.com:delving/rapid.git $GOPATH/src/github.com/delving

Or use `go get`

    $ go get github.com/delving/rapid

In order to get the dependencies, install [dep]:

    $ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    $ dep ensure

Start the server with the default configuration.

    $ rapid http

## Developing

### Branching Model

We use the [GitFlow](https://github.com/nvie/gitflow) branching model. Instead of merging features and releases locally, you should use `git flow feature publish <your feature>` and then create a pull-request.

### Dependency Management

We have decided to use [dep] for vendoring. All pinned dependencies are stored in the `./vendor` directory.

Here are some basic commands to work with dep.

Install the dependencies and revisions listed in the lock file into the vendor directory. If no lock file exists an update is run.

    $ dep ennsure

Install the latest dependencies into the vendor directory matching the version resolution information. The complete dependency tree is installed. A lock file is created from the final output.

    $ dep update

Add a new dependency to the `Gopkg.toml`, install the dependency, and re-resolve the dependency tree. Optionally, put a version after an anchor.

    $ dep ensure --add github.com/foo/bar

or 

    $ dep ensure --add github.com/foo/bar#^1.2.3

### System Dependencies

Currently, RAPID depends on a few external dependencies.

- A search engine: We currently support the [Apache Lucene]-based search-engine [ElasticSearch] 5.
- A triple store (optional): We currently support [Apache Fuseki] out of the box. But any Triple store with a [SPARQL 1.1] and [SPARQL Update 1.1] endpoint should work.

Current development and production uses the following versions:

- [ElasticSearch] 5.6 (configured port: 9200)
- [Apache Fuseki] 3.5.0 (configured port: 3030)

All of these dependencies are available on the major platforms via package-managers or direct download: Linux, MacOS and Windows. Please follow the respective installation instructions to get them up and running. RAPID will complain at startup when these are not available.

The easiest setup for development is using the supplied `docker-compose.yml`. Run the following command to start all the external dependencies. 

    `$ docker-compose up`

It also includes [Kibana] also runs on `http://localhost:5601` to provide a development console for ElasticSearch.

### Please help out

This project is still under development. Feedback and suggestions are very
welcome and we encourage you to use the [Issues
list](http://github.com/delving/rapid/issues) on Github to provide that
feedback.

Feel free to fork this repo and to commit your additions. For a list of all
contributors, please see the [AUTHORS](AUTHORS) file.

### Contributing

Fork the repository.  Then, run:

    git clone --recursive git@github.com:<username>/rapid.git
    cd rapid 
    git branch master origin/master
    git flow init -d
    git flow feature start <your feature>

Then, do work and commit your changes.  

    git flow feature publish <your feature>

When done, open a pull request to your feature branch.


## License

Copyright (c) 2017-present Delving B.V.

Licensed under [Apache 2.0](./License)

[Golang Installation]: https://golang.org/doc/install
[dep]: https://golang.github.io/dep/ 
[Kibana]: https://www.elastic.co/products/kibana
[golang]: https://golang.org/
[Apache Lucene]: https://lucene.apache.org/ 
[Apache Fuseki]: https://jena.apache.org/documentation/fuseki2/index.html
[ElasticSearch]: https://www.elastic.co/guide/en/elasticsearch/reference/6.2/getting-started.html
[SPARQL 1.1]: https://www.w3.org/TR/sparql11-query/
[SPARQL Update 1.1]: https://www.w3.org/TR/sparql11-update/
