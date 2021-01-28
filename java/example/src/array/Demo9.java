package array;

import java.lang.reflect.Array;
import java.util.Arrays;

/**
 * Created by ixfosa on 2021/1/28 15:04
 */

// 数组扩容
public class Demo9 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4};

        // 手动实现单一类型数组扩容
        int[] newArr1 = new int[10];
        for (int i = 0; i < arr.length; i++) {
            newArr1[i] = arr[i];
        }
        System.out.println(Arrays.toString(newArr1)); // [1, 2, 3, 4, 0, 0, 0, 0, 0, 0]


        System.out.println("------------------------------");


        // System.arraycopy() 方法实现数组扩容
        int[] newArr2 = new int[10];
        System.arraycopy(arr, 0, newArr2, 0, arr.length);
        System.out.println(Arrays.toString(newArr2)); // [1, 2, 3, 4, 0, 0, 0, 0, 0, 0]


        System.out.println("------------------------------");


        // 扩展已经填满的数组方法 Arrays.copyOf()
        int[] arr3 = {1, 2, 3, 4};
        int[] copy = Arrays.copyOf(arr3, arr3.length << 1);
        System.out.println(Arrays.toString(copy));
    }
}
