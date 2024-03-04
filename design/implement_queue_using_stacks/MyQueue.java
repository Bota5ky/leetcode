package design.implement_queue_using_stacks;

import java.util.Stack;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/implement-queue-using-stacks/">232. 用栈实现队列</a>
 * @since 2024-03-04 08:32
 */
class MyQueue {
    Stack<Integer> in;
    Stack<Integer> out;

    public MyQueue() {
        in = new Stack<>();
        out = new Stack<>();
    }

    public void push(int x) {
        in.push(x);
    }

    public int pop() {
        if (out.isEmpty()) {
            while (!in.isEmpty()) {
                out.push(in.pop());
            }
        }
        return out.pop();
    }

    public int peek() {
        if (out.isEmpty()) {
            while (!in.isEmpty()) {
                out.push(in.pop());
            }
        }
        return out.peek();
    }

    public boolean empty() {
        return out.isEmpty() && in.isEmpty();
    }
}

/*
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue obj = new MyQueue();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.peek();
 * boolean param_4 = obj.empty();
 */
