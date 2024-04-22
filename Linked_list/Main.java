import java.util.*;

class Node {
    int value;
    Node next;

    Node(int value) {
        this.value = value;
        this.next = null;
    }
}

class Linked_list {
    Node head;

    void insert_at_beginning(int data) {
        Node newNode = new Node(data);
        newNode.next = head;
        head = newNode;
    }

    void print_list() {
        if (head == null) {
            System.out.println("Linked List is empty");
            return;
        }

        Node itr = head;
        while (itr != null) {
            System.out.print(itr.value + "--");
            itr = itr.next;
        }
        System.out.println();
    }

    void insert_at_end(int data) {
        if (head == null) {
            head = new Node(data);
            return;
        }

        Node itr = head;
        while (itr.next != null) {
            itr = itr.next;
        }
        itr.next = new Node(data);
    }

    void insert_value(int[] dataList) {
        head = null;
        for (int data : dataList) {
            insert_at_end(data);
        }
    }

    int get_length() {
        int count = 0;
        Node itr = head;
        while (itr != null) {
            count++;
            itr = itr.next;
        }
        return count;
    }

    void remove_at(int index) {
        if (index < 0 || index >= get_length()) {
            throw new IllegalArgumentException("Invalid index");
        }

        if (index == 0) {
            head = head.next;
            return;
        }

        int count = 0;
        Node itr = head;
        while (itr != null) {
            if (count == index - 1) {
                itr.next = itr.next.next;
                break;
            }
            count++;
            itr = itr.next;
        }
    }

    void insert_at(int index, int data) {
        if (index < 0 || index > get_length()) {
            throw new IllegalArgumentException("Invalid index");
        }

        if (index == 0) {
            insert_at_beginning(data);
            return;
        }

        int count = 0;
        Node itr = head;
        while (itr != null) {
            if (count == index - 1) {
                Node newNode = new Node(data);
                newNode.next = itr.next;
                itr.next = newNode;
                break;
            }
            count++;
            itr = itr.next;
        }
    }
}

public class Main {
    public static void main(String[] args) {
        Linked_list li = new Linked_list();
        li.insert_at_beginning(1);
        li.insert_at_beginning(2);
        li.insert_at_beginning(3);
        li.insert_at_beginning(4);
        li.insert_at_end(9);
        li.insert_at_beginning(0);
        li.print_list();
        
        Linked_list li2 = new Linked_list();
        int[] data = {2, 4, 6, 8};
        li2.insert_value(data);
        li2.print_list();
        System.out.println("li2 length: " + li2.get_length());
        li2.remove_at(2);
        System.out.println("li2 after remove node");
        li2.print_list();
        System.out.println("li2 after insert new value to a specific index");
        li2.insert_at(1, 5);
        li2.print_list();
    }
}
