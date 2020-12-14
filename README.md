# stackexchange-xml-to-csv

CLI tool that allows you to convert [Stack Exchange data dumps](https://archive.org/download/stackexchange) from `XML` to `CSV` format, which is more suitable for importing to the different databases.

Table of contents
=================
* [Getting started](#getting-started)
    * [Download database dump](#download-database-dump)
    * [Extract](#extract)
    * [Building of stackexchange-xml-to-csv](#building-of-stackexchange-xml-to-csv)
    * [XML to CSV converting](#xml-to-csv-converting)
* [RDBMS schema examples](#rdbms-schema-examples)
* [License](#license)


Getting started
===============
Before, ensure that you have:
* Working [Go environment](https://golang.org/doc/install) with go version >= 1.14. Execute in the console `go version` command. It should display the current version of the compiler.
* Archiver that can extract `.7z` files. Possible candidate is [7z](https://www.7-zip.org/).


Download database dump
======================

Choose and download the [database dump](https://archive.org/download/stackexchange) that you are going to convert.

**Important: Stackoverflow dump stored in 8 separated 7z archives:**

* [stackoverflow.com-Badges.7z](https://archive.org/download/stackexchange/stackoverflow.com-Badges.7z) ( ~ **70M** compressed / **4G** uncompressed / **37M** rows )
* [stackoverflow.com-Comments.7z](https://archive.org/download/stackexchange/stackoverflow.com-Comments.7z) ( ~ **4.5G** compressed / **22G** uncompressed / **76M** rows )
* [stackoverflow.com-PostHistory.7z](https://archive.org/download/stackexchange/stackoverflow.com-PostHistory.7z) ( ~ **28.0G** compressed / **138G** uncompressed / **133M** rows)
* [stackoverflow.com-PostLinks.7z](https://archive.org/download/stackexchange/stackoverflow.com-PostLinks.7z) ( ~ **100M** compressed / **800M** uncompressed / **7M** rows)
* [stackoverflow.com-Posts.7z](https://archive.org/download/stackexchange/stackoverflow.com-Posts.7z) ( ~ **16G** compressed / **80G** uncompressed / **50M** rows)
* [stackoverflow.com-Tags.7z](https://archive.org/download/stackexchange/stackoverflow.com-Tags.7z) ( ~ **900K** compressed / **5.0M** uncompressed / **60K** rows)
* [stackoverflow.com-Users.7z](https://archive.org/download/stackexchange/stackoverflow.com-Users.7z) ( ~ **650M** compressed / **4.0G** uncompressed / **13M** rows)
* [stackoverflow.com-Votes.7z](https://archive.org/download/stackexchange/stackoverflow.com-Votes.7z) ( ~ **1.0G** compressed / **20G** uncompressed / **200M** rows)

Extract
=======

Extract archive(s) content file(s) to the directory from where you will convert files using `7z` or another archiver.

Example with with [academia.stackexchange.com.7z](https://archive.org/download/stackexchange/academia.stackexchange.com.7z) dump:
```shell
$ mkdir xml csv
$ 7z e academia.stackexchange.com.7z -oxml
$ ls xml/
Badges.xml  Comments.xml  PostHistory.xml  PostLinks.xml  Posts.xml  Tags.xml  Users.xml  Votes.xml
```

Building of stackexchange-xml-to-csv
====================================

Clone & build `stackexchange-xml-to-csv` converter:

```shell
$ git clone https://github.com/SkobelevIgor/stackexchange-xml-to-csv
$ cd stackexchange-xml-to-csv/
$ go build
```

XML to CSV converting
=====================

Now you have `stackexchange-xml-to-csv` executable file. Let’s convert XML files:
```
./stackexchange-xml-to-csv -—source-path=../xml --store-to-dir=../csv
```
### List of possible flags:

* `source-path` (**Required**) Absolute or relative path to the directory with an XML file(s) or to the separate XML file.
* `store-to-dir` (**Optional**) Absolute or relative path to the directory where to store result CSV files.
* `skip-html-decoding` (**Optional**) Some of the files (e.g., Posts.xml) contain escaped HTML. By default, the converter will decode them. To disable this behavior, use this flag.


RDBMS schema examples
=====================
Here you can find examples of the schema for the different databases:

* [PostgreSQL](./example/postgresql_ddl.sql)

License
=======

[MIT License](./LICENSE)
