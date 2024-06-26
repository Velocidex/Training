<!-- content optional -->

## The Linux Journal Logs

* Recent Linux systems use `systemd`
* Logs are structured and incorporate indexes for fast searching.

* You can view the logs with `journalctl`

```
journalctl --file /run/log/journal/*/*.journal
```

---

<!-- content small-font optional -->

## Exercise: Parsing Journal Logs: Execve

* You can use `execve()` to run an external binary
* Use `journalctl -o json --file X.journal` to read the journal file
* Use `parse_jsonl()` to parse the output into structured data.

---

<!-- content small-font optional -->

## Exercise: Parsing Journal Logs: Natively

* It is useful to be able to parse the log file directly.

* Write a VQL parser for the Journal log file.
* The format is documented
  https://www.freedesktop.org/wiki/Software/systemd/journal-files/

* Get a sample file from https://github.com/Velocidex/velociraptor/raw/master/artifacts/testdata/files/system.journal


---

<!-- content small-font optional -->

## Exercise: Parsing Journal Logs

* Structure is:
   1. Header: Provides metadata
   2. Object Header: Provides Type and Size
      * Trick: Size is 8 byte aligned.
   3. Different Objects follow depending on Type.
   4. We only care about:
      * DATA_OBJECT: Contains one item
      * ENTRY_OBJECT: Contains one log line - refers to multiple DATA_OBJECT.

---

<!-- content small-font optional -->

## Exercise: Parsing Journal Logs

* Parse header

<div class="solution solution-closed">

```
    LET JournalProfile = '''[
    ["Header", "x=>x.header_size", [
      ["Signature", 0, "String", {
          "length": 8,
      }],
      ["header_size", 88, "uint64"],
      ["arena_size", 96, "uint64"],
      ["n_objects", 144, uint64],
      ["n_entries", 152, uint64],
      ["Objects", "x=>x.header_size", "Array", {
          "type": "ObjectHeader",
          "count": "x=>x.n_objects",
          "max_count": 100000
      }]
    ]],
```

</div>

* Parse Object Headers

<div class="solution solution-closed">

```
    ["ObjectHeader", "x=>x.size", [
     ["Offset", 0, "Value", {
        "value": "x=>x.StartOf",
     }],
     ["type", 0, "Enumeration",{
         "type": "uint8",
         "choices": {
          "0": OBJECT_UNUSED,
          "1": OBJECT_DATA,
          "2": OBJECT_FIELD,
          "3": OBJECT_ENTRY,
          "4": OBJECT_DATA_HASH_TABLE,
          "5": OBJECT_FIELD_HASH_TABLE,
          "6": OBJECT_ENTRY_ARRAY,
          "7": OBJECT_TAG,
         }
     }],
     ["flags", 1, "uint8"],
     ["__real_size", 8, "uint64"],
     ["__round_size", 8, "Value", {
         "value": "x=>int(int=x.__real_size / 8) * 8",
     }],
     ["size", 0, "Value", {
         "value": "x=>if(condition=x.__real_size = x.__round_size, then=x.__round_size, else=x.__round_size + 8)",
     }],
     ["payload", 16, Union, {
         "selector": "x=>x.type",
         "choices": {
             "OBJECT_DATA": DataObject,
             "OBJECT_ENTRY": EntryObject,
         }
     }]
    ]],
```
</div>

* [Full solution](https://github.com/Velocidex/velociraptor/blob/master/artifacts/definitions/Linux/Forensics/Journal.yaml)
