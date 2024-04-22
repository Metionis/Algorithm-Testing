#include <iostream>
#include <vector>
#include <stdexcept>

template<typename T>
class Stack {
private:
    std::vector<T> items;

public:
    // Push element onto the stack
    void push(const T& item) {
        items.push_back(item);
    }

    // Remove and return the top element from the stack
    T pop() {
        if (isEmpty()) {
            throw std::out_of_range("Stack is empty");
        }
        T item = items.back();
        items.pop_back();
        return item;
    }

    // Return the top element of the stack without removing it
    T peek() {
        if (isEmpty()) {
            throw std::out_of_range("Stack is empty");
        }
        return items.back();
    }

    // Check if the stack is empty
    bool isEmpty() {
        return items.empty();
    }

    // Return the number of elements in the stack
    size_t size() {
        return items.size();
    }

    // Print the elements of the stack
    void printStack() {
        for (const T& item : items) {
            std::cout << item << " ";
        }
        std::cout << std::endl;
    }
};

int main() {
    Stack<int> stack;
    stack.push(1);
    stack.push(2);
    stack.push(3);
    
    std::cout << "Stack: ";
    stack.printStack(); // Output: 1 2 3
    
    std::cout << "Size of stack: " << stack.size() << std::endl; // Output: 3
    
    std::cout << "Top of stack: " << stack.peek() << std::endl; // Output: 3
    
    std::cout << "Popped element: " << stack.pop() << std::endl; // Output: 3
    
    std::cout << "Stack after pop: ";
    stack.printStack(); // Output: 1 2
    
    return 0;
}
