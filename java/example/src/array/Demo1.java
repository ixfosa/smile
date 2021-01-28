package array;

import java.util.Arrays;

/**
 * Created by ixfosa on 2021/1/27 20:17
 */

// 数组排序及元素查找
public class Demo1 {
    public static void main(String[] args) {
        int[] array = { 2, 5, -2, 6, -3, 8, 0, -7, -9, 4 };
        Arrays.sort(array);
        System.out.println("sort: " + Arrays.toString(array)); // sort: [-9, -7, -3, -2, 0, 2, 4, 5, 6, 8]

        int i = Arrays.binarySearch(array, -7);
        System.out.println(i); // 元素 -7 在第 1 个位置
    }
}
