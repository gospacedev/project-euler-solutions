public class p2 {
    public static void main(String[] args) {
        int x = 1, y = 2, n = 2;

        while (y < 4000000) {
            x += y;
            if (x % 2 == 0) {
                n += x;
            }

            y += x;
            if (y % 2 == 0) {
                n += y;
            }
        }

        System.out.println(n);
    }
}
