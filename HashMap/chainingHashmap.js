class HashTable {
    constructor() {
        this.MAX = 100;
        this.arr = new Array(this.MAX).fill(null).map(() => []);
    }

    getHash(key) {
        let h = 0;
        for (let char of key) {
            h += char.charCodeAt(0);
        }
        return h % this.MAX;
    }

    setItem(key, value) {
        const h = this.getHash(key);
        let found = false;
        for (let i = 0; i < this.arr[h].length; i++) {
            const element = this.arr[h][i];
            if (element.length === 2 && element[0] === key) {
                this.arr[h][i] = [key, value];
                found = true;
                break;
            }
        }
        if (!found) {
            this.arr[h].push([key, value]);
        }
    }

    getItem(key) {
        const h = this.getHash(key);
        for (let element of this.arr[h]) {
            if (element[0] === key) {
                return element[1];
            }
        }
    }

    deleteItem(key) {
        const h = this.getHash(key);
        for (let i = 0; i < this.arr[h].length; i++) {
            const element = this.arr[h][i];
            if (element[0] === key) {
                this.arr[h].splice(i, 1);
                break;
            }
        }
    }
}

// Example usage:
const ht = new HashTable();
ht.setItem('name', 'John');
ht.setItem('age', 30);
console.log(ht.getItem('name')); // Output: John
ht.deleteItem('name');
console.log(ht.getItem('name')); // Output: undefined
