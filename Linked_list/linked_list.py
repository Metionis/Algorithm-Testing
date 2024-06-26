class Node:
    def __init__ (self, value = None, next = None):
        self.next = next;
        self.value = value;

class Linked_list:
    def __init__ (self):
        self.head = None

    def insert_at_beginning (self, data):
        node = Node(data, self.head)
        self.head = node

    def print_list (self):
        if self.head is None:
            print('Linked List is empty')
            return
        
        itr = self.head
        listr = ''
        while itr:
            listr += str(itr.value) + '--'
            itr = itr.next
        print(listr)

    def insert_at_end(self, data):
        if self.head is None:
            self.head = Node(data, None)
            return
    
        itr = self.head
        while itr.next:
            itr = itr.next

        itr.next = Node(data, None)

    def insert_value (self, data_list):
        self.head = None
        for data in data_list:
            self.insert_at_end(data)

    def get_length (self):
        count = 0
        itr = self.head
        while itr:
            count += 1
            itr = itr.next

        return count
    
    def remove_at (self, index):
        if index < 0 or index >= self.get_length():
            raise Exception("Invalid Value of index")

        if index == 0:
            self.head = self.head.next
            return
        
        count = 0
        itr = self.head
        while itr:
            if count == index - 1:
                itr.next = itr.next.next
                break

            count += 1
            itr = itr.next


    def insert_at(self, index, data):
        if index < 0 or index > self.get_length():
            raise Exception("Invalid Value of index")
        if index == 0:
            self.insert_at_beginning(data)
            return
        
        count = 0
        itr = self.head
        while itr:
            if count == index - 1:
                node = Node(data, itr.next)
                itr.next = node
                return

            count += 1
            itr = itr.next

        
if __name__ == '__main__':
    li = Linked_list()
    li.insert_at_beginning(1)
    li.insert_at_beginning(2)
    li.insert_at_beginning(3)
    li.insert_at_beginning(4)
    li.insert_at_end(9)
    li.insert_at_beginning(0)
    li.print_list()
            
    li2 = Linked_list()
    data_list = ['Orange', 'water', 'cocaca', 'banana']
    li2.insert_value(data_list)
    li2.print_list()
    print("li2 length: ", li2.get_length())
    li2.remove_at(2)
    print("li2 after remove node")
    li2.print_list()
    print("li2 after insert new value to a specific Aindex")
    li2.insert_at(1, "cheese")
    li2.print_list()