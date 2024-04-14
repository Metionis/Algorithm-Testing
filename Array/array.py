stock_prize = [2, 3, 5, 6]
stock_names = ["APPL", "IBM", "TATA"]
stock_data = [
    {"ticker": "APPL", "price": 302},
    {"ticker": "IBM", "price": 312},
    {"ticker": "TATA", "price": 354},
]

# Dimension Array
stock_prizes = [
    [2, 3, 5, 6],
    [12, 14, 15, 16],
    [5, 7, 1, 9]
]

# Printing stock_prize
print("stock_prize:")
for price in stock_prize:
    print(price)

# Printing stock_names
print("\nstock_names:")
for name in stock_names:
    print(name)

# Printing stock_data
print("\nstock_data:")
for data in stock_data:
    print(f"Ticker: {data['ticker']}, Price: {data['price']}")

# Printing stock_prizes
print("\nstock_prizes:")
for row in stock_prizes:
    for prize in row:
        print(prize, end=" ")
    print()  # Print new line for each row
