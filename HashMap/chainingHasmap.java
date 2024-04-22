package HashMap;

import java.util.ArrayList;
import java.util.List;

public class HashTable {
    private final int MAX;
    private final List<List<Pair>> arr;

    public HashTable() {
        this.MAX = 100;
        this.arr = new ArrayList<>(MAX);
        for (int i = 0; i < MAX; i++) {
            this.arr.add(new ArrayList<>());
        }
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
        boolean found = false;
        for (Pair pair : arr.get(index)) {
            if (pair.getKey().equals(key)) {
                pair.setValue(value);
                found = true;
                break;
            }
        }
        if (!found) {
            arr.get(index).add(new Pair(key, value));
        }
    }

    public Object get(String key) {
        int index = getHash(key);
        for (Pair pair : arr.get(index)) {
            if (pair.getKey().equals(key)) {
                return pair.getValue();
            }
        }
        return null;
    }

    public void remove(String key) {
        int index = getHash(key);
        arr.get(index).removeIf(pair -> pair.getKey().equals(key));
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

    static class Pair {
        private final String key;
        private Object value;

        public Pair(String key, Object value) {
            this.key = key;
            this.value = value;
        }

        public String getKey() {
            return key;
        }

        public Object getValue() {
            return value;
        }

        public void setValue(Object value) {
            this.value = value;
        }
    }
}
