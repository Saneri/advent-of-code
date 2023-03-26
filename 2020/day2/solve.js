const { strict } = require("assert");
const assert = require("assert");

const utils = require("../utils");

function solver(input) {
  let correctPassAmount = 0;
  input.forEach((element) => {
    const parts = [...element.matchAll("(\\d+)-(\\d+) (\\w): (\\w+)")][0];
    const lowIndex = parseInt(parts[1]) - 1;
    const highIndex = parseInt(parts[2]) - 1;
    const alphabet = parts[3];
    const pass = parts[4];

    if ((pass[lowIndex] == alphabet) != (pass[highIndex] == alphabet)) {
      correctPassAmount += 1;
    }
  });
  return correctPassAmount;
}

assert(solver(["1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"]) === 1);

async function main() {
  const input = await utils.readInput("day2/input.txt", utils.types.STRING);
  console.log(solver(input));
}

main();
