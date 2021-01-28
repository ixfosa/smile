package array;

import java.util.Arrays;

/**
 * Created by ixfosa on 2021/1/27 20:24
 */
// 数组添加元素
public class Demo2 {
    public static void main(String[] args) {
        int[] arr = {1, 3, 4, 5, 6};
        System.out.println(Arrays.toString(insertElement(arr, 2, 1))); // [1, 2, 3, 4, 5, 6]
        System.out.println(Arrays.toString(insertElement(arr, 0, 0))); // [0, 1, 3, 4, 5, 6]
    }

    // arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
    // Object src : 原数组
    // int srcPos : 从元数据的起始位置开始
    // Object dest : 目标数组
    // int destPos : 目标数组的开始起始位置
    // int length  : 要copy的数组的长度

    public static int[] insertElement(int[] arr, int element, int index) {
        int[] target = new int[arr.length + 1];
        System.arraycopy(arr, 0, target,0,  index);
        target[index] = element;
        System.arraycopy(arr, index, target, index + 1, arr.length - index);
        return target;
    }
}
