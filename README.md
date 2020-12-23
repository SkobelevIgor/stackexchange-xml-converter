# stackexchange-xml-converter

CLI tool that allows you to convert [Stack Exchange data dumps](https://archive.org/download/stackexchange) from `XML` to `CSV` or `JSON` formats, which is more suitable for importing to the different databases.

Table of contents
=================
* [RDBMS schema examples](#rdbms-schema-examples)
    * [PostgreSQL]()
    * [MySQL]()
* [Getting started](#getting-started)
    * [Download database dump](#download-database-dump)
    * [Extract](#extract)
    * [Build the stackexchange-xml-converter](#build-the-stackexchange-xml-converter)
    * [XML to CSV converting](#xml-to-csv-converting)
* [License](#license)

RDBMS schema examples
=====================
Here you can find the examples of the schema for the different databases:

* [PostgreSQL](./schema_example/postgresql_ddl.sql)
* [MySQL](./schema_example/mysql_ddl.sql)

Getting started
===============
Before, ensure that you have:
* Working [Go environment](https://golang.org/doc/install) with go version >= 1.14. Execute in the console `go version` command. It should display the current version of the compiler.
* Archiver that can extract `.7z` files. Possible candidate is [7z](https://www.7-zip.org/).


### Download database dump

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

### Extract

Extract archive(s) content file(s) to the directory from where you will convert XML files.

Example with [academia.stackexchange.com.7z](https://archive.org/download/stackexchange/academia.stackexchange.com.7z) dump:
```shell
$ mkdir xml csv
$ 7z e academia.stackexchange.com.7z -oxml
$ ls xml/
Badges.xml  Comments.xml  PostHistory.xml  PostLinks.xml  Posts.xml  Tags.xml  Users.xml  Votes.xml
```

### Build the stackexchange-xml-converter

Clone & build `
stackexchange-xml-converter` converter:

```shell
$ git clone https://github.com/SkobelevIgor/stackexchange-xml-converter
$ cd stackexchange-xml-converter/
$ go build
```

### XML to CSV/JSON converting


Now you have the `stackexchange-xml-converter` executable file. Let’s convert XML files to the CSV format:
```
./stackexchange-xml-converter -result-format=csv -source-path=../xml -store-to-dir=../csv
```
#### List of possible flags:

* `result-format` (**Required**) Result format (csv or json)
* `source-path` (**Required**) Absolute or relative path to the directory with an XML file(s) or to the separate XML file.
* `store-to-dir` (**Optional**) Absolute or relative path to the directory where to store result CSV files.
* `skip-html-decoding` (**Optional**) Some of the files (e.g., Posts.xml) contain escaped HTML. By default, the converter will decode them. To disable this behavior, use this flag.

License
=======

[MIT License](./LICENSE)
