package string;

//  删除字符串中的一个字符
//   substring() 函数来删除字符串中的一个字符，我们将功能封装在 removeCharAt 函数中。
public class Demo3 {
    public static void main(String[] args) {
        String str = "hello ixfosa!";
        str = removeCharAt(str, 0);
        System.out.println(str);
    }

    public static String removeCharAt(String str, int pos) {
        return str.substring(0, pos) + str.substring(pos + 1);
    }
}
