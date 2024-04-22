class Stack {
    constructor() {
        this.items = [];
    }

    // Push element to the top of the stack
    push(element) {
        this.items.push(element);
    }

    // Remove and return the top element from the stack
    pop() {
        if (this.isEmpty()) {
            throw new Error("Stack is empty");
        }
        return this.items.pop();
    }

    // Return the top element of the stack without removing it
    peek() {
        if (this.isEmpty()) {
            throw new Error("Stack is empty");
        }
        return this.items[this.items.length - 1];
    }

    // Check if the stack is empty
    isEmpty() {
        return this.items.length === 0;
    }

    // Return the number of elements in the stack
    size() {
        return this.items.length;
    }

    // Print the elements of the stack
    printStack() {
        console.log(this.items.toString());
    }
}

// Example usage:
const stack = new Stack();
stack.push(1);
stack.push(2);
stack.push(3);
stack.printStack(); // Output: 1,2,3
console.log("Size of stack:", stack.size()); // Output: 3
console.log("Top of stack:", stack.peek()); // Output: 3
console.log("Popped element:", stack.pop()); // Output: 3
stack.printStack(); // Output: 1,2
