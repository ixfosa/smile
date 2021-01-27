package string;

/**
 * Created by ixfosa on 2021/1/27 15:08
 */

// 字符串性能比较测试
// 以下实例演示了通过两种方式创建字符串，并测试其性能：

public class Demo11 {
    public static void main(String[] args) {
        String[] strArr1 = new String[50000];
        long startTime1 = System.nanoTime();
        for (int i = 0; i < 50000; i++) {
            strArr1[i] = "hello";
        }
        long endTime1 = System.nanoTime();
        System.out.println("String关键字：" + (endTime1 - startTime1)); // String关键字：646500


        String[] strArr2 = new String[50000];
        long startTime2 = System.nanoTime();
        for (int i = 0; i < 50000; i++) {
            strArr2[i] = new String("hello");
        }
        long endTime2 = System.nanoTime();
        System.out.println("new String：" + (endTime2 - startTime2)); // new String：1785300
    }

}
