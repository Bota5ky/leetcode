package design.implement_stack_using_queues;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/implement-stack-using-queues/">225. 用队列实现栈</a>
 * @since 2024-03-03 21:33
 */
class MyStack {
    Queue<Integer> out;
    Queue<Integer> in;

    public MyStack() {
        out = new LinkedList<>();
        in = new LinkedList<>();
    }

    public void push(int x) {
        in.offer(x);
        while (!out.isEmpty()) {
            in.offer(out.poll());
        }
        Queue<Integer> temp = in;
        in = out;
        out = temp;
    }

    public int pop() {
        return out.poll();
    }

    public int top() {
        return out.peek();
    }

    public boolean empty() {
        return out.isEmpty();
    }
}

/*
 * Your MyStack object will be instantiated and called as such:
 * MyStack obj = new MyStack();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.top();
 * boolean param_4 = obj.empty();
 */
