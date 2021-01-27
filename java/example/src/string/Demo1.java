package string;

// 字符串比较
// 通过字符串函数 compareTo (string) ，compareToIgnoreCase(String) 及 compareTo(object string) 来比较两个字符串，
// 并返回字符串中第一个字母ASCII的差值。
public class Demo1 {
    public static void main(String[] args) {
        String str = "Hello ixfosa";
        String anotherStr = "hello ixfosa";
        Object objStr = str;
        System.out.println(str.compareTo(anotherStr)); // -32
        System.out.println(str.compareToIgnoreCase(anotherStr)); //0
        System.out.println(str.compareTo(objStr.toString())); //0
    }
}
