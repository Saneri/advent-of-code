const _ = require("lodash");
const assert = require("assert");

const utils = require("../utils");

const mandatoryFields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];
const optionalFields = ["cid"];
const eyeColors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];

function solver(input) {
  let validPassports = 0;
  let fields = Array.from(mandatoryFields);

  input.forEach((line) => {
    const elements = line.split(" ");
    if (elements[0] !== "") {
      elements.forEach((element) => {
        //console.log(element.split(":"));
        const key = element.split(":")[0];
        const value = element.split(":")[1];
        //console.log(validate(key, value));
        if (fields.includes(key) && validate(key, value)) {
          fields.splice(fields.indexOf(key), 1);
        } else if (
          !mandatoryFields.includes(key) &&
          !optionalFields.includes(key)
        ) {
          throw new Error("invalid key error");
        }
      });
    } else {
      if (_.isEmpty(fields)) {
        validPassports++;
      }
      fields = Array.from(mandatoryFields);
    }
  });
  if (_.isEmpty(fields)) {
    validPassports++;
  }
  fields = Array.from(mandatoryFields);
  return validPassports;
}

function validate(key, value) {
  switch (key) {
    case "byr":
      return parseInt(value) >= 1920 && parseInt(value) <= 2002;
    case "iyr":
      return parseInt(value) >= 2010 && parseInt(value) <= 2020;
    case "eyr":
      return parseInt(value) >= 2020 && parseInt(value) <= 2030;
    case "hgt":
      try {
        const length = parseInt(value.slice(0, -2));
        const unit = value.slice(-2);
        if (unit === "cm" && 150 <= length && length <= 193) return true;
        else if (unit === "in" && 59 <= length && length <= 76) return true;
      } catch {}
      return false;
    case "hcl":
      return /^#[0-9a-j]{6}$/.test(value);
    case "ecl":
      return eyeColors.includes(value);
    case "pid":
      return /^[0-9]{9}$/.test(value);
  }
}

assert(validate("byr", "abc") === false);
assert(validate("iyr", "2011") === true);
assert(validate("hgt", "190cm") === true);
assert(validate("hcl", "#abcaa1") === true);
assert(validate("ecl", "gry") === true);
assert(validate("pid", "003456789") === true);

assert(
  solver([
    "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
    "byr:1937 iyr:2017 cid:147 hgt:183cm",
    "",
    "eyr:1972 cid:100",
    "hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
    "",
    "eyr:2029 ecl:blu cid:129 byr:1989",
    "iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
  ]) === 2
);

async function main() {
  const input = await utils.readInput("day4/input.txt", utils.types.STRING);
  console.log(solver(input));
}

main();
