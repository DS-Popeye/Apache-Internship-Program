# Apache AGE Golang Driver

[Apache AGE](https://age.apache.org/) is a PostgreSQL extension that provides graph database functionality. The goal of the Apache AGE project is to create single storage that can handle both relational and graph model data so that users can use standard ANSI SQL along with openCypher, the Graph query language. This repository hosts the development of the Python driver for this Apache extension (currently in Incubator status). Thanks for checking it out.

A graph consists of a set of vertices (also called nodes) and edges, where each individual vertex and edge possesses a map of properties. A vertex is the basic object of a graph, that can exist independently of everything else in the graph. An edge creates a directed connection between two vertices. A graph database is simply composed of vertices and edges. This type of database is useful when the meaning is in the relationships between the data. Relational databases can easily handle direct relationships, but indirect relationships are more difficult to deal with in relational databases. A graph database stores relationship information as a first-class entity. Apache AGE gives you the best of both worlds, simultaneously.

Apache AGE is:

- **Powerful** -- AGE adds graph database support to the already popular PostgreSQL database: PostgreSQL is used by organizations including Apple, Spotify, and NASA.
- **Flexible** -- AGE allows you to perform openCypher queries, which make complex queries much easier to write.
- **Intelligent** -- AGE allows you to perform graph queries that are the basis for many next level web services such as fraud & intrustion detection, master data management, product recommendations, identity and relationship management, experience personalization, knowledge management and more.

### Features

- Cypher query support for database/sql PostreSQL driver (enables cypher queries directly)
- Deserialize AGE result (AGType) to Vertex, Edge, Path

### Go get & gomod

- over Go 1.16
- This module runs on golang standard api [database/sql](https://golang.org/pkg/database/sql/) and [antlr4-go](https://github.com/antlr/antlr4/tree/master/runtime/Go/antlr)

```(shell)
go get github.com/DS-Popeye/Apache-Internship-Program/Question-3/Go

```

#### gomod

```(go)
require  github.com/DS-Popeye/Apache-Internship-Program/Question-3/Go {version}

```

Check [latest version](https://github.com/DS-Popeye/Apache-Internship-Program/Question-3/Go/releases)

### For more information about [Apache AGE](https://age.apache.org/)

- Apache Incubator Age : https://age.apache.org/
- Github : https://github.com/apache/incubator-age
- Document : https://age.incubator.apache.org/docs/

### Check AGE loaded on your PostgreSQL

Connect to your containerized Postgres instance and then run the following commands:

```(sql)
# psql
CREATE EXTENSION age;
LOAD 'age';
SET search_path = ag_catalog, "$user", public;
```

### Test

Check out and rewrite DSN in incubator-age/drivers/golang/age_test.go

```
cd apache-age-go
go test -v .

```

### Samples

**Usage 1: using database/sql API and Cypher execution function 'ExecCypher'**

> - Sample : [samples/sql_api_sample.go](samples/sql_api_sample.go)

**Usage 2: using Age Wrapper**

> - Sample : [samples/age_wrapper_sample.go](samples/age_wrapper_sample.go)

**Run Samples**

> Sample : [samples/main.go](samples/main.go)

### License

Apache-2.0 License
