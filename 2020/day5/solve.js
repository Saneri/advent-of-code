const assert = require("assert");

const utils = require("../utils");

const rowBlock = 7;
const columnBlock = 3;

function solver(input) {
  const seats = [];
  input.forEach((binaryPartition) => {
    const row = binaryPartition.slice(0, rowBlock);
    const column = binaryPartition.slice(rowBlock, rowBlock + columnBlock);
    const seatId = decode(row, "F", "B") * 8 + decode(column, "L", "R");
    seats.push(seatId);
  });
  seats.sort((a, b) => a - b);
  let lastId = seats[0];
  seats.slice(1).some((id) => {
    lastId++;
    return id != lastId;
  });
  return lastId;
}

function decode(str, lowerKey, upperKey) {
  const nOfLetters = str.length;
  let upper = Math.pow(2, nOfLetters);
  let lower = 0;
  for (let char of str) {
    const average = (lower + upper) / 2;
    if (char === lowerKey) {
      upper = average;
    } else if (char === upperKey) {
      lower = average;
    } else {
      throw "invalid character in str";
    }
  }
  if (lower != upper - 1) {
    throw "Calculation error";
  }
  return lower;
}

assert(decode("FBFBBFF", "F", "B") === 44);
assert(decode("RLR", "L", "R") === 5);
//assert(solver(["FBFBBFFRLR"]) === 357);

async function main() {
  const input = await utils.readInput("day5/input.txt", utils.types.STRING);
  console.log(solver(input));
}

main();
