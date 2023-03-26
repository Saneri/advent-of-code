const fs = require("fs");

const types = {
  INT: "int",
  STRING: "string",
};

function readInput(path, type) {
  return new Promise(function (resolve, reject) {
    fs.readFile(path, "utf8", (err, data) => {
      if (err) {
        console.error(err);
        reject(err);
      }

      switch (type) {
        case types.INT:
          resolve(data.split("\r\n").map((x) => parseInt(x)));
        case types.STRING:
          resolve(data.split("\r\n"));
      }
    });
  });
}

module.exports = { readInput, types };
