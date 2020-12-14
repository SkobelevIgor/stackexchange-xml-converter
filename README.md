# stackexchange-xml-to-csv

**stackexchange-xml-to-csv** is a CLI tool that allows you to convert [Stack Exchange data dumps](https://archive.org/download/stackexchange) from `XML` to `CSV` format, which is more suitable for importing to the different databases.

Table of contents.
=================
* [Getting started](#get_start)
    * [Download database dump](#download-dump)
    * [Extract archive(s)](#extract)
    * [stackexchange-xml-to-csv building](#stackexchange-xml-to-csv-build)
    * [XML to CSV Convertation](#xml-to-csv)
* [RDBMS schema examples](#examples)
    * [PostgreSQL](#pg)
    * [MySQL](#mysql)
* [License](#license)


Getting started.
================
Before, ensure that you have a working [Go environment](https://golang.org/doc/install) with go version >= 1.14. Execute in the console `go version` command. It should display the current version of the compiler.


1. Download database dump.
==========================

Choose and download the [database dump](https://archive.org/download/stackexchange) that you are going to convert.

**Important: Stackoverflow dump stored in 8 separated 7z archives:**

* [stackoverflow.com-Badges.7z](https://archive.org/download/stackexchange/stackoverflow.com-Badges.7z)
* [stackoverflow.com-Comments.7z](https://archive.org/download/stackexchange/stackoverflow.com-Comments.7z)
* [stackoverflow.com-PostHistory.7z](https://archive.org/download/stackexchange/stackoverflow.com-PostHistory.7z)
* [stackoverflow.com-PostLinks.7z](https://archive.org/download/stackexchange/stackoverflow.com-PostLinks.7z)
* [stackoverflow.com-Posts.7z](https://archive.org/download/stackexchange/stackoverflow.com-Posts.7z)
* [stackoverflow.com-Tags.7z](https://archive.org/download/stackexchange/stackoverflow.com-Tags.7z)
* [stackoverflow.com-Users.7z](https://archive.org/download/stackexchange/stackoverflow.com-Users.7z)
* [stackoverflow.com-Votes.7z](https://archive.org/download/stackexchange/stackoverflow.com-Votes.7z)

## 2. Extract archive(s).
========================

Extract archive(s) content file(s) to the directory from where you will convert files using [7z](https://www.7-zip.org/) or another archiver.

Example with with [academia.stackexchange.com.7z](https://archive.org/download/stackexchange/academia.stackexchange.com.7z) dump:
```shell
$ mkdir xml csv
$ 7z e academia.stackexchange.com.7z -oxml
$ ls xml/
Badges.xml  Comments.xml  PostHistory.xml  PostLinks.xml  Posts.xml  Tags.xml  Users.xml  Votes.xml
```

## 3. stackexchange-xml-to-csv building.
===========================================

Clone & build `stackexchange-xml-to-csv` converter:

```shell
$ git clone https://github.com/SkobelevIgor/stackexchange-xml-to-csv
$ cd stackexchange-xml-to-csv/
$ go build
```

## 4. XML to CSV Convertation.
=============================

Now you have `stackexchange-xml-to-csv` executable file. Let’s convert XML files:
```
./stackexchange-xml-to-csv -—source-path=../xml --store-to-dir=../csv
```
### List of possible flags:

* `source-path` (**Required**) Absolute or relative path to the directory with an XML file(s) or to the separate XML file.
* `store-to-dir` (**Optional**) Absolute or relative path to the directory where to store result CSV files.
* `skip-html-decoding` (**Optional**) Some of the files (e.g., Posts.xml) contain escaped HTML. By default, the converter will decode them. To disable this behavior, use this flag.


Schema examples.
================
Here you can find examples of the schema for different databases:
    * [PostgreSQL](example/postgres_ddl.sql)


License
=======

[MIT License](LICENSE)
