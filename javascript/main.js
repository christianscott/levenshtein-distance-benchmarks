var fs = require("fs");

var lines = fs.readFileSync("../sample.txt").toString().split("\n");

var start = Date.now();
benchmark(lines);
var end = Date.now() - start;

console.log(`${end / 1000}`);

function benchmark(lines) {
  for (var i = 0; i < 1e4; i++) {
    var lastValue = "";
    for (var j = 0; j < lines.length; j++) {
      levenshteinDistance(lastValue, lines[j]);
      lastValue = lines[j];
    }
  }
}

function levenshteinDistance(source, target) {
  if (source.length === 0) {
    return target.length;
  }

  if (target.length === 0) {
    return source.length;
  }

  var cache = new Uint32Array(target.length + 1);
  for (var i = 0; i < target.length + 1; i++) {
    cache[i] = i;
  }

  var sourceLen = source.length;
  for (var i = 0; i < sourceLen; i++) {
    var nextDistance = i + 1;
    var sourceChar = source.charAt(i);

    var targetLen = target.length;
    for (var j = 0; j < targetLen; j++) {
      var currentDistance = nextDistance;

      var areCharCodesEqual = sourceChar === target.charAt(j);

      var distanceIfSubstitute = (cache[j] + (areCharCodesEqual ? 0 : 1)) | 0;
      var distanceIfInsert = (currentDistance + 1) | 0;
      var distanceIfDelete = (cache[j + 1] + 1) | 0;

      nextDistance = Math.min(
        distanceIfSubstitute,
        distanceIfInsert,
        distanceIfDelete
      );

      cache[j] = currentDistance;
    }

    cache[target.length] = nextDistance;
  }

  return cache[target.length];
}
