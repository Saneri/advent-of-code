import { readFileSync } from 'fs';
import { join } from 'path';

const sumBinary = (
  gammaBits: number[],
  index: number,
  bitArray: number[][]
): number[] => {
  const updatedBits = gammaBits.map((x: number, i: number): number => {
    return x + bitArray[index][i];
  });
  if (index === bitArray.length - 1) {
    return updatedBits;
  }
  return sumBinary(updatedBits, index + 1, bitArray);
};

const solve3 = (bitArray: number[][]): number => {
  const initialBits = new Array(bitArray[0].length).fill(0);
  const gammaBits = sumBinary(initialBits, 0, bitArray);
  const gammaString = gammaBits.map((bit) => {
    return String(Number(bitArray.length / 2 < bit));
  });

  const epsilonString = gammaBits.map((bit) => {
    return String(Number(bitArray.length / 2 >= bit));
  });
  const calcRate = (bitString: string[]): number => {
    return parseInt(bitString.join(''), 2);
  };
  return calcRate(gammaString) * calcRate(epsilonString);
};

const data = readFileSync(join(__dirname, 'input.txt'), 'utf8')
  .split('\r\n')
  .map((x) => x.split('').map(Number));

console.log(solve3(data));
