

class LinkedList {
    #head = undefined;
    #tail = undefined;

    enqueue(value) {
        const newNode = { value };
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
        } else {
            if (this.tail) {
                this.tail.next = newNode;
            }
            this.tail = newNode;
        }
    }

    dequeue() {
        if (!this.head) {
            return null;
        }
        const value = this.head.value;
        this.head = this.head.next;
        return value;
    }
    peek() {
        if (!this.head) {
            return null;
        }
        return this.head.value;
    }
}

const queue = new LinkedList();
queue.enqueue(1);
queue.enqueue(2);
queue.enqueue(3);
console.log(queue.dequeue()); // 1
console.log(queue.dequeue()); // 2
console.log(queue.peek()); // 3