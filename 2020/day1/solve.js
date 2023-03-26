const assert = require("assert");
const utils = require("../utils");

const wantedNumber = 2020;

function solver(input) {
  for (let i = 0; i < input.length; i++) {
    for (let j = 0; j < input.length; j++) {
      for (let k = 0; k < input.length; k++) {
        if ((i != j) != k && input[i] + input[j] + input[k] === wantedNumber) {
          return input[i] * input[j] * input[k];
        }
      }
    }
  }
}

assert(solver([1721, 979, 366, 299, 675, 1456]) == 241861950);

async function main() {
  const input = await utils.readInput("day1/input.txt", utils.types.INT);
  console.log(solver(input));
}

main();
