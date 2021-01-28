package array;


import java.util.Arrays;

/**
 * Created by ixfosa on 2021/1/27 21:45
 */

// 数组填充
// 以下实例我们通过 Java Util 类的 Arrays.fill(arrayname,value) 方法和
// Arrays.fill(arrayname ,starting index ,ending index ,value) 方法向数组中填充元素：

public class Demo8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4};
        Arrays.fill(arr, 100);
        System.out.println(Arrays.toString(arr)); // [100, 100, 100, 100]

        Arrays.fill(arr, 2, 4, 66);
        System.out.println(Arrays.toString(arr)); // [100, 100, 66, 66]
    }
}
