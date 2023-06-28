
<!-- .slide: class="content small-font" -->

## Searching for files - glob()

* One of the most common operations in DFIR is searching for files
  efficiently.
* Velociraptor has the `glob()` plugin to search for files using a
  glob expression.
* Glob expressions use wildcards to search the filesystem for matches.
    * Paths are separated by / or \ into components
    * A `*` is a wildcard match (e.g. `*.exe` matches all files ending
      with .exe)
    * Alternatives are expressed as comma separated strings in `{}`
      e.g. `*.{exe,dll,sys}`
    * A `**` denotes recursive search. `e.g. C:\Users\**\*.exe`

---

<!-- .slide: class="content small-font" -->

## Exercise: Search for exe

* Search user’s home directory for binaries.

```sql
SELECT * FROM glob(globs='C:\\Users\\**\\*.exe')
```

Note the need to escape `\` in strings. You can use `/` instead and
specify multiple globs to search all at the same time:

```sql
SELECT * FROM glob(globs=['C:/Users/**/*.exe',
                          'C:/Users/**/*.dll'])
```

---

<!-- .slide: class="content" -->
## Filesystem accessors

* Glob is a very useful concept to search hierarchical trees
* Velociraptor supports direct access to many different such trees via accessors (essentially FS drivers):

    * `file` - uses OS APIs to access files.
    * `ntfs` - uses raw NTFS parsing to access low level files
    * `reg` - uses OS APIs to access the windows registry
    * `raw_reg` - search in a registry hive
    * `zip` - Search inside zip files

---

<!-- .slide: class="content small-font" -->
## The registry accessor

* Uses the OS API to access the registry
* Top level consists of the major hives (`HKEY_USERS` etc)
* Values appear as files, Keys appear as directories
* Default value is named “@”
* Value content is included inside the Data attribute
* Can escape components with / using quotes

`HKLM\Microsoft\Windows\"http://www.microsoft.com/"`


---

<!-- .slide: class="content small-font" -->

## The registry accessor

* The `OSPath` column includes the key (as directory) and the value (as a
  filename) in the path.

* The Registry accessor also includes value contents if they are small
  enough in the `Data` column.


---

<!-- .slide: class="content small-font" -->

## Exercise - RunOnce artifact

* Write an artifact which hashes every binary mentioned in Run/RunOnce
  keys.

* “Run and RunOnce registry keys cause programs to run each time that
  a user logs on.”

*  https://learn.microsoft.com/en-us/windows/win32/setupapi/run-and-runonce-registry-keys

---

<!-- .slide: class="content small-font" -->

## Exercise: Hash all files provided in the globs

Create an artifact that hashes files found by user provided globs.

---

<!-- .slide: class="title" -->

# Searching data

---

<!-- .slide: class="content small-font" -->

## Searching data

* A powerful DFIR technique is searching bulk data for patterns
* Searching for CC data in process memory
* Searching for URLs in process memory
* Searching binaries for malware signatures
* Searching registry for patterns

Bulk searching helps to identify evidence without needing to parse
file formats

---

<!-- .slide: class="content small-font" -->

## YARA - The swiss army knife

* YARA is a powerful keyword scanner
* Uses rules designed to identify binary patterns in bulk data
* YARA is optimized to scan for many rules simultaneously.
* Velociraptor supports YARA scanning of bulk data (via accessors) and memory.

`yara()` and `proc_yara()`

---

<!-- .slide: class="content small-font" -->

## YARA rules

Yara rules has a special domain specific language

```yara
rule X {
   strings:
       $a = “hello” nocase
       $b = “Goodbye” wide
       $c = /[a-z]{5,10}[0-9]/i

   condition:
       $a and ($b or $c)
}
```

---

<!-- .slide: class="content small-font" -->

## Exercise: drive by download

* You suspect a user was compromised by a drive by download (i.e. they
  clicked and downloaded malware delivered by mail, ads etc).
* You think the user used the Edge browser but you have no idea of the
  internal structure of the browser cache/history etc.
* Write an artifact to extract potential URLs from the Edge browser
  directory (also where is it?)

---

<!-- .slide: class="content small-font" -->

## Step 1: Figure out where to look


Looks like somewhere in `C:\Users\<name>\AppData\Local\Microsoft\Edge\**`

---

<!-- .slide: class="content small-font" -->

## Step 2: Recover URLs

* We don't exactly understand how Edge stores data but we know roughly
  what a URL is supposed to look like!

* Yara is our sledgehammer !

```
rule URL {
  strings: $a = /https?:\\/\\/[a-z0-9\\/+&#:\\?.-]+/i
  condition: any of them
}
```

---

<!-- .slide: class="content small-font" -->


## Step 3: Let’s do this!


---

<!-- .slide: class="content small-font" -->

## YARA best practice

* You can get yara rules from many sources (threat intel, blog posts etc)
* YARA is really a first level triage tool:
    * Depending on signature many false positives expected
    * Some signatures are extremely specific so make a great signal
    * Try to collect additional context around the hits to eliminate
      false positives.
    * Yara scanning is relatively expensive! consider more targeted
      glob expressions and client side throttling since usually YARA
      scanning is not time critical.

---

<!-- .slide: class="content small-font" -->

## Uploading files

* Velociraptor can collect file data.
    * Over the network
    * Locally to a collection zip file.
    * Driven by VQL

The `upload()` VQL function copies a file using an accessor to the
relevant container

---

<!-- .slide: class="content small-font" -->

## Exercise

Collect all executables in users’ home directory


Write your own VQL by combining `glob()` and `upload()`
