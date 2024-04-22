package HashMap;

import java.util.Arrays;

public class HashTable {
    private final int MAX;
    private final Object[] arr;

    public HashTable() {
        this.MAX = 100;
        this.arr = new Object[MAX];
        Arrays.fill(this.arr, null); // Initialize array with null values
    }

    private int getHash(String key) {
        int h = 0;
        for (char c : key.toCharArray()) {
            h += (int) c; // Convert character to ASCII value and add to hash
        }
        return h % MAX; // Get index within array bounds
    }

    public void set(String key, Object value) {
        int index = getHash(key);
        arr[index] = value;
    }

    public Object get(String key) {
        int index = getHash(key);
        return arr[index];
    }

    public void remove(String key) {
        int index = getHash(key);
        arr[index] = null;
    }

    public static void main(String[] args) {
        HashTable t = new HashTable();
        t.set("march 6", 90);
        t.set("april 1", 120);
        t.set("june 24", 190);

        // Retrieving values
        System.out.println("Value for 'march 6': " + t.get("march 6"));
        System.out.println("Value for 'april 1': " + t.get("april 1"));
        System.out.println("Value for 'june 24': " + t.get("june 24"));

        // Removing a key-value pair
        t.remove("april 1");
        System.out.println("Value for 'april 1' after removal: " + t.get("april 1")); // Should print null
    }
}
