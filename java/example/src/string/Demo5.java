package string;


// 字符串反转
public class Demo5 {
    public static void main(String[] args) {
        String str = "hello ixfosa";
        String reverseStr = new StringBuilder(str).reverse().toString();
       // System.out.println(reverseStr);

        System.out.println(reverse(str));
    }

    public static String reverse(String str) {
        char[] chars = str.toCharArray();
        for (int i = 0, j = str.length() - 1; i < j; i++, j--) {
            char temp = chars[i];
            chars[i] = chars[j];
            chars[j] = temp;
        }
        return new String(chars);
    }
}
