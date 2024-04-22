class Node {
    constructor(data) {
        this.data = data;
        this.next = null;
    }
}

class Linked_list {
    constructor() {
        this.head = null;
    }

    insertAtBeginning(data) {
        const newNode = new Node(data);
        newNode.next = this.head;
        this.head = newNode;
    }

    printList() {
        if (!this.head) {
            console.log('Linked List is empty');
            return;
        }

        let itr = this.head;
        let listStr = '';
        while (itr) {
            listStr += `${itr.data}-->`;
            itr = itr.next;
        }
        console.log(listStr);
    }

    insertAtEnd(data) {
        const newNode = new Node(data);
        if (!this.head) {
            this.head = newNode;
            return;
        }
        let itr = this.head;
        while (itr.next) {
            itr = itr.next;
        }
        itr.next = newNode;
    }

    insertValue(dataList) {
        this.head = null;
        for (const data of dataList) {
            this.insertAtEnd(data);
        }
    }

    getLength() {
        let count = 0;
        let itr = this.head;
        while (itr) {
            count++;
            itr = itr.next;
        }
        return count;
    }

    removeAt(index) {
        if (index < 0 || index >= this.getLength()) {
            throw new Error("Invalid index");
        }

        if (index === 0) {
            this.head = this.head.next;
            return;
        }

        let count = 0;
        let itr = this.head;
        while (itr) {
            if (count === index - 1) {
                itr.next = itr.next.next;
                break;
            }
            count++;
            itr = itr.next;
        }
    }

    insertAt(index, data) {
        if (index < 0 || index > this.getLength()) {
            throw new Error("Invalid index");
        }

        if (index === 0) {
            this.insertAtBeginning(data);
            return;
        }

        let count = 0;
        let itr = this.head;
        while (itr) {
            if (count === index - 1) {
                const newNode = new Node(data);
                newNode.next = itr.next;
                itr.next = newNode;
                break;
            }
            count++;
            itr = itr.next;
        }
    }
}

// Example usage:
const li = new Linked_list();
li.insertAtBeginning(1);
li.insertAtBeginning(2);
li.insertAtBeginning(3);
li.insertAtBeginning(4);
li.insertAtEnd(9);
li.insertAtBeginning(0);
li.printList();

const li2 = new Linked_list();
const dataList = ['Orange', 'water', 'cocaca', 'banana'];
li2.insertValue(dataList);
li2.printList();
console.log("li2 length:", li2.getLength());
li2.removeAt(2);
console.log("li2 after remove node");
li2.printList();
console.log("li2 after insert new value to a specific index");
li2.insertAt(1, "cheese");
li2.printList();
