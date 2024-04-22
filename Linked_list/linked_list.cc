#include <iostream>

class Node {
public:
    int value;
    Node* next;

    Node(int data) : value(data), next(nullptr) {}
};

class Linked_list {
private:
    Node* head;

public:
    Linked_list() : head(nullptr) {}

    void insert_at_beginning(int data) {
        Node* newNode = new Node(data);
        newNode->next = head;
        head = newNode;
    }

    void print_list() {
        if (head == nullptr) {
            std::cout << "Linked List is empty" << std::endl;
            return;
        }

        Node* itr = head;
        while (itr != nullptr) {
            std::cout << itr->value << "--";
            itr = itr->next;
        }
        std::cout << std::endl;
    }

    void insert_at_end(int data) {
        if (head == nullptr) {
            head = new Node(data);
            return;
        }

        Node* itr = head;
        while (itr->next != nullptr) {
            itr = itr->next;
        }
        itr->next = new Node(data);
    }

    void insert_value(int* dataList, int size) {
        head = nullptr;
        for (int i = 0; i < size; ++i) {
            insert_at_end(dataList[i]);
        }
    }

    int get_length() {
        int count = 0;
        Node* itr = head;
        while (itr != nullptr) {
            count++;
            itr = itr->next;
        }
        return count;
    }

    void remove_at(int index) {
        if (index < 0 || index >= get_length()) {
            throw std::invalid_argument("Invalid index");
        }

        if (index == 0) {
            Node* temp = head;
            head = head->next;
            delete temp;
            return;
        }

        int count = 0;
        Node* itr = head;
        while (itr != nullptr) {
            if (count == index - 1) {
                Node* temp = itr->next;
                itr->next = itr->next->next;
                delete temp;
                break;
            }
            count++;
            itr = itr->next;
        }
    }

    void insert_at(int index, int data) {
        if (index < 0 || index > get_length()) {
            throw std::invalid_argument("Invalid index");
        }

        if (index == 0) {
            insert_at_beginning(data);
            return;
        }

        int count = 0;
        Node* itr = head;
        while (itr != nullptr) {
            if (count == index - 1) {
                Node* newNode = new Node(data);
                newNode->next = itr->next;
                itr->next = newNode;
                break;
            }
            count++;
            itr = itr->next;
        }
    }

    ~Linked_list() {
        Node* temp;
        while (head != nullptr) {
            temp = head;
            head = head->next;
            delete temp;
        }
    }
};

int main() {
    Linked_list li;
    li.insert_at_beginning(1);
    li.insert_at_beginning(2);
    li.insert_at_beginning(3);
    li.insert_at_beginning(4);
    li.insert_at_end(9);
    li.insert_at_beginning(0);
    li.print_list();
    
    Linked_list li2;
    int data[] = {2, 4, 6, 8};
    li2.insert_value(data, sizeof(data) / sizeof(data[0]));
    li2.print_list();
    std::cout << "li2 length: " << li2.get_length() << std::endl;
    li2.remove_at(2);
    std::cout << "li2 after remove node" << std::endl;
    li2.print_list();
    std::cout << "li2 after insert new value to a specific index" << std::endl;
    li2.insert_at(1, 5);
    li2.print_list();

    return 0;
}
