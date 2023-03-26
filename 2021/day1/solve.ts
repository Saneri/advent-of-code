import { readFileSync } from 'fs';
import { join } from 'path';

const solve1a = (data: Array<number>): number => {
  const increaseAmount = data.reduce((previous, _, index, arr): number => {
    if (arr[index] > arr[index - 1]) {
      return previous + 1;
    }

    return previous;
  }, 0);
  return increaseAmount;
};

const solve1b = (data: Array<number>): number => {
  const slidingWindowSum = (arr: Array<number>, index: number): number => {
    return arr[index - 1] + arr[index] + arr[index + 1];
  };

  const increaseAmount = data.reduce((previous, _, index, arr): number => {
    if (
      arr[index - 1] &&
      arr[index + 1] &&
      slidingWindowSum(arr, index) > slidingWindowSum(arr, index - 1)
    ) {
      return previous + 1;
    }

    return previous;
  }, 0);
  return increaseAmount;
};

const data = readFileSync(join(__dirname, 'input.txt'), 'utf8')
  .split('\n')
  .map(Number);

console.log(`1a: ${solve1a(data)}`);
console.log(`1b: ${solve1b(data)}`);
