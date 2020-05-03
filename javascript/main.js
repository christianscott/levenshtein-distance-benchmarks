const fs = require('fs');

const lines = fs
  .readFileSync('../sample.txt')
  .toString()
  .split('\n');

const start = Date.now();
benchmark(lines);
const end = Date.now() - start;

console.log(`${end / 1000}`);

function benchmark(lines) {
  for (var i = 0; i < 1e4; i++) {
    let lastValue = '';
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

  const cache = new Array(target.length + 1);
  for (let i = 0; i < target.length + 1; i++) {
    cache[i] = i;
  }

  for (let i = 0; i < source.length; i++) {
    let nextDistance = i + 1;
    const sourceChar = source.charAt(i);

    for (let j = 0; j < target.length; j++) {
      let currentDistance = nextDistance;

      const areCharCodesEqual = sourceChar === target.charAt(j);

      const distanceIfSubstitute = cache[j] + (areCharCodesEqual ? 0 : 1);
      const distanceIfInsert = currentDistance + 1;
      const distanceIfDelete = cache[j + 1] + 1;

      nextDistance = Math.min(distanceIfSubstitute, distanceIfInsert, distanceIfDelete);

      cache[j] = currentDistance;
    }

    cache[target.length] = nextDistance;
  }

  return cache[target.length];
}
