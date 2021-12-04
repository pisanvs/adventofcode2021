import { readFile } from "fs";

readFile("../inputs/3.txt", "utf8", (err, data) => {
    if (err) console.log(err);
    const input = data.split("\n");

    var zeroes = [0,0,0,0,0,0,0,0,0,0,0,0];
    var ones = [0,0,0,0,0,0,0,0,0,0,0,0];
    
    input.forEach((line, i) => {
        line.split("").forEach((char, ii) => {
            if (char === '0') {
                zeroes[ii]++;
            }
            if (char === '1') {
                ones[ii]++;
            }
        });
    });

    var final = [0,0,0,0,0,0,0,0,0,0,0,0]

    zeroes.forEach((num, i) => {
        if (num > ones[i]) {
            final[i] = '0';
        } else {
            final[i] = '1';
        }
    });

    console.log(final.join(''));

    var epsilon = [0,0,0,0,0,0,0,0,0,0,0,0];

    final.forEach((num, i) => {
        if (num === '0') {
            epsilon[i] = '1';
        } else {
            epsilon[i] = '0';
        }
    });

    console.log(epsilon.join(''));

    epsilon = parseInt(epsilon.join(''), 2);
    var gamma = parseInt(final.join(''), 2);

    console.log(epsilon*gamma);
    // finished at 23:07
})