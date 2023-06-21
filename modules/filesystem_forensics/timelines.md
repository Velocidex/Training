<!-- .slide: class="title" -->

# Timelines

---

<!-- .slide: class="content " -->

## What is a timeline?

* It is a way to visualize time based rows from multiple sources.
* The main concepts:
    * `Timeline`: Just a series of rows keyed on a time column. The
      rows can be anything at all, as long as a single column is
      specified as the time column and it is sorted by time order.
    * `Super Timeline`: A grouping of several timelines viewed
      together on the same timeline.

---

<!-- .slide: class="content small-font" -->

## Timeline workflow

* Timelines are created from post processed results from the notebook:
    1. Collect a set of artifacts with relevant information:
       * e.g. MFT entries, Prefetch, Event logs etc.
    2. Create a `Supertimeline` to hold all the timelines together.
    3. Reduce the data from each artifact source by manipulating the
       VQL query:
       * Reduce the number of rows by limiting only interesting rows.
       * Reduce the columns by adding only important columns.
    4. Add the table to the `Super Timeline` by selecting the time
       column.

---

<!-- .slide: class="content " -->

## Example:

TODO
