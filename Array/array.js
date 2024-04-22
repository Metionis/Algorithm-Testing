let stock_prize = [2, 3, 5, 6];
let stock_names = ["APPL", "IBM", "TATA"];
let stock_data = [
    { "ticker": "APPL", "price": 302 },
    { "ticker": "IBM", "price": 312 },
    { "ticker": "TATA", "price": 354 }
];

// Dimension Array
let stock_prizes = [
    [2, 3, 5, 6],
    [12, 14, 15, 16],
    [5, 7, 1, 9]
];

// Printing stock_prize
console.log("stock_prize:");
for (let price of stock_prize) {
    console.log(price);
}

// Printing stock_names
console.log("\nstock_names:");
for (let name of stock_names) {
    console.log(name);
}

// Printing stock_data
console.log("\nstock_data:");
for (let data of stock_data) {
    console.log(`Ticker: ${data['ticker']}, Price: ${data['price']}`);
}

// Printing stock_prizes
console.log("\nstock_prizes:");
for (let row of stock_prizes) {
    for (let prize of row) {
        process.stdout.write(prize + " ");
    }
    console.log(); // Print new line for each row
}
