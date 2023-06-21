<!-- .slide: class="title" -->

# Parsing

## Processing and analysing evidence on the endpoint

---

<!-- .slide: class="content small-font" -->

## Parsing evidence on the endpoint

* By analyzing files directly on the endpoint we can extract relevant
  data immediately.

* Velociraptor supports sophisticated parsing strategies that allow
  VQL artifacts to extract maximum details directly on the endpoint.

   * Built in parsers (`parse_ntfs`, `parse_xml`, `parse_json`)
   * Text based parsers (`parse_string_with_regex`, `split`)
   * Binary parser

* By eliminating the need for post processing we can scale analysis
  across larger number of endpoints


---

<!-- .slide: class="content small-font" -->

## Built in parsers - SQLite

TODO

---

<!-- .slide: class="content small-font" -->

## Parsing with Regular Expressions

TODO

---

<!-- .slide: class="content " -->

## Complex RegEx parsing

* Sometimes log files are less structures and a regex based approach
  is not reliable enough.

* In this case think about how to split the data in a reliable way and
  apply regular expressions multiple times.

* Divide and Concur

---

<!-- .slide: class="content " -->

## Exercise: MPLogs

* [Mind the MPLog: Leveraging Microsoft Protection Logging for Forensic Investigations](https://www.crowdstrike.com/blog/how-to-use-microsoft-protection-logging-for-forensic-investigations/)

* MPLog files are found in `C:\ProgramData\Microsoft\Windows Defender\Support`
* Events described in [This Reference](https://learn.microsoft.com/en-us/microsoft-365/security/defender-endpoint/troubleshoot-performance-issues?view=o365-worldwide)

Write a VQL Parser to parse these logs.

---

<!-- .slide: class="content small-font" -->

## Steps for solution

1. Locate data from disk and split into separate log lines (records).
2. Find a strategy to parse each record:
    * Will one pass Regex work?
    * What is the structure of the line

3. You may use the following plugins:
    * `split`, `utf16`, `parse_lines`, `read_file`, `to_dict`, `foreach`
    * `parse_string_with_regex`, `glob`
    * Dict addition

---

<!-- .slide: class="content small-font" -->

## Possible solution

* Not really perfect because log is not very consistent.

```sql
LET LogGlob = '''C:\ProgramData\Microsoft\Windows Defender\Support\MPLog*.log'''

LET AllLines = SELECT * FROM foreach(row={
  SELECT utf16(string=read_file(filename=OSPath, length=10000000)) AS Data,
        OSPath
  FROM glob(globs=LogGlob)
}, query={
  SELECT Line, OSPath
  FROM parse_lines(filename=Data, accessor="data")
})

LET ParseData(Data) = to_dict(item={
  SELECT split(sep_string=":", string=_value)[0] AS _key,
         split(sep_string=":", string=_value)[1] AS _value
  FROM foreach(row=split(sep=", ", string=Data))
})

LET Lines = SELECT OSPath, Line,
   parse_string_with_regex(string=Line,
       regex="^(?P<Timestamp>[^ ]+) (?P<Data>.+)") AS P
FROM X
WHERE P.Timestamp

SELECT * FROM foreach(row={
   SELECT dict(Timestamp=P.Timestamp, _Line=Line, _OSPath=OSPath) +
          ParseData(Data=P.Data) AS Data
   FROM Lines
}, column="Data")
```


---

<!-- .slide: class="title" -->

# The Binary Parser

---

<!-- .slide: class="content " -->

## Parsing binary data

* A lot of data we want to parse is binary only
* Having a powerful binary parser built into VQL allows the VQL query
  to parse many more things!
* [VQL Binary parser](https://github.com/Velocidex/vtypes) is declerative.

---

<!-- .slide: class="content small-font" -->

## What is binary data?

* Serialized representation of abstract data structures
* Declare the layout of the data and let the parser recover the data
  from the binary stream.

TODO
