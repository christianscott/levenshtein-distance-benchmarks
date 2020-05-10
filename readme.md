# Levenshtein edit distance benchmarks

Several implementations of a DP levenshtein edit distance alogorithm in different programming languages.

- `yarn bench`: run the benchmarks
- `yarn check`: test the output of each of the implementations against expected.txt

Benchmarks at the time of writing:

```bash
$ yarn bench
yarn run v1.22.0
$ node run.js bench
go: 1.477370
javascript: 3.222
rust: 2.790513908
✨  Done in 8.45s.
```

Python and c are omitted by renaming bench.sh to _bench.sh. Python is omitted due to taking ~200 seconds (100x!!!), and C is omitted because I can't write C 🤠.
