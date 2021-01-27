package string;

//  查找字符串最后一次出现的位置
//  public int lastIndexOf(String str)
public class Demo2 {
    public static void main(String[] args) {
        String str = "hello ixfosa, hello long, hello zhong";
        int idx = str.lastIndexOf("long");

        if (idx == -1) {
            System.out.println("没有找到字符串 long");
        } else {
            System.out.println("找到" + idx);
        }
    }
}
