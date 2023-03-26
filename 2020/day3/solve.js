const assert = require("assert");
const utils = require("../utils");

const slopes = [
  [1, 1],
  [3, 1],
  [5, 1],
  [7, 1],
  [1, 2],
];

function calculateTrees(input, horizontalSlope, verticalSlope) {
  let verticalPosition = 0;
  let horizontalPosition = 0;
  let treeCount = 0;
  const height = input.length - 1;
  const width = input[0].length - 1;

  while (horizontalPosition < height) {
    horizontalPosition += horizontalSlope;
    verticalPosition += verticalSlope;

    if (verticalPosition > width) {
      verticalPosition = (verticalPosition % width) - 1;
    }

    if (input[horizontalPosition][verticalPosition] == "#") {
      treeCount++;
    }
  }

  return treeCount;
}

function solver(input, slopes) {
  treeProduct = 1;
  slopes.forEach((slope) => {
    treeProduct *= calculateTrees(input, slope[1], slope[0]);
  });
  return treeProduct;
}

assert(
  calculateTrees(
    [
      "..##.......",
      "#...#...#..",
      ".#....#..#.",
      "..#.#...#.#",
      ".#...##..#.",
      "..#.##.....",
      ".#.#.#....#",
      ".#........#",
      "#.##...#...",
      "#...##....#",
      ".#..#...#.#",
    ],
    1,
    3
  ) === 7
);

assert(
  solver(
    [
      "..##.......",
      "#...#...#..",
      ".#....#..#.",
      "..#.#...#.#",
      ".#...##..#.",
      "..#.##.....",
      ".#.#.#....#",
      ".#........#",
      "#.##...#...",
      "#...##....#",
      ".#..#...#.#",
    ],
    slopes
  ) === 336
);

async function main() {
  const input = await utils.readInput("day3/input.txt", utils.types.STRING);
  console.log(solver(input, slopes));
}

main();
