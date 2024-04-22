class HashTable {
    constructor() {
        this.MAX = 100;
        this.arr = new Array(this.MAX).fill(null);
    }

    getHash(key) {
        let h = 0;
        for (let char of key) {
            h += char.charCodeAt(0);
        }
        return h % this.MAX;
    }

    set(key, value) {
        const h = this.getHash(key);
        this.arr[h] = value;
    }

    get(key) {
        const h = this.getHash(key);
        return this.arr[h];
    }

    remove(key) {
        const h = this.getHash(key);
        this.arr[h] = null;
    }
}

// Example usage:
const t = new HashTable();
t.set("march 6", 90);
t.set("april 1", 120);
t.set("june 24", 190);

// Retrieving values
console.log("Value for 'march 6':", t.get("march 6"));
console.log("Value for 'april 1':", t.get("april 1"));
console.log("Value for 'june 24':", t.get("june 24"));

// Removing a key-value pair
t.remove("april 1");
console.log("Value for 'april 1' after removal:", t.get("april 1")); // Should print null
