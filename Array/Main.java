import java.util.*;

public class Main {
    public static void main(String[] args) {
        int[] stockPrices = new int[5];

        stockPrices[0] = 298;
        stockPrices[1] = 305;
        stockPrices[2] = 312;
        stockPrices[3] = 319;
        stockPrices[4] = 323;

        ArrayList<Integer> stockPricesList = new ArrayList<Integer>();

        stockPricesList.add(298);
        stockPricesList.add(305);
        stockPricesList.add(312);
        stockPricesList.add(319);
        stockPricesList.add(323);

        System.out.println("Printing using array:");
        for (int i = 0; i < 5; i++) {
            System.out.println(stockPrices[i]);
        }

        System.out.println("Printing using ArrayList:");
        for (int price : stockPricesList) {
            System.out.println(price);
        }
    }
}
