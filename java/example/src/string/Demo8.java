package string;

import java.util.StringTokenizer;

/**
 * Created by ixfosa on 2021/1/27 14:40
 */

// 字符串分隔(StringTokenizer)

// Java 中我们可以使用 StringTokennizer 设置不同分隔符来分隔字符串，默认的分隔符是：
// 空格、制表符（\t）、换行符(\n）、回车符（\r）。

// 以下实例演示了 StringTokennizer 使用空格和等号来分隔字符串：
public class Demo8 {
    public static void main(String[] args) {
        String str = "This is String , split by StringTokenizer, created by ixfosa";
        StringTokenizer st = new StringTokenizer(str);
        while (st.hasMoreElements()) {
            System.out.println(st.nextElement());
        }

        StringTokenizer st1 = new StringTokenizer(str, ",");
        while (st1.hasMoreElements()) {
            System.out.println(st1.nextElement());
        }
    }
}
