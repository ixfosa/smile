package array;

/**
 * Created by ixfosa on 2021/1/27 20:41
 */

// 获取数组长度
public class Demo3 {
    public static void main(String[] args) {
        int[][] arr = new int[2][3];
        System.out.println("第一维数组长度: " + arr.length); // 第一维数组长度: 2
        System.out.println("第二维数组长度: " + arr[0].length); // 第二维数组长度: 3
        System.out.println("第二维数组长度: " + arr[1].length); // 第二维数组长度: 3
    }
}
