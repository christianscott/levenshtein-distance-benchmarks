const execa = require("execa");
const fg = require("fast-glob");
const fs = require("fs");
const path = require("path");

function main() {
  const args = process.argv.slice(process.argv.indexOf(__filename) + 1);
  const [command] = args;
  switch (command) {
    case "bench":
      return bench();
    case "check":
      return check();
    default:
      throw new Error(`unknown command ${command}`);
  }
}

async function bench() {
  for await (const { dir, stdout } of run()) {
    console.log(`${dir}: ${stdout}`);
  }
}

async function check() {
  const expected = fs
    .readFileSync("./expected.txt", { encoding: "utf8" })
    .trim();
  let anyFailed = false;
  for await (const { dir, stderr } of run()) {
    if (stderr.trim() === expected) {
      console.log(`${dir}: ok`);
    } else {
      anyFailed = true;
      console.log(
        `${dir}: wrong output\n  expected: ${expected}\n  got: ${stderr.trim()}`
      );
    }
  }

  if (anyFailed) {
    console.error("one or more programs produced the wrong output");
    process.exit(1);
  }
}

async function* run() {
  const benchFiles = fg.sync("./**/bench.sh");
  for (const benchFile of benchFiles) {
    const dir = path.dirname(benchFile);
    const { stdout, stderr } = await execa("bash", ["bench.sh"], {
      cwd: dir,
    });
    yield { dir, stdout, stderr };
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
