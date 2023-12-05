import java.lang.management.ManagementFactory;
import java.lang.management.MemoryUsage;
import java.lang.management.OperatingSystemMXBean;
import java.lang.management.MemoryPoolMXBean;
import java.util.List;


public class JavaPerfTest1 {

    public static void main(String[] args) {
        elapsedTime();

        OperatingSystemMXBean osBean = ManagementFactory.getOperatingSystemMXBean();

        System.out.println("CPU Usage: " + osBean.getSystemLoadAverage() + "%");
        memoryUsage();
    }

    public static void memoryUsage() {

        List<MemoryPoolMXBean> memoryPoolMXBeans = ManagementFactory.getMemoryPoolMXBeans();

        // Iterate through each MemoryPoolMXBean
        for (MemoryPoolMXBean memoryPoolMXBean : memoryPoolMXBeans) {
            String poolName = memoryPoolMXBean.getName();
            MemoryUsage usage = memoryPoolMXBean.getUsage();

            // Display information
            System.out.println("Memory Pool: " + poolName);
            System.out.println("  Committed Memory: " + usage.getCommitted() / (1024 * 1024) + " MB");
            System.out.println("  Used Memory: " + usage.getUsed() / (1024 * 1024) + " MB");
            System.out.println("  Max Memory: " + usage.getMax() / (1024 * 1024) + " MB");
            if (memoryPoolMXBean.isUsageThresholdSupported()) {
                System.out.println("  Usage Threshold: " + memoryPoolMXBean.getUsageThreshold() / (1024 * 1024) + " MB");
            } else {
                System.out.println("  Usage Threshold not supported for this memory pool.");
            }

            System.out.println("--------------------------------------------------");
        }    
    }


    private static String formatMemoryUsage(MemoryUsage memoryUsage) {
        long max = memoryUsage.getMax() / (1024 * 1024);
        long used = memoryUsage.getUsed() / (1024 * 1024);
        long committed = memoryUsage.getCommitted() / (1024 * 1024);

        return "Used: " + used + "MB, Committed: " + committed + "MB, Max: " + max + "MB";
    }

    private static void runYourCode() {
        // Allocate a large array to consume memory
        int n = 1000*1000*1000;
        int[] largeArray = new int[n];

        for (int i = 0; i < n; i++) {
            largeArray[i] = calculateProduct();        
        }
    }


    private static int calculateProduct() {
        int product = 1;
        for (int i = 1; i <= 100*1; i++) {
            product *= i;
        }
        return product;
    }

    private static void elapsedTime() {

        long startTime = System.currentTimeMillis();
        runYourCode();
        long endTime = System.currentTimeMillis();
        long elapsedTime = endTime - startTime;
        System.out.println("Elapsed Time: " + elapsedTime + " milliseconds");

    }
}
