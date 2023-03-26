import { readFileSync } from 'fs';
import { join } from 'path';

type Bingo = { bingoNumbers: number[]; bingoBoards: number[][][] };

const parseInput = (filename: string): Bingo => {
  const data = readFileSync(filename, 'utf8').split('\r\n\r\n');
  const bingoBoards = data
    .slice(1)
    .map((board) =>
      board.split('\r\n').map((row) => row.split(/\s+/).map(Number))
    );
  const bingoNumbers = data[0].split(',').map(Number);

  return {
    bingoNumbers,
    bingoBoards
  };
};

const solve4 = (bingoData: Bingo): number => {
  return -99;
};

const bingoData = parseInput(join(__dirname, 'input.txt'));
console.log(solve4(bingoData));
