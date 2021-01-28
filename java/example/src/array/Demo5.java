package array;

import java.util.Arrays;

/**
 * Created by ixfosa on 2021/1/27 21:05
 */

// 数组输出
public class Demo5 {
    public static void main(String[] args) {
        String[] arr = {"ixfosa", "long", "zhong"};

        System.out.println("Arrays.toString: " + Arrays.toString(arr));

        System.out.println("-------------for----------");
        for (int i = 0; i < arr.length; i++) {
            String s = arr[i];
            System.out.println(s);
        }

        System.out.println("-------------for in----------");
        for (String s : arr) {
            System.out.println(s);
        }
    }
}
