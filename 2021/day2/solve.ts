import { readFileSync } from 'fs';
import { join } from 'path';

type Location = {
  depth: number;
  horizontal: number;
  aim: number;
};

const solve2b = (data: Array<string>): number => {
  const parseCommand = (
    data: Array<string>,
    index: number,
    location: Location
  ): Location => {
    const [command, amountString]: Array<string> = data[index].split(' ');
    const amount: number = +amountString;
    let { depth, horizontal, aim } = location;
    switch (command) {
      case 'forward':
        horizontal += amount;
        depth += aim * amount;
        break;
      case 'up':
        aim -= amount;
        break;
      case 'down':
        aim += amount;
        break;
      default:
        console.error(`invalid command: ${command}`);
        break;
    }
    if (index === data.length - 1) return { depth, horizontal, aim };

    return parseCommand(data, index + 1, { depth, horizontal, aim });
  };

  const initialLocation = Object.freeze({ depth: 0, horizontal: 0, aim: 0 });
  const location = parseCommand(data, 0, initialLocation);
  return location.depth * location.horizontal;
};

const data = readFileSync(join(__dirname, 'input.txt'), 'utf8').split('\n');

console.log(solve2b(data));
